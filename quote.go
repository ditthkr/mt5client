package mt5client

import "fmt"

// QuoteService จัดการราคา
type QuoteService struct {
	client *Client
}

// Get ดึงราคาของสัญลักษณ์
func (s *QuoteService) Get(symbol string) (*Quote, error) {
	if s.client.token == "" {
		return nil, fmt.Errorf("not connected")
	}

	queryParams := map[string]string{
		"id":     s.client.token,
		"symbol": symbol,
	}

	var quote Quote
	err := s.client.get("/GetQuote", queryParams, &quote)
	if err != nil {
		return nil, err
	}

	return &quote, nil
}

// GetMany ดึงราคาหลายสัญลักษณ์
func (s *QuoteService) GetMany(symbols []string) ([]Quote, error) {
	if s.client.token == "" {
		return nil, fmt.Errorf("not connected")
	}

	queryParams := map[string]string{
		"id": s.client.token,
	}

	// เพิ่ม symbols เป็น array
	for i, symbol := range symbols {
		queryParams[fmt.Sprintf("symbols[%d]", i)] = symbol
	}

	var quotes []Quote
	err := s.client.get("/GetQuoteMany", queryParams, &quotes)
	if err != nil {
		return nil, err
	}

	return quotes, nil
}

// GetTickValueMany ดึง tick value หลายสัญลักษณ์
func (s *QuoteService) GetTickValueMany(symbols []string) (map[string]float64, error) {
	if s.client.token == "" {
		return nil, fmt.Errorf("not connected")
	}

	queryParams := map[string]string{
		"id": s.client.token,
	}

	for i, symbol := range symbols {
		queryParams[fmt.Sprintf("symbols[%d]", i)] = symbol
	}

	var tickValues map[string]float64
	err := s.client.get("/GetTickValueMany", queryParams, &tickValues)
	if err != nil {
		return nil, err
	}

	return tickValues, nil
}

// GetTickValueWithSize ดึง tick value พร้อมขนาด
func (s *QuoteService) GetTickValueWithSize(symbol string, volume float64) (float64, error) {
	if s.client.token == "" {
		return 0, fmt.Errorf("not connected")
	}

	queryParams := map[string]string{
		"id":     s.client.token,
		"symbol": symbol,
		"volume": fmt.Sprintf("%.2f", volume),
	}

	var tickValue float64
	err := s.client.get("/TickValueWithSize", queryParams, &tickValue)
	if err != nil {
		return 0, err
	}

	return tickValue, nil
}

// IsQuoteSession ตรวจสอบว่าอยู่ในเซสชันราคาหรือไม่
func (s *QuoteService) IsQuoteSession(symbol string) (bool, error) {
	if s.client.token == "" {
		return false, fmt.Errorf("not connected")
	}

	queryParams := map[string]string{
		"id":     s.client.token,
		"symbol": symbol,
	}

	var result bool
	err := s.client.get("/IsQuoteSession", queryParams, &result)
	if err != nil {
		return false, err
	}

	return result, nil
}

// IsQuoteSessionMany ตรวจสอบหลายสัญลักษณ์
func (s *QuoteService) IsQuoteSessionMany(symbols []string) (map[string]bool, error) {
	if s.client.token == "" {
		return nil, fmt.Errorf("not connected")
	}

	queryParams := map[string]string{
		"id": s.client.token,
	}

	for i, symbol := range symbols {
		queryParams[fmt.Sprintf("symbols[%d]", i)] = symbol
	}

	var results map[string]bool
	err := s.client.get("/IsQuoteSessionMany", queryParams, &results)
	if err != nil {
		return nil, err
	}

	return results, nil
}

// IsTradeSession ตรวจสอบว่าอยู่ในเซสชันเทรดหรือไม่
func (s *QuoteService) IsTradeSession(symbol string) (bool, error) {
	if s.client.token == "" {
		return false, fmt.Errorf("not connected")
	}

	queryParams := map[string]string{
		"id":     s.client.token,
		"symbol": symbol,
	}

	var result bool
	err := s.client.get("/IsTradeSession", queryParams, &result)
	if err != nil {
		return false, err
	}

	return result, nil
}

// IsTradeSessionMany ตรวจสอบหลายสัญลักษณ์
func (s *QuoteService) IsTradeSessionMany(symbols []string) (map[string]bool, error) {
	if s.client.token == "" {
		return nil, fmt.Errorf("not connected")
	}

	queryParams := map[string]string{
		"id": s.client.token,
	}

	for i, symbol := range symbols {
		queryParams[fmt.Sprintf("symbols[%d]", i)] = symbol
	}

	var results map[string]bool
	err := s.client.get("/IsTradeSessionMany", queryParams, &results)
	if err != nil {
		return nil, err
	}

	return results, nil
}
