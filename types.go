package mt5client

import (
	"encoding/json"
	"time"
)

// ConnectParams พารามิเตอร์สำหรับการเชื่อมต่อ
type ConnectParams struct {
	User     int64  `json:"user"`
	Password string `json:"password"`
	Host     string `json:"host"`
	Port     int    `json:"port"`
}

// Account ข้อมูลบัญชี
type Account struct {
	Login       int64   `json:"login"`
	Balance     float64 `json:"balance"`
	Equity      float64 `json:"equity"`
	Profit      float64 `json:"profit"`
	Margin      float64 `json:"margin"`
	MarginFree  float64 `json:"marginFree"`
	MarginLevel float64 `json:"marginLevel"`
	Credit      float64 `json:"credit"`
	Currency    string  `json:"currency"`
	Company     string  `json:"company"`
	Name        string  `json:"name"`
	Server      string  `json:"server"`
	Leverage    int     `json:"leverage"`
}

// Order คำสั่งซื้อขาย
type Order struct {
	Ticket            uint64    `json:"ticket"`
	Profit            float64   `json:"profit"`
	OrderType         string    `json:"orderType"`
	Symbol            string    `json:"symbol"`
	Lots              float64   `gorm:"not null"`
	OpenPrice         float64   `json:"openPrice"`
	OpenTime          time.Time `json:"openTime"`
	OpenTimestampUTC  int64     `json:"openTimestampUTC"`
	Volume            uint64    `json:"volume"`
	ClosePrice        float64   `json:"closePrice"`
	CloseTime         time.Time `json:"closeTime"`
	CloseTimestampUTC int64     `json:"closeTimestampUTC"`
	CloseVolume       uint64    `json:"closeVolume"`
	TakeProfit        float64   `json:"takeProfit"`
	StopLoss          float64   `json:"stopLoss"`
	Swap              float64   `json:"swap"`
	Commission        float64   `json:"commission"`
	Fee               float64   `json:"fee"`
	State             string    `json:"state"`
	Comment           string    `json:"comment"`
}

// UnmarshalJSON custom unmarshal สำหรับ Order
func (o *Order) UnmarshalJSON(data []byte) error {
	type Alias Order
	aux := &struct {
		OpenTime  string `json:"openTime"`
		CloseTime string `json:"closeTime"`
		*Alias
	}{
		Alias: (*Alias)(o),
	}

	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}

	if aux.OpenTime != "" {
		t, err := parseTime(aux.OpenTime)
		if err == nil {
			o.OpenTime = t
		}
	}

	return nil
}

// OrderRequest พารามิเตอร์สำหรับส่งคำสั่ง
type OrderRequest struct {
	Symbol      string  `json:"symbol"`
	Type        string  `json:"type"`
	Volume      float64 `json:"volume"`
	Price       float64 `json:"price,omitempty"`
	StopLoss    float64 `json:"stopLoss,omitempty"`
	TakeProfit  float64 `json:"takeProfit,omitempty"`
	Comment     string  `json:"comment,omitempty"`
	MagicNumber int64   `json:"magicNumber,omitempty"`
}

// TradeResult ผลลัพธ์การเทรด
type TradeResult struct {
	Success bool    `json:"success"`
	Order   int64   `json:"order"`
	Deal    int64   `json:"deal"`
	Volume  float64 `json:"volume"`
	Price   float64 `json:"price"`
	Bid     float64 `json:"bid"`
	Ask     float64 `json:"ask"`
	Comment string  `json:"comment"`
	Request string  `json:"request"`
}

// Quote ราคา
type Quote struct {
	Symbol    string    `json:"symbol"`
	Bid       float64   `json:"bid"`
	Ask       float64   `json:"ask"`
	Last      float64   `json:"last"`
	Volume    float64   `json:"volume"`
	Time      time.Time `json:"time"`
	Spread    int       `json:"spread"`
	TickValue float64   `json:"tickValue"`
}

// UnmarshalJSON custom unmarshal สำหรับ Quote
func (q *Quote) UnmarshalJSON(data []byte) error {
	type Alias Quote
	aux := &struct {
		Time string `json:"time"`
		*Alias
	}{
		Alias: (*Alias)(q),
	}

	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}

	if aux.Time != "" {
		t, err := parseTime(aux.Time)
		if err == nil {
			q.Time = t
		}
	}

	return nil
}

// Symbol ข้อมูลสัญลักษณ์
type Symbol struct {
	Name              string  `json:"name"`
	Description       string  `json:"description"`
	Path              string  `json:"path"`
	Digits            int     `json:"digits"`
	Point             float64 `json:"point"`
	Spread            int     `json:"spread"`
	StopsLevel        int     `json:"stopsLevel"`
	TradeMode         string  `json:"tradeMode"`
	VolumeMin         float64 `json:"volumeMin"`
	VolumeMax         float64 `json:"volumeMax"`
	VolumeStep        float64 `json:"volumeStep"`
	ContractSize      float64 `json:"contractSize"`
	CurrencyBase      string  `json:"currencyBase"`
	CurrencyProfit    string  `json:"currencyProfit"`
	CurrencyMargin    string  `json:"currencyMargin"`
	SwapLong          float64 `json:"swapLong"`
	SwapShort         float64 `json:"swapShort"`
	QuoteSessionStart string  `json:"quoteSessionStart"`
	QuoteSessionEnd   string  `json:"quoteSessionEnd"`
}

// Bar แท่งเทียน
type Bar struct {
	Time   time.Time `json:"time"`
	Open   float64   `json:"open"`
	High   float64   `json:"high"`
	Low    float64   `json:"low"`
	Close  float64   `json:"close"`
	Volume int64     `json:"volume"`
	Spread int       `json:"spread"`
}

// UnmarshalJSON custom unmarshal สำหรับ Bar
func (b *Bar) UnmarshalJSON(data []byte) error {
	type Alias Bar
	aux := &struct {
		Time string `json:"time"`
		*Alias
	}{
		Alias: (*Alias)(b),
	}

	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}

	if aux.Time != "" {
		t, err := parseTime(aux.Time)
		if err == nil {
			b.Time = t
		}
	}

	return nil
}

// Tick ข้อมูล tick
type Tick struct {
	Time   time.Time `json:"time"`
	Bid    float64   `json:"bid"`
	Ask    float64   `json:"ask"`
	Last   float64   `json:"last"`
	Volume float64   `json:"volume"`
	Flags  int       `json:"flags"`
}

// UnmarshalJSON custom unmarshal สำหรับ Tick
func (t *Tick) UnmarshalJSON(data []byte) error {
	type Alias Tick
	aux := &struct {
		Time string `json:"time"`
		*Alias
	}{
		Alias: (*Alias)(t),
	}

	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}

	if aux.Time != "" {
		tm, err := parseTime(aux.Time)
		if err == nil {
			t.Time = tm
		}
	}

	return nil
}

// TradeSummary สรุปการเทรด
type TradeSummary struct {
	OpenTrades  int     `json:"openTrades"`
	OpenProfit  float64 `json:"openProfit"`
	DayProfit   float64 `json:"dayProfit"`
	WeekProfit  float64 `json:"weekProfit"`
	MonthProfit float64 `json:"monthProfit"`
	TotalProfit float64 `json:"totalProfit"`
}

// MarketTradeCount จำนวนเทรดแยกตาม symbol
type MarketTradeCount struct {
	MarketName string `json:"marketName"`
	Count      int    `json:"count"`
}

// Profitability ข้อมูลความสามารถทำกำไร
type Profitability struct {
	WonTrades         int64   `json:"wonTrades"`
	WonTradesPercent  float64 `json:"wonTradesPercent"`
	LostTrades        int64   `json:"lostTrades"`
	LostTradesPercent float64 `json:"lostTradesPercent"`
}

// AveragePipsUsd ค่าเฉลี่ย pips และ USD
type AveragePipsUsd struct {
	AveragePips float64 `json:"averagePips"`
	AverageUsd  float64 `json:"averageUsd"`
}

// Won ข้อมูลการชนะ
type Won struct {
	WonCount   int     `json:"wonCount"`
	All        int     `json:"all"`
	WonPercent float64 `json:"wonPersent"` // typo ใน swagger
}

// ProfitData ข้อมูลกำไร
type ProfitData struct {
	Ticket int64     `json:"tiket"` // typo ใน swagger
	Date   time.Time `json:"date"`
	Profit float64   `json:"profit"`
}

// UnmarshalJSON custom unmarshal สำหรับ ProfitData
func (p *ProfitData) UnmarshalJSON(data []byte) error {
	type Alias ProfitData
	aux := &struct {
		Date string `json:"date"`
		*Alias
	}{
		Alias: (*Alias)(p),
	}

	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}

	if aux.Date != "" {
		t, err := parseTime(aux.Date)
		if err == nil {
			p.Date = t
		}
	}

	return nil
}

// ZScore Z-Score
type ZScore struct {
	ZScoreDecimal     float64 `json:"zScoreDecimal"`
	ZScoreProbability float64 `json:"zScoreProbability"`
}

// Expectancy Expectancy
type Expectancy struct {
	Pips   float64 `json:"pips"`
	Dollar float64 `json:"dollar"`
}

// TradeStats สถิติการเทรดแบบละเอียด
type TradeStats struct {
	Summary                    TradeSummary       `json:"summary"`
	MaxBalanceDrawdownRaw      float64            `json:"maxBalanceDrawdownRaw"`
	MaxBalanceDrawdownRelative float64            `json:"maxBalanceDrawdownRelative"`
	MaxEquityDrawdownRaw       float64            `json:"maxEquityDrawdownRaw"`
	MaxEquityDrawdownRelative  float64            `json:"maxEquityDrawdownRelative"`
	Markets                    []MarketTradeCount `json:"markets"`
	Profitability              Profitability      `json:"profitability"`
	Pips                       float64            `json:"pips"`
	Lots                       float64            `json:"lots"`
	Commissions                float64            `json:"comissions"` // typo ใน swagger
	AverageWin                 AveragePipsUsd     `json:"averageWin"`
	AverageLost                AveragePipsUsd     `json:"averageLost"`
	LongsWon                   Won                `json:"longsWon"`
	ShortsWon                  Won                `json:"shortsWon"`
	BestTrade                  ProfitData         `json:"bestTrade"`
	WorstTrade                 ProfitData         `json:"worstTrade"`
	BestTradePips              ProfitData         `json:"bestTradePips"`
	WorstTradePips             ProfitData         `json:"worstTradePips"`
	AverageTradeLength         string             `json:"averageTradeLength"` // format: date-span
	ProfitFactor               float64            `json:"profitFactor"`
	StandardDeviation          float64            `json:"standardDeviation"`
	SharpeRatio                float64            `json:"sharpeRatio"`
	ZScore                     ZScore             `json:"zScore"`
	Expectancy                 Expectancy         `json:"expectancy"`
	GHPR                       float64            `json:"ghpr"`
	Trades                     int64              `json:"trades"`
}

// Position ตำแหน่ง (alias ของ Order สำหรับความเข้ากันได้)
type Position = Order

// Deal ดีล
type Deal struct {
	Ticket      int64     `json:"ticket"`
	Order       int64     `json:"order"`
	Time        time.Time `json:"time"`
	Type        string    `json:"type"`
	Entry       string    `json:"entry"`
	Symbol      string    `json:"symbol"`
	Volume      float64   `json:"volume"`
	Price       float64   `json:"price"`
	Commission  float64   `json:"commission"`
	Swap        float64   `json:"swap"`
	Profit      float64   `json:"profit"`
	Fee         float64   `json:"fee"`
	Comment     string    `json:"comment"`
	MagicNumber int64     `json:"magicNumber"`
}

// UnmarshalJSON custom unmarshal สำหรับ Deal
func (d *Deal) UnmarshalJSON(data []byte) error {
	type Alias Deal
	aux := &struct {
		Time string `json:"time"`
		*Alias
	}{
		Alias: (*Alias)(d),
	}

	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}

	if aux.Time != "" {
		t, err := parseTime(aux.Time)
		if err == nil {
			d.Time = t
		}
	}

	return nil
}

// HistoryPosition ตำแหน่งที่ปิดแล้ว
type HistoryPosition struct {
	PositionId  int64     `json:"positionId"`
	Symbol      string    `json:"symbol"`
	OpenTime    time.Time `json:"openTime"`
	CloseTime   time.Time `json:"closeTime"`
	Volume      float64   `json:"volume"`
	OpenPrice   float64   `json:"openPrice"`
	ClosePrice  float64   `json:"closePrice"`
	Profit      float64   `json:"profit"`
	Commission  float64   `json:"commission"`
	Swap        float64   `json:"swap"`
	Comment     string    `json:"comment"`
	MagicNumber int64     `json:"magicNumber"`
}

// UnmarshalJSON custom unmarshal สำหรับ HistoryPosition
func (h *HistoryPosition) UnmarshalJSON(data []byte) error {
	type Alias HistoryPosition
	aux := &struct {
		OpenTime  string `json:"openTime"`
		CloseTime string `json:"closeTime"`
		*Alias
	}{
		Alias: (*Alias)(h),
	}

	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}

	if aux.OpenTime != "" {
		t, err := parseTime(aux.OpenTime)
		if err == nil {
			h.OpenTime = t
		}
	}

	if aux.CloseTime != "" {
		t, err := parseTime(aux.CloseTime)
		if err == nil {
			h.CloseTime = t
		}
	}

	return nil
}

// MarketWatch ข้อมูล Market Watch
type MarketWatch struct {
	Symbol string  `json:"symbol"`
	Bid    float64 `json:"bid"`
	Ask    float64 `json:"ask"`
	Time   string  `json:"time"`
}

// Mail อีเมล
type Mail struct {
	Time    time.Time `json:"time"`
	Subject string    `json:"subject"`
	Sender  string    `json:"sender"`
	Body    string    `json:"body"`
}

// UnmarshalJSON custom unmarshal สำหรับ Mail
func (m *Mail) UnmarshalJSON(data []byte) error {
	type Alias Mail
	aux := &struct {
		Time string `json:"time"`
		*Alias
	}{
		Alias: (*Alias)(m),
	}

	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}

	if aux.Time != "" {
		t, err := parseTime(aux.Time)
		if err == nil {
			m.Time = t
		}
	}

	return nil
}

// OrderBook Order Book
type OrderBook struct {
	Symbol string           `json:"symbol"`
	Time   time.Time        `json:"time"`
	Bids   []OrderBookLevel `json:"bids"`
	Asks   []OrderBookLevel `json:"asks"`
}

// UnmarshalJSON custom unmarshal สำหรับ OrderBook
func (o *OrderBook) UnmarshalJSON(data []byte) error {
	type Alias OrderBook
	aux := &struct {
		Time string `json:"time"`
		*Alias
	}{
		Alias: (*Alias)(o),
	}

	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}

	if aux.Time != "" {
		t, err := parseTime(aux.Time)
		if err == nil {
			o.Time = t
		}
	}

	return nil
}

// OrderBookLevel ระดับราคาใน Order Book
type OrderBookLevel struct {
	Price  float64 `json:"price"`
	Volume float64 `json:"volume"`
}

// OhlcData ข้อมูล OHLC
type OhlcData struct {
	Symbol    string `json:"symbol"`
	Timeframe string `json:"timeframe"`
	Bar       Bar    `json:"bar"`
}
