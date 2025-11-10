package mt5client

import "fmt"

// StatsService จัดการสถิติการเทรด
type StatsService struct {
	client *Client
}

// GetTradeStats ดึงสถิติการเทรดทั้งหมด
func (s *StatsService) GetTradeStats() (*TradeStats, error) {
	if s.client.token == "" {
		return nil, fmt.Errorf("not connected")
	}

	queryParams := map[string]string{
		"id": s.client.token,
	}

	var stats TradeStats
	err := s.client.get("/TradeStats", queryParams, &stats)
	if err != nil {
		return nil, err
	}

	return &stats, nil
}

// GetEquityHistory ดึงประวัติ Equity (สำหรับกราฟ)
func (s *StatsService) GetEquityHistory(from, to string) ([]map[string]interface{}, error) {
	if s.client.token == "" {
		return nil, fmt.Errorf("not connected")
	}

	queryParams := map[string]string{
		"id":   s.client.token,
		"from": from,
		"to":   to,
	}

	var history []map[string]interface{}
	err := s.client.get("/TradeStatsEquityHistory", queryParams, &history)
	if err != nil {
		return nil, err
	}

	return history, nil
}
