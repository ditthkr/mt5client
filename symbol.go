package mt5client

import "fmt"

// SymbolService จัดการสัญลักษณ์
type SymbolService struct {
	client *Client
}

// GetList ดึงรายการสัญลักษณ์ทั้งหมด
func (s *SymbolService) GetList() ([]string, error) {
	if s.client.token == "" {
		return nil, fmt.Errorf("not connected")
	}

	queryParams := map[string]string{
		"id": s.client.token,
	}

	var symbols []string
	err := s.client.get("/SymbolList", queryParams, &symbols)
	if err != nil {
		return nil, err
	}

	return symbols, nil
}

// GetParams ดึงพารามิเตอร์ของสัญลักษณ์
func (s *SymbolService) GetParams(symbol string) (*Symbol, error) {
	if s.client.token == "" {
		return nil, fmt.Errorf("not connected")
	}

	queryParams := map[string]string{
		"id":     s.client.token,
		"symbol": symbol,
	}

	var symbolInfo Symbol
	err := s.client.get("/SymbolParams", queryParams, &symbolInfo)
	if err != nil {
		return nil, err
	}

	return &symbolInfo, nil
}

// GetParamsMany ดึงพารามิเตอร์หลายสัญลักษณ์
func (s *SymbolService) GetParamsMany(symbols []string) ([]Symbol, error) {
	if s.client.token == "" {
		return nil, fmt.Errorf("not connected")
	}

	queryParams := map[string]string{
		"id": s.client.token,
	}

	for i, symbol := range symbols {
		queryParams[fmt.Sprintf("symbols[%d]", i)] = symbol
	}

	var symbolInfos []Symbol
	err := s.client.get("/SymbolParamsMany", queryParams, &symbolInfos)
	if err != nil {
		return nil, err
	}

	return symbolInfos, nil
}

// GetSessions ดึงข้อมูลเซสชันของสัญลักษณ์
func (s *SymbolService) GetSessions(symbol string) (map[string]interface{}, error) {
	if s.client.token == "" {
		return nil, fmt.Errorf("not connected")
	}

	queryParams := map[string]string{
		"id":     s.client.token,
		"symbol": symbol,
	}

	var sessions map[string]interface{}
	err := s.client.get("/SymbolSessions", queryParams, &sessions)
	if err != nil {
		return nil, err
	}

	return sessions, nil
}

// GetAll ดึงสัญลักษณ์ทั้งหมด (รวมข้อมูล)
func (s *SymbolService) GetAll() ([]Symbol, error) {
	if s.client.token == "" {
		return nil, fmt.Errorf("not connected")
	}

	queryParams := map[string]string{
		"id": s.client.token,
	}

	var symbols []Symbol
	err := s.client.get("/Symbols", queryParams, &symbols)
	if err != nil {
		return nil, err
	}

	return symbols, nil
}

// GetSubscribed ดึงสัญลักษณ์ที่ subscribe อยู่
func (s *SymbolService) GetSubscribed() ([]string, error) {
	if s.client.token == "" {
		return nil, fmt.Errorf("not connected")
	}

	queryParams := map[string]string{
		"id": s.client.token,
	}

	var symbols []string
	err := s.client.get("/SubscribedSymbols", queryParams, &symbols)
	if err != nil {
		return nil, err
	}

	return symbols, nil
}
