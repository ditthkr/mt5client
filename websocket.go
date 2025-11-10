package mt5client

import (
	"encoding/json"
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/gorilla/websocket"
)

// WebSocketClient WebSocket client
type WebSocketClient struct {
	client          *Client
	connections     map[string]*websocket.Conn // key = path เช่น "/OnQuote"
	handlers        *EventHandlers
	mu              sync.RWMutex
	done            chan struct{}
	isActive        bool
	autoReconnect   bool
	reconnectDelay  time.Duration
	subscribedPaths map[string]func([]byte) // เก็บ paths ที่ subscribe ไว้สำหรับ reconnect
}

// EventHandlers handlers สำหรับ events ต่างๆ
type EventHandlers struct {
	OnConnect             func()
	OnDisconnect          func()
	OnError               func(error)
	OnQuote               func(*Quote)
	OnTickValue           func(*TickValueEvent)
	OnOrderUpdate         func(*OrderUpdateEvent)
	OnOrderProfit         func(*OrderProfitEvent)
	OnMarketWatch         func(*MarketWatch)
	OnTickHistory         func(*TickHistoryEvent)
	OnMail                func(*Mail)
	OnOpenedOrdersTickets func([]int64)
	OnOrderBook           func(*OrderBook)
	OnOhlc                func(*OhlcData)
}

type SocketResponse struct {
	Type string          `json:"type"`
	Data json.RawMessage `json:"data"`
}

// TickValueEvent event สำหรับ tick value
type TickValueEvent struct {
	Symbol    string  `json:"symbol"`
	TickValue float64 `json:"tickValue"`
	Time      string  `json:"time"`
}

// OrderUpdateEvent event สำหรับ order update
type OrderUpdateEvent struct {
	Update struct {
		Type  string `json:"type"`
		Order Order  `json:"order"`
	} `json:"update"`
}

// OrderProfitEvent event สำหรับกำไร/ขาดทุน
type OrderProfitEvent struct {
	Orders []Order `json:"orders"`
}

// TickHistoryEvent event สำหรับ tick history
type TickHistoryEvent struct {
	Symbol string `json:"symbol"`
	Ticks  []Tick `json:"ticks"`
}

// NewWebSocketClient สร้าง WebSocket client
func (c *Client) NewWebSocketClient() *WebSocketClient {
	return &WebSocketClient{
		client:          c,
		connections:     make(map[string]*websocket.Conn),
		handlers:        &EventHandlers{},
		done:            make(chan struct{}),
		autoReconnect:   true, // เปิด auto-reconnect โดยdefault
		reconnectDelay:  5 * time.Second,
		subscribedPaths: make(map[string]func([]byte)),
	}
}

// SetAutoReconnect ตั้งค่า auto-reconnect
func (ws *WebSocketClient) SetAutoReconnect(enable bool, delay time.Duration) {
	ws.mu.Lock()
	defer ws.mu.Unlock()
	ws.autoReconnect = enable
	if delay > 0 {
		ws.reconnectDelay = delay
	}
}

// SetHandlers ตั้งค่า event handlers
func (ws *WebSocketClient) SetHandlers(handlers *EventHandlers) {
	ws.mu.Lock()
	defer ws.mu.Unlock()
	ws.handlers = handlers
}

// Connect เชื่อมต่อ WebSocket (เตรียมพร้อม)
func (ws *WebSocketClient) Connect() error {
	ws.mu.Lock()
	defer ws.mu.Unlock()

	if ws.isActive {
		return fmt.Errorf("already connected")
	}

	if ws.client.token == "" {
		return fmt.Errorf("not connected to MT5")
	}

	ws.isActive = true
	ws.done = make(chan struct{})

	// เรียก OnConnect handler
	if ws.handlers.OnConnect != nil {
		go ws.handlers.OnConnect()
	}

	return nil
}

// connectToPath เชื่อมต่อไปยัง path เฉพาะ
func (ws *WebSocketClient) connectToPath(path string, handler func([]byte)) error {
	if ws.client.token == "" {
		return fmt.Errorf("not connected to MT5")
	}

	// แปลง http:// เป็น ws://
	wsURL := ws.client.baseURL
	if len(wsURL) > 7 && wsURL[:7] == "http://" {
		wsURL = "ws://" + wsURL[7:]
	} else if len(wsURL) > 8 && wsURL[:8] == "https://" {
		wsURL = "wss://" + wsURL[8:]
	}

	// สร้าง WebSocket URL: ws://host:port/OnQuote?id=token
	url := fmt.Sprintf("%s%s?id=%s", wsURL, path, ws.client.token)

	conn, _, err := websocket.DefaultDialer.Dial(url, nil)
	if err != nil {
		return fmt.Errorf("failed to connect to %s: %w", path, err)
	}

	ws.mu.Lock()
	ws.connections[path] = conn
	ws.subscribedPaths[path] = handler // เก็บไว้สำหรับ reconnect
	ws.mu.Unlock()

	// อ่านข้อความใน goroutine
	go ws.readMessages(path, conn, handler)

	return nil
}

// reconnectPath reconnect ไปยัง path
func (ws *WebSocketClient) reconnectPath(path string, handler func([]byte)) {
	for {
		select {
		case <-ws.done:
			return
		default:
			ws.mu.RLock()
			shouldReconnect := ws.autoReconnect && ws.isActive
			delay := ws.reconnectDelay
			ws.mu.RUnlock()

			if !shouldReconnect {
				return
			}

			log.Printf("Reconnecting to %s in %v...", path, delay)
			time.Sleep(delay)

			err := ws.connectToPath(path, handler)
			if err != nil {
				log.Printf("Failed to reconnect to %s: %v", path, err)
				continue
			}

			log.Printf("✅ Reconnected to %s", path)
			return
		}
	}
}

// SubscribeQuote subscribe รับ quote real-time
func (ws *WebSocketClient) SubscribeQuote() error {
	return ws.connectToPath("/OnQuote", ws.handleQuoteMessage)
}

// SubscribeTickValue subscribe รับ tick value
func (ws *WebSocketClient) SubscribeTickValue() error {
	return ws.connectToPath("/OnTickValue", ws.handleTickValueMessage)
}

// SubscribeOrderUpdate subscribe รับ order updates
func (ws *WebSocketClient) SubscribeOrderUpdate() error {
	return ws.connectToPath("/OnOrderUpdate", ws.handleOrderUpdateMessage)
}

// SubscribeOrderProfit subscribe รับ profit updates
func (ws *WebSocketClient) SubscribeOrderProfit() error {
	return ws.connectToPath("/OnOrderProfit", ws.handleOrderProfitMessage)
}

// SubscribeMarketWatch subscribe รับ market watch updates
func (ws *WebSocketClient) SubscribeMarketWatch() error {
	return ws.connectToPath("/OnMarketWatch", ws.handleMarketWatchMessage)
}

// SubscribeTickHistory subscribe รับ tick history
func (ws *WebSocketClient) SubscribeTickHistory() error {
	return ws.connectToPath("/OnTickHistory", ws.handleTickHistoryMessage)
}

// SubscribeMail subscribe รับ mail
func (ws *WebSocketClient) SubscribeMail() error {
	return ws.connectToPath("/OnMail", ws.handleMailMessage)
}

// SubscribeOpenedOrdersTickets subscribe รับ opened orders tickets
func (ws *WebSocketClient) SubscribeOpenedOrdersTickets(interval int) error {
	path := fmt.Sprintf("/OnOpenedOrdersTickets?interval=%d", interval)
	return ws.connectToPath(path, ws.handleOpenedOrdersTicketsMessage)
}

// SubscribeOrderBook subscribe รับ order book
func (ws *WebSocketClient) SubscribeOrderBook() error {
	return ws.connectToPath("/OnOrderBook", ws.handleOrderBookMessage)
}

// SubscribeOhlc subscribe รับ OHLC updates
func (ws *WebSocketClient) SubscribeOhlcWS() error {
	return ws.connectToPath("/OnOhlc", ws.handleOhlcMessage)
}

// readMessages อ่านข้อความจาก WebSocket
func (ws *WebSocketClient) readMessages(path string, conn *websocket.Conn, handler func([]byte)) {
	defer func() {
		if r := recover(); r != nil {
			log.Printf("WebSocket %s panic recovered: %v", path, r)
		}
		conn.Close()
		ws.mu.Lock()
		delete(ws.connections, path)
		ws.mu.Unlock()

		// Auto-reconnect ถ้าเปิดใช้งาน
		ws.mu.RLock()
		shouldReconnect := ws.autoReconnect && ws.isActive
		ws.mu.RUnlock()

		if shouldReconnect {
			go ws.reconnectPath(path, handler)
		}
	}()

	for {
		select {
		case <-ws.done:
			return
		default:
			_, message, err := conn.ReadMessage()
			if err != nil {
				if ws.handlers.OnError != nil {
					ws.handlers.OnError(fmt.Errorf("%s error: %v", path, err))
				}
				return
			}

			// ประมวลผลข้อความ
			handler(message)
		}
	}
}

// Message handlers
func (ws *WebSocketClient) handleQuoteMessage(data []byte) {
	if ws.handlers.OnQuote == nil {
		return
	}

	type RawQuote struct {
		Type  string `json:"type"`
		Quote Quote  `json:"data"`
	}
	var raw RawQuote
	if err := json.Unmarshal(data, &raw); err != nil {
		if ws.handlers.OnError != nil {
			ws.handlers.OnError(fmt.Errorf("failed to parse quote: %v", err))
		}
		return
	}

	ws.handlers.OnQuote(&raw.Quote)
}

func (ws *WebSocketClient) handleTickValueMessage(data []byte) {
	if ws.handlers.OnTickValue == nil {
		return
	}

	var event TickValueEvent
	if err := json.Unmarshal(data, &event); err != nil {
		if ws.handlers.OnError != nil {
			ws.handlers.OnError(fmt.Errorf("failed to parse tick value: %v", err))
		}
		return
	}

	ws.handlers.OnTickValue(&event)
}

func (ws *WebSocketClient) handleOrderUpdateMessage(data []byte) {
	if ws.handlers.OnOrderUpdate == nil {
		return
	}

	var raw SocketResponse
	if err := json.Unmarshal(data, &raw); err != nil {
		if ws.handlers.OnError != nil {
			ws.handlers.OnError(fmt.Errorf("failed to parse raw order update: %v", err))
		}
		return
	}
	switch raw.Type {
	case "OrderUpdate":
		var event OrderUpdateEvent
		if err := json.Unmarshal(raw.Data, &event); err != nil {
			if ws.handlers.OnError != nil {
				ws.handlers.OnError(fmt.Errorf("failed to parse raw order update: %v", err))
			}
			return
		}
		ws.handlers.OnOrderUpdate(&event)
	}

}

func (ws *WebSocketClient) handleOrderProfitMessage(data []byte) {
	if ws.handlers.OnOrderProfit == nil {
		return
	}

	var raw SocketResponse
	if err := json.Unmarshal(data, &raw); err != nil {
		if ws.handlers.OnError != nil {
			ws.handlers.OnError(fmt.Errorf("failed to parse raw order profit: %v", err))
		}
		return
	}

	switch raw.Type {
	case "OrderProfit":
		var event OrderProfitEvent
		if err := json.Unmarshal(raw.Data, &event); err != nil {
			if ws.handlers.OnError != nil {
				ws.handlers.OnError(fmt.Errorf("failed to parse order profit: %v", err))
			}
			return
		}
		ws.handlers.OnOrderProfit(&event)
	}

}

func (ws *WebSocketClient) handleMarketWatchMessage(data []byte) {
	if ws.handlers.OnMarketWatch == nil {
		return
	}

	var mw MarketWatch
	if err := json.Unmarshal(data, &mw); err != nil {
		if ws.handlers.OnError != nil {
			ws.handlers.OnError(fmt.Errorf("failed to parse market watch: %v", err))
		}
		return
	}

	ws.handlers.OnMarketWatch(&mw)
}

func (ws *WebSocketClient) handleTickHistoryMessage(data []byte) {
	if ws.handlers.OnTickHistory == nil {
		return
	}

	var event TickHistoryEvent
	if err := json.Unmarshal(data, &event); err != nil {
		if ws.handlers.OnError != nil {
			ws.handlers.OnError(fmt.Errorf("failed to parse tick history: %v", err))
		}
		return
	}

	ws.handlers.OnTickHistory(&event)
}

func (ws *WebSocketClient) handleMailMessage(data []byte) {
	if ws.handlers.OnMail == nil {
		return
	}

	var mail Mail
	if err := json.Unmarshal(data, &mail); err != nil {
		if ws.handlers.OnError != nil {
			ws.handlers.OnError(fmt.Errorf("failed to parse mail: %v", err))
		}
		return
	}

	ws.handlers.OnMail(&mail)
}

func (ws *WebSocketClient) handleOpenedOrdersTicketsMessage(data []byte) {
	if ws.handlers.OnOpenedOrdersTickets == nil {
		return
	}

	var tickets []int64
	if err := json.Unmarshal(data, &tickets); err != nil {
		if ws.handlers.OnError != nil {
			ws.handlers.OnError(fmt.Errorf("failed to parse opened orders tickets: %v", err))
		}
		return
	}

	ws.handlers.OnOpenedOrdersTickets(tickets)
}

func (ws *WebSocketClient) handleOrderBookMessage(data []byte) {
	if ws.handlers.OnOrderBook == nil {
		return
	}

	var book OrderBook
	if err := json.Unmarshal(data, &book); err != nil {
		if ws.handlers.OnError != nil {
			ws.handlers.OnError(fmt.Errorf("failed to parse order book: %v", err))
		}
		return
	}

	ws.handlers.OnOrderBook(&book)
}

func (ws *WebSocketClient) handleOhlcMessage(data []byte) {
	if ws.handlers.OnOhlc == nil {
		return
	}

	var ohlc OhlcData
	if err := json.Unmarshal(data, &ohlc); err != nil {
		if ws.handlers.OnError != nil {
			ws.handlers.OnError(fmt.Errorf("failed to parse ohlc: %v", err))
		}
		return
	}

	ws.handlers.OnOhlc(&ohlc)
}

// Disconnect ตัดการเชื่อมต่อทั้งหมด
func (ws *WebSocketClient) Disconnect() error {
	ws.mu.Lock()
	defer ws.mu.Unlock()

	if !ws.isActive {
		return nil
	}

	ws.isActive = false
	close(ws.done)

	// ปิด connections ทั้งหมด
	for path, conn := range ws.connections {
		conn.Close()
		delete(ws.connections, path)
	}

	// ล้าง subscribed paths
	ws.subscribedPaths = make(map[string]func([]byte))

	// เรียก OnDisconnect handler
	if ws.handlers.OnDisconnect != nil {
		go ws.handlers.OnDisconnect()
	}

	return nil
}

// IsConnected ตรวจสอบสถานะการเชื่อมต่อ
func (ws *WebSocketClient) IsConnected() bool {
	ws.mu.RLock()
	defer ws.mu.RUnlock()
	return ws.isActive
}

// KeepAlive ส่ง ping เป็นระยะ (สำหรับ connections ทั้งหมด)
func (ws *WebSocketClient) KeepAlive(interval time.Duration) {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	for {
		select {
		case <-ws.done:
			return
		case <-ticker.C:
			ws.mu.RLock()
			for _, conn := range ws.connections {
				conn.WriteMessage(websocket.PingMessage, []byte{})
			}
			ws.mu.RUnlock()
		}
	}
}
