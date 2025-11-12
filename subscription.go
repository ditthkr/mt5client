package mt5client

import "fmt"

// SubscriptionService จัดการ subscription
type SubscriptionService struct {
	client *Client
}

// Subscribe subscribe ราคาสัญลักษณ์
func (r *SubscriptionService) Subscribe(symbol string) error {
	if r.client.token == "" {
		return fmt.Errorf("not connected")
	}

	queryParams := map[string]string{
		"id":     r.client.token,
		"symbol": symbol,
	}

	var result string
	err := r.client.get("/Subscribe", queryParams, &result)
	return err
}

// SubscribeMany subscribe หลายสัญลักษณ์
func (r *SubscriptionService) SubscribeMany(symbols []string) error {
	if r.client.token == "" {
		return fmt.Errorf("not connected")
	}

	queryParams := map[string]string{
		"id": r.client.token,
	}

	for i, symbol := range symbols {
		queryParams[fmt.Sprintf("symbols[%d]", i)] = symbol
	}

	var result string
	err := r.client.get("/SubscribeMany", queryParams, &result)
	return err
}

// Unsubscribe unsubscribe สัญลักษณ์
func (r *SubscriptionService) Unsubscribe(symbol string) error {
	if r.client.token == "" {
		return fmt.Errorf("not connected")
	}

	queryParams := map[string]string{
		"id":     r.client.token,
		"symbol": symbol,
	}

	var result string
	err := r.client.get("/UnSubscribe", queryParams, &result)
	return err
}

// UnsubscribeMany unsubscribe หลายสัญลักษณ์
func (r *SubscriptionService) UnsubscribeMany(symbols []string) error {
	if r.client.token == "" {
		return fmt.Errorf("not connected")
	}

	queryParams := map[string]string{
		"id": r.client.token,
	}

	for i, symbol := range symbols {
		queryParams[fmt.Sprintf("symbols[%d]", i)] = symbol
	}

	var result string
	err := r.client.get("/UnSubscribeMany", queryParams, &result)
	return err
}

// SubscribeTickValue subscribe tick value
func (r *SubscriptionService) SubscribeTickValue(symbol string) error {
	if r.client.token == "" {
		return fmt.Errorf("not connected")
	}

	queryParams := map[string]string{
		"id":     r.client.token,
		"symbol": symbol,
	}

	var result string
	err := r.client.get("/SubscribeTickValue", queryParams, &result)
	return err
}

// SubscribeOhlc subscribe OHLC
func (r *SubscriptionService) SubscribeOhlc(symbol, timeframe string) error {
	if r.client.token == "" {
		return fmt.Errorf("not connected")
	}

	queryParams := map[string]string{
		"id":        r.client.token,
		"symbol":    symbol,
		"timeframe": timeframe,
	}

	var result string
	err := r.client.get("/SubscribeOhlc", queryParams, &result)
	return err
}

// UnsubscribeOhlc unsubscribe OHLC
func (r *SubscriptionService) UnsubscribeOhlc(symbol, timeframe string) error {
	if r.client.token == "" {
		return fmt.Errorf("not connected")
	}

	queryParams := map[string]string{
		"id":        r.client.token,
		"symbol":    symbol,
		"timeframe": timeframe,
	}

	var result string
	err := r.client.get("/UnsubscribeOhlc", queryParams, &result)
	return err
}

// SubscribeOrderBook subscribe order book
func (r *SubscriptionService) SubscribeOrderBook(symbol string) error {
	if r.client.token == "" {
		return fmt.Errorf("not connected")
	}

	queryParams := map[string]string{
		"id":     r.client.token,
		"symbol": symbol,
	}

	var result string
	err := r.client.get("/SubscribeOrderBook", queryParams, &result)
	return err
}

// UnsubscribeOrderBook unsubscribe order book
func (r *SubscriptionService) UnsubscribeOrderBook(symbol string) error {
	if r.client.token == "" {
		return fmt.Errorf("not connected")
	}

	queryParams := map[string]string{
		"id":     r.client.token,
		"symbol": symbol,
	}

	var result string
	err := r.client.get("/UnsubscribeOrderBook", queryParams, &result)
	return err
}

// SubscribeMarketWatch subscribe market watch
func (r *SubscriptionService) SubscribeMarketWatch() error {
	if r.client.token == "" {
		return fmt.Errorf("not connected")
	}

	queryParams := map[string]string{
		"id": r.client.token,
	}

	var result string
	err := r.client.get("/SubscribeMarketWatch", queryParams, &result)
	return err
}

// GetWebSocketURL ดึง WebSocket URL สำหรับ events
func (r *SubscriptionService) GetWebSocketURL() (string, error) {
	// WebSocket endpoints จะใช้ผ่าน /On* paths
	// เช่น /OnQuote, /OnOrderUpdate, etc.
	// ต้องใช้ WebSocket client แยกต่างหาก
	return r.client.baseURL, nil
}
