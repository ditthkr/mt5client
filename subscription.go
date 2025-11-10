package mt5client

import "fmt"

// SubscriptionService จัดการ subscription
type SubscriptionService struct {
	client *Client
}

// Subscribe subscribe ราคาสัญลักษณ์
func (s *SubscriptionService) Subscribe(symbol string) error {
	if s.client.token == "" {
		return fmt.Errorf("not connected")
	}

	queryParams := map[string]string{
		"id":     s.client.token,
		"symbol": symbol,
	}

	var result string
	err := s.client.get("/Subscribe", queryParams, &result)
	return err
}

// SubscribeMany subscribe หลายสัญลักษณ์
func (s *SubscriptionService) SubscribeMany(symbols []string) error {
	if s.client.token == "" {
		return fmt.Errorf("not connected")
	}

	queryParams := map[string]string{
		"id": s.client.token,
	}

	for i, symbol := range symbols {
		queryParams[fmt.Sprintf("symbols[%d]", i)] = symbol
	}

	var result string
	err := s.client.get("/SubscribeMany", queryParams, &result)
	return err
}

// Unsubscribe unsubscribe สัญลักษณ์
func (s *SubscriptionService) Unsubscribe(symbol string) error {
	if s.client.token == "" {
		return fmt.Errorf("not connected")
	}

	queryParams := map[string]string{
		"id":     s.client.token,
		"symbol": symbol,
	}

	var result string
	err := s.client.get("/UnSubscribe", queryParams, &result)
	return err
}

// UnsubscribeMany unsubscribe หลายสัญลักษณ์
func (s *SubscriptionService) UnsubscribeMany(symbols []string) error {
	if s.client.token == "" {
		return fmt.Errorf("not connected")
	}

	queryParams := map[string]string{
		"id": s.client.token,
	}

	for i, symbol := range symbols {
		queryParams[fmt.Sprintf("symbols[%d]", i)] = symbol
	}

	var result string
	err := s.client.get("/UnSubscribeMany", queryParams, &result)
	return err
}

// SubscribeTickValue subscribe tick value
func (s *SubscriptionService) SubscribeTickValue(symbol string) error {
	if s.client.token == "" {
		return fmt.Errorf("not connected")
	}

	queryParams := map[string]string{
		"id":     s.client.token,
		"symbol": symbol,
	}

	var result string
	err := s.client.get("/SubscribeTickValue", queryParams, &result)
	return err
}

// SubscribeOhlc subscribe OHLC
func (s *SubscriptionService) SubscribeOhlc(symbol, timeframe string) error {
	if s.client.token == "" {
		return fmt.Errorf("not connected")
	}

	queryParams := map[string]string{
		"id":        s.client.token,
		"symbol":    symbol,
		"timeframe": timeframe,
	}

	var result string
	err := s.client.get("/SubscribeOhlc", queryParams, &result)
	return err
}

// UnsubscribeOhlc unsubscribe OHLC
func (s *SubscriptionService) UnsubscribeOhlc(symbol, timeframe string) error {
	if s.client.token == "" {
		return fmt.Errorf("not connected")
	}

	queryParams := map[string]string{
		"id":        s.client.token,
		"symbol":    symbol,
		"timeframe": timeframe,
	}

	var result string
	err := s.client.get("/UnsubscribeOhlc", queryParams, &result)
	return err
}

// SubscribeOrderBook subscribe order book
func (s *SubscriptionService) SubscribeOrderBook(symbol string) error {
	if s.client.token == "" {
		return fmt.Errorf("not connected")
	}

	queryParams := map[string]string{
		"id":     s.client.token,
		"symbol": symbol,
	}

	var result string
	err := s.client.get("/SubscribeOrderBook", queryParams, &result)
	return err
}

// UnsubscribeOrderBook unsubscribe order book
func (s *SubscriptionService) UnsubscribeOrderBook(symbol string) error {
	if s.client.token == "" {
		return fmt.Errorf("not connected")
	}

	queryParams := map[string]string{
		"id":     s.client.token,
		"symbol": symbol,
	}

	var result string
	err := s.client.get("/UnsubscribeOrderBook", queryParams, &result)
	return err
}

// SubscribeMarketWatch subscribe market watch
func (s *SubscriptionService) SubscribeMarketWatch() error {
	if s.client.token == "" {
		return fmt.Errorf("not connected")
	}

	queryParams := map[string]string{
		"id": s.client.token,
	}

	var result string
	err := s.client.get("/SubscribeMarketWatch", queryParams, &result)
	return err
}

// GetWebSocketURL ดึง WebSocket URL สำหรับ events
func (s *SubscriptionService) GetWebSocketURL() (string, error) {
	// WebSocket endpoints จะใช้ผ่าน /On* paths
	// เช่น /OnQuote, /OnOrderUpdate, etc.
	// ต้องใช้ WebSocket client แยกต่างหาก
	return s.client.baseURL, nil
}
