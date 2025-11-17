package mt5client

import "fmt"

// ServiceFunctions ฟังก์ชันบริการทั่วไป
type ServiceFunctions struct {
	client *Client
}

// GetVersion ดึงเวอร์ชันของ API
func (r *ServiceFunctions) GetVersion() (string, error) {
	var version string
	err := r.client.get("/Version", nil, &version)
	return version, err
}

// Ping ทดสอบการเชื่อมต่อ
func (r *ServiceFunctions) Ping() (string, error) {
	var result string
	err := r.client.get("/Ping", nil, &result)
	return result, err
}

// PingHost ทดสอบการเชื่อมต่อไปยัง host
func (r *ServiceFunctions) PingHost(host string) (bool, error) {
	queryParams := map[string]string{
		"host": host,
	}

	var result bool
	err := r.client.get("/PingHost", queryParams, &result)
	return result, err
}

// PingHostMany ทดสอบหลาย hosts
func (r *ServiceFunctions) PingHostMany(hosts []string) (map[string]bool, error) {
	queryParams := map[string]string{}

	for i, host := range hosts {
		queryParams[fmt.Sprintf("hosts[%d]", i)] = host
	}

	var results map[string]bool
	err := r.client.get("/PingHostMany", queryParams, &results)
	return results, err
}

// Search ค้นหาสัญลักษณ์
func (r *ServiceFunctions) Search(keyword string) ([]string, error) {
	if r.client.token == "" {
		return nil, fmt.Errorf("not connected")
	}

	queryParams := map[string]string{
		"id":      r.client.token,
		"keyword": keyword,
	}

	var results []string
	err := r.client.get("/Search", queryParams, &results)
	return results, err
}

// GetServerTimezone ดึงข้อมูล timezone ของ server
func (r *ServiceFunctions) GetServerTimezone() (string, error) {
	if r.client.token == "" {
		return "", fmt.Errorf("not connected")
	}

	queryParams := map[string]string{
		"id": r.client.token,
	}

	var timezone string
	err := r.client.get("/ServerTimezone", queryParams, &timezone)
	return timezone, err
}

// GetClusterDetails ดึงข้อมูล cluster
func (r *ServiceFunctions) GetClusterDetails() (map[string]interface{}, error) {
	if r.client.token == "" {
		return nil, fmt.Errorf("not connected")
	}

	queryParams := map[string]string{
		"id": r.client.token,
	}

	var details map[string]interface{}
	err := r.client.get("/ClusterDetails", queryParams, &details)
	return details, err
}

// ChangePassword เปลี่ยนรหัสผ่าน
func (r *ServiceFunctions) ChangePassword(oldPassword, newPassword string) error {
	if r.client.token == "" {
		return fmt.Errorf("not connected")
	}

	queryParams := map[string]string{
		"id":          r.client.token,
		"oldPassword": oldPassword,
		"newPassword": newPassword,
	}

	var result string
	err := r.client.get("/ChangePassword", queryParams, &result)
	return err
}

// GetDemo ขอบัญชี demo
func (r *ServiceFunctions) GetDemo(server, name, email string) (map[string]interface{}, error) {
	queryParams := map[string]string{
		"server": server,
		"name":   name,
		"email":  email,
	}

	var result map[string]interface{}
	err := r.client.get("/GetDemo", queryParams, &result)
	return result, err
}

// GetRequiredMargin คำนวณ margin ที่ต้องการ
func (r *ServiceFunctions) GetRequiredMargin(symbol string, volume float64) (float64, error) {
	if r.client.token == "" {
		return 0, fmt.Errorf("not connected")
	}

	queryParams := map[string]string{
		"id":     r.client.token,
		"symbol": symbol,
		"volume": fmt.Sprintf("%.2f", volume),
	}

	var margin float64
	err := r.client.get("/RequiredMargin", queryParams, &margin)
	return margin, err
}

// GetMails ดึงอีเมล
func (r *ServiceFunctions) GetMails() ([]Mail, error) {
	if r.client.token == "" {
		return nil, fmt.Errorf("not connected")
	}

	queryParams := map[string]string{
		"id": r.client.token,
	}

	var mails []Mail
	err := r.client.get("/Mails", queryParams, &mails)
	return mails, err
}

// GetMarketWatchMany ดึง market watch หลายสัญลักษณ์
func (r *ServiceFunctions) GetMarketWatchMany(symbols []string) ([]MarketWatch, error) {
	if r.client.token == "" {
		return nil, fmt.Errorf("not connected")
	}

	queryParams := map[string]string{
		"id": r.client.token,
	}

	for i, symbol := range symbols {
		queryParams[fmt.Sprintf("symbols[%d]", i)] = symbol
	}

	var marketWatch []MarketWatch
	err := r.client.get("/MarketWatchMany", queryParams, &marketWatch)
	return marketWatch, err
}

// GetQuoteClient ดึงข้อมูล quote client
func (r *ServiceFunctions) GetQuoteClient() (map[string]interface{}, error) {
	if r.client.token == "" {
		return nil, fmt.Errorf("not connected")
	}

	queryParams := map[string]string{
		"id": r.client.token,
	}

	var client map[string]interface{}
	err := r.client.get("/QuoteClient", queryParams, &client)
	return client, err
}

// LoadServersDat โหลดไฟล์ servers.dat
func (r *ServiceFunctions) LoadServersDat(data []byte) error {
	var result string
	err := r.client.post("/LoadServersDat", nil, data, &result)
	return err
}

// GetMetricsApiKey ดึง API key สำหรับ metrics
func (r *ServiceFunctions) GetMetricsApiKey() (string, error) {
	if r.client.token == "" {
		return "", fmt.Errorf("not connected")
	}

	queryParams := map[string]string{
		"id": r.client.token,
	}

	var apiKey string
	err := r.client.get("/MetricsApiKey", queryParams, &apiKey)
	return apiKey, err
}

// GetReadMe ดึง README
func (r *ServiceFunctions) GetReadMe() (string, error) {
	var readme string
	err := r.client.get("/ReadMe", nil, &readme)
	return readme, err
}

// LotSizeParams พารามิเตอร์สำหรับคำนวณ lot size
type LotSizeParams struct {
	Symbol      string  // สัญลักษณ์ เช่น EURUSD
	EntryPrice  float64 // ราคาเข้า
	StopLoss    float64 // ราคา Stop Loss
	RiskAmount  float64 // ยอมขาดทุนเป็นเงิน (เช่น 100 = $100)
	RiskPercent float64 // ยอมขาดทุนเป็น % (เช่น 2.0 = 2% ของพอร์ต)
}

// LotSizeResult ผลลัพธ์การคำนวณ lot size
type LotSizeResult struct {
	LotSize       float64 `json:"lotSize"`       // ขนาด lot ที่คำนวณได้
	RiskAmount    float64 `json:"riskAmount"`    // จำนวนเงินที่เสี่ยง
	EntryPrice    float64 `json:"entryPrice"`    // ราคา entry ที่ใช้คำนวณ (ถ้าส่งมาเป็น 0 จะเป็นราคา market)
	PriceDistance float64 `json:"priceDistance"` // ระยะห่าง entry ถึง SL (เป็น price)
	PointDistance float64 `json:"pointDistance"` // ระยะห่างเป็น points
	TickValue     float64 `json:"tickValue"`     // มูลค่าต่อ tick
	Symbol        string  `json:"symbol"`        // สัญลักษณ์
}

// CalculateLotSize คำนวณ lot size จาก risk amount (เงิน)
// สูตร: Lot Size = Risk Amount / (Point Distance × Tick Value)
func (r *ServiceFunctions) CalculateLotSize(symbol string, entryPrice, stopLoss, riskAmount float64) (*LotSizeResult, error) {
	if r.client.token == "" {
		return nil, fmt.Errorf("not connected")
	}

	if stopLoss <= 0 {
		return nil, fmt.Errorf("stop loss must be greater than 0")
	}

	if riskAmount <= 0 {
		return nil, fmt.Errorf("risk amount must be greater than 0")
	}

	// ถ้า entryPrice = 0 ให้ใช้ราคา market ปัจจุบัน
	if entryPrice == 0 {
		quote, err := r.client.Quote.Get(symbol)
		if err != nil {
			return nil, fmt.Errorf("failed to get current price: %w", err)
		}

		// คำนวณ mid price
		midPrice := (quote.Bid + quote.Ask) / 2

		// ตัดสินใจว่าเป็น BUY หรือ SELL จาก stopLoss
		if stopLoss < midPrice {
			// BUY order: SL ต่ำกว่า market → ซื้อที่ Ask
			entryPrice = quote.Ask
		} else if stopLoss > midPrice {
			// SELL order: SL สูงกว่า market → ขายที่ Bid
			entryPrice = quote.Bid
		} else {
			// SL = midPrice (ไม่น่าเกิด แต่ป้องกันไว้)
			return nil, fmt.Errorf("stop loss cannot be equal to current market price")
		}
	} else if entryPrice < 0 {
		return nil, fmt.Errorf("entry price must be greater than or equal to 0 (use 0 for market price)")
	}

	// ดึงข้อมูล Symbol เพื่อได้ Point value
	symbolParams, err := r.client.Symbol.GetParams(symbol)
	if err != nil {
		return nil, fmt.Errorf("failed to get symbol info: %w", err)
	}

	if symbolParams.SymbolInfo.Points <= 0 {
		return nil, fmt.Errorf("invalid point value for symbol %s", symbol)
	}

	// คำนวณ Tick Value ถ้า API ส่งมาเป็น 0
	tickValue := symbolParams.SymbolInfo.TickValue
	if tickValue <= 0 {
		tickValue, err = r.calculateTickValue(&symbolParams.SymbolInfo, symbol)
		if err != nil {
			return nil, fmt.Errorf("failed to calculate tick value: %w", err)
		}
	}

	// คำนวณระยะห่างราคา
	priceDistance := entryPrice - stopLoss
	if priceDistance < 0 {
		priceDistance = -priceDistance
	}

	// คำนวณระยะห่างเป็น points
	pointDistance := priceDistance / symbolParams.SymbolInfo.Points

	// คำนวณ lot size
	// สูตร: Lot = Risk / (Points × TickValue)
	lotSize := riskAmount / (pointDistance * tickValue)

	// ปัดเศษให้เป็นทศนิยม 2 ตำแหน่ง (standard lot size)
	lotSize = float64(int(lotSize*100)) / 100

	// ตรวจสอบขอบเขต lot size
	if symbolParams.SymbolGroup.MinLots > 0 && lotSize < symbolParams.SymbolGroup.MinLots {
		lotSize = symbolParams.SymbolGroup.MinLots
	}
	if symbolParams.SymbolGroup.MaxLots > 0 && lotSize > symbolParams.SymbolGroup.MaxLots {
		return nil, fmt.Errorf("calculated lot size (%.2f) exceeds maximum (%.2f)", lotSize, symbolParams.SymbolGroup.MaxLots)
	}

	return &LotSizeResult{
		LotSize:       lotSize,
		RiskAmount:    riskAmount,
		EntryPrice:    entryPrice,
		PriceDistance: priceDistance,
		PointDistance: pointDistance,
		TickValue:     tickValue,
		Symbol:        symbol,
	}, nil
}

// CalculateLotSizeByPercent คำนวณ lot size จาก risk % ของพอร์ต
func (r *ServiceFunctions) CalculateLotSizeByPercent(symbol string, entryPrice, stopLoss, riskPercent float64) (*LotSizeResult, error) {
	if r.client.token == "" {
		return nil, fmt.Errorf("not connected")
	}

	if riskPercent <= 0 || riskPercent > 100 {
		return nil, fmt.Errorf("risk percent must be between 0 and 100")
	}

	// ดึงข้อมูลบัญชีเพื่อได้ Balance
	account, err := r.client.Account.GetInfo()
	if err != nil {
		return nil, fmt.Errorf("failed to get account info: %w", err)
	}

	// คำนวณ risk amount จาก % ของ balance
	riskAmount := account.Balance * (riskPercent / 100.0)

	// เรียกใช้ CalculateLotSize
	return r.CalculateLotSize(symbol, entryPrice, stopLoss, riskAmount)
}

// CalculatePipValue คำนวณมูลค่าต่อ pip สำหรับ lot size ที่กำหนด
func (r *ServiceFunctions) CalculatePipValue(symbol string, lotSize float64) (float64, error) {
	if r.client.token == "" {
		return 0, fmt.Errorf("not connected")
	}

	if lotSize <= 0 {
		return 0, fmt.Errorf("lot size must be greater than 0")
	}

	// ดึงข้อมูล Symbol
	symbolParams, err := r.client.Symbol.GetParams(symbol)
	if err != nil {
		return 0, fmt.Errorf("failed to get symbol info: %w", err)
	}

	// คำนวณ Tick Value ถ้า API ส่งมาเป็น 0
	tickValue := symbolParams.SymbolInfo.TickValue
	if tickValue <= 0 {
		tickValue, err = r.calculateTickValue(&symbolParams.SymbolInfo, symbol)
		if err != nil {
			return 0, fmt.Errorf("failed to calculate tick value: %w", err)
		}
	}

	// สำหรับ forex ที่มี 5 digits, 1 pip = 10 points
	// สำหรับ forex ที่มี 3 digits, 1 pip = 1 point
	pipsPerPoint := 1.0
	if symbolParams.SymbolInfo.Digits == 5 || symbolParams.SymbolInfo.Digits == 3 {
		pipsPerPoint = 10.0
	}

	// Pip Value = Tick Value × Lot Size × pipsPerPoint
	pipValue := tickValue * lotSize * pipsPerPoint

	return pipValue, nil
}

// calculateTickValue คำนวณ Tick Value เมื่อ API ส่งมาเป็น 0
func (r *ServiceFunctions) calculateTickValue(symbolInfo *SymbolInfo, symbol string) (float64, error) {
	// ตรวจสอบว่ามีข้อมูลที่จำเป็นหรือไม่
	if symbolInfo.Points <= 0 {
		return 0, fmt.Errorf("invalid points value for %s", symbol)
	}
	if symbolInfo.ContractSize <= 0 {
		return 0, fmt.Errorf("invalid contract size for %s", symbol)
	}

	// กรณีที่ 1: ProfitCurrency เป็น USD (เช่น EURUSD, GBPUSD)
	// TickValue = Points × ContractSize
	if symbolInfo.ProfitCurrency == "USD" {
		tickValue := symbolInfo.Points * symbolInfo.ContractSize
		return tickValue, nil
	}

	// กรณีที่ 2: ProfitCurrency เป็น JPY (เช่น USDJPY, EURJPY)
	// ต้องดึง current price มาใช้คำนวณ
	if symbolInfo.ProfitCurrency == "JPY" {
		// ดึง current price
		quote, err := r.client.Quote.Get(symbol)
		if err != nil {
			return 0, fmt.Errorf("failed to get quote for %s: %w", symbol, err)
		}

		// ใช้ mid price (เฉลี่ย bid/ask)
		currentPrice := (quote.Bid + quote.Ask) / 2
		if currentPrice <= 0 {
			return 0, fmt.Errorf("invalid quote price for %s", symbol)
		}

		// TickValue = (Points × ContractSize) / CurrentPrice
		tickValue := (symbolInfo.Points * symbolInfo.ContractSize) / currentPrice
		return tickValue, nil
	}

	// กรณีที่ 3: ProfitCurrency เป็นสกุลอื่นๆ (Cross pairs)
	// ต้องแปลงผ่าน conversion rate จาก ProfitCurrency -> USD

	// ตรวจสอบว่ามี suffix หรือไม่ (เช่น AUDNZDm -> m, EURUSD.raw -> .raw)
	suffix := ""
	baseCurrencyPairLen := 6 // ความยาวปกติของ currency pair (เช่น EURUSD)
	if len(symbol) > baseCurrencyPairLen {
		suffix = symbol[baseCurrencyPairLen:]
	}

	// หา conversion pair (เช่น NZD -> NZDUSD, EUR -> EURUSD) พร้อม suffix
	conversionSymbol := symbolInfo.ProfitCurrency + "USD" + suffix

	// ลองดึง quote ของ conversion pair
	conversionQuote, err := r.client.Quote.Get(conversionSymbol)
	if err != nil {
		// ถ้าไม่เจอ เช่น NZDUSD อาจจะต้องลอง USDNZD แทน
		reverseSymbol := "USD" + symbolInfo.ProfitCurrency + suffix
		conversionQuote, err = r.client.Quote.Get(reverseSymbol)
		if err != nil {
			return 0, fmt.Errorf("unsupported profit currency %s for symbol %s, cannot find conversion pair %s or %s",
				symbolInfo.ProfitCurrency, symbol, conversionSymbol, reverseSymbol)
		}

		// ถ้าใช้ reverse pair (USDNZD) ต้องใช้ 1/price
		conversionRate := (conversionQuote.Bid + conversionQuote.Ask) / 2
		if conversionRate <= 0 {
			return 0, fmt.Errorf("invalid conversion rate for %s", reverseSymbol)
		}

		// TickValue = (Points × ContractSize) / ConversionRate
		tickValue := (symbolInfo.Points * symbolInfo.ContractSize) / conversionRate
		return tickValue, nil
	}

	// ถ้าเจอ direct pair (เช่น NZDUSD, EURUSD)
	conversionRate := (conversionQuote.Bid + conversionQuote.Ask) / 2
	if conversionRate <= 0 {
		return 0, fmt.Errorf("invalid conversion rate for %s", conversionSymbol)
	}

	// TickValue = (Points × ContractSize) × ConversionRate
	tickValue := symbolInfo.Points * symbolInfo.ContractSize * conversionRate
	return tickValue, nil
}
