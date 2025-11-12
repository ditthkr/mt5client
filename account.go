package mt5client

import "fmt"

// AccountService -
type AccountService struct {
	client *Client
}

// GetInfo ดึงข้อมูลบัญชีทั้งหมด
func (r *AccountService) GetInfo() (*Account, error) {
	if r.client.token == "" {
		return nil, fmt.Errorf("not connected")
	}

	queryParams := map[string]string{
		"id": r.client.token,
	}

	var account Account
	err := r.client.get("/Account", queryParams, &account)
	if err != nil {
		return nil, err
	}

	return &account, nil
}

// GetDetails ดึงรายละเอียดบัญชี
func (r *AccountService) GetDetails() (*Account, error) {
	if r.client.token == "" {
		return nil, fmt.Errorf("not connected")
	}

	queryParams := map[string]string{
		"id": r.client.token,
	}

	var account Account
	err := r.client.get("/AccountDetails", queryParams, &account)
	if err != nil {
		return nil, err
	}

	return &account, nil
}

// GetSummary ดึงสรุปบัญชี
func (r *AccountService) GetSummary() (map[string]interface{}, error) {
	if r.client.token == "" {
		return nil, fmt.Errorf("not connected")
	}

	queryParams := map[string]string{
		"id": r.client.token,
	}

	var summary map[string]interface{}
	err := r.client.get("/AccountSummary", queryParams, &summary)
	if err != nil {
		return nil, err
	}

	return summary, nil
}

// GetEquityHistory ดึงประวัติ Equity
func (r *AccountService) GetEquityHistory(from, to string) ([]map[string]interface{}, error) {
	if r.client.token == "" {
		return nil, fmt.Errorf("not connected")
	}

	queryParams := map[string]string{
		"id":   r.client.token,
		"from": from,
		"to":   to,
	}

	var history []map[string]interface{}
	err := r.client.get("/EquityHistory", queryParams, &history)
	if err != nil {
		return nil, err
	}

	return history, nil
}
