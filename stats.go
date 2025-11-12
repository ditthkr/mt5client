package mt5client

import "fmt"

// StatsService จัดการสถิติการเทรด
type StatsService struct {
	client *Client
}

// GetTradeStats ดึงสถิติการเทรดทั้งหมด
func (r *StatsService) GetTradeStats() (*TradeStats, error) {
	if r.client.token == "" {
		return nil, fmt.Errorf("not connected")
	}

	queryParams := map[string]string{
		"id": r.client.token,
	}

	var stats TradeStats
	err := r.client.get("/TradeStats", queryParams, &stats)
	if err != nil {
		return nil, err
	}

	return &stats, nil
}

// GetEquityHistory ดึงประวัติ Equity (สำหรับกราฟ)
func (r *StatsService) GetEquityHistory(from, to string) ([]map[string]interface{}, error) {
	if r.client.token == "" {
		return nil, fmt.Errorf("not connected")
	}

	queryParams := map[string]string{
		"id":   r.client.token,
		"from": from,
		"to":   to,
	}

	var history []map[string]interface{}
	err := r.client.get("/TradeStatsEquityHistory", queryParams, &history)
	if err != nil {
		return nil, err
	}

	return history, nil
}
