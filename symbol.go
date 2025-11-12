package mt5client

import "fmt"

// SymbolService จัดการสัญลักษณ์
type SymbolService struct {
	client *Client
}

// GetList ดึงรายการสัญลักษณ์ทั้งหมด
func (r *SymbolService) GetList() ([]string, error) {
	if r.client.token == "" {
		return nil, fmt.Errorf("not connected")
	}

	queryParams := map[string]string{
		"id": r.client.token,
	}

	var symbols []string
	err := r.client.get("/SymbolList", queryParams, &symbols)
	if err != nil {
		return nil, err
	}

	return symbols, nil
}

// GetParams ดึงพารามิเตอร์ของสัญลักษณ์
func (r *SymbolService) GetParams(symbol string) (*SymbolParams, error) {
	if r.client.token == "" {
		return nil, fmt.Errorf("not connected")
	}

	queryParams := map[string]string{
		"id":     r.client.token,
		"symbol": symbol,
	}

	var symbolParams SymbolParams
	err := r.client.get("/SymbolParams", queryParams, &symbolParams)
	if err != nil {
		return nil, err
	}

	return &symbolParams, nil
}

// GetParamsMany ดึงพารามิเตอร์หลายสัญลักษณ์
func (r *SymbolService) GetParamsMany(symbols []string) ([]SymbolParams, error) {
	if r.client.token == "" {
		return nil, fmt.Errorf("not connected")
	}

	queryParams := map[string]string{
		"id": r.client.token,
	}

	for i, symbol := range symbols {
		queryParams[fmt.Sprintf("symbols[%d]", i)] = symbol
	}

	var symbolParams []SymbolParams
	err := r.client.get("/SymbolParamsMany", queryParams, &symbolParams)
	if err != nil {
		return nil, err
	}

	return symbolParams, nil
}

// GetSessions ดึงข้อมูลเซสชันของสัญลักษณ์
func (r *SymbolService) GetSessions(symbol string) (map[string]interface{}, error) {
	if r.client.token == "" {
		return nil, fmt.Errorf("not connected")
	}

	queryParams := map[string]string{
		"id":     r.client.token,
		"symbol": symbol,
	}

	var sessions map[string]interface{}
	err := r.client.get("/SymbolSessions", queryParams, &sessions)
	if err != nil {
		return nil, err
	}

	return sessions, nil
}

// GetAll ดึงสัญลักษณ์ทั้งหมด (รวมข้อมูล)
func (r *SymbolService) GetAll() ([]SymbolInfo, error) {
	if r.client.token == "" {
		return nil, fmt.Errorf("not connected")
	}

	queryParams := map[string]string{
		"id": r.client.token,
	}

	var symbols []SymbolInfo
	err := r.client.get("/Symbols", queryParams, &symbols)
	if err != nil {
		return nil, err
	}

	return symbols, nil
}

// GetSubscribed ดึงสัญลักษณ์ที่ subscribe อยู่
func (r *SymbolService) GetSubscribed() ([]string, error) {
	if r.client.token == "" {
		return nil, fmt.Errorf("not connected")
	}

	queryParams := map[string]string{
		"id": r.client.token,
	}

	var symbols []string
	err := r.client.get("/SubscribedSymbols", queryParams, &symbols)
	if err != nil {
		return nil, err
	}

	return symbols, nil
}
