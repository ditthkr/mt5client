package mt5client

import "fmt"

// AccountService -
type AccountService struct {
	client *Client
}

// GetInfo ดึงข้อมูลบัญชีทั้งหมด
func (s *AccountService) GetInfo() (*Account, error) {
	if s.client.token == "" {
		return nil, fmt.Errorf("not connected")
	}

	queryParams := map[string]string{
		"id": s.client.token,
	}

	var account Account
	err := s.client.get("/Account", queryParams, &account)
	if err != nil {
		return nil, err
	}

	return &account, nil
}

// GetDetails ดึงรายละเอียดบัญชี
func (s *AccountService) GetDetails() (*Account, error) {
	if s.client.token == "" {
		return nil, fmt.Errorf("not connected")
	}

	queryParams := map[string]string{
		"id": s.client.token,
	}

	var account Account
	err := s.client.get("/AccountDetails", queryParams, &account)
	if err != nil {
		return nil, err
	}

	return &account, nil
}

// GetSummary ดึงสรุปบัญชี
func (s *AccountService) GetSummary() (map[string]interface{}, error) {
	if s.client.token == "" {
		return nil, fmt.Errorf("not connected")
	}

	queryParams := map[string]string{
		"id": s.client.token,
	}

	var summary map[string]interface{}
	err := s.client.get("/AccountSummary", queryParams, &summary)
	if err != nil {
		return nil, err
	}

	return summary, nil
}

// GetEquityHistory ดึงประวัติ Equity
func (s *AccountService) GetEquityHistory(from, to string) ([]map[string]interface{}, error) {
	if s.client.token == "" {
		return nil, fmt.Errorf("not connected")
	}

	queryParams := map[string]string{
		"id":   s.client.token,
		"from": from,
		"to":   to,
	}

	var history []map[string]interface{}
	err := s.client.get("/EquityHistory", queryParams, &history)
	if err != nil {
		return nil, err
	}

	return history, nil
}
