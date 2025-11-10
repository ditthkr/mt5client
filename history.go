package mt5client

import "fmt"

// HistoryService จัดการประวัติ
type HistoryService struct {
	client *Client
}

// GetOrders ดึงประวัติคำสั่งซื้อขาย
func (s *HistoryService) GetOrders(from, to string) ([]Order, error) {
	if s.client.token == "" {
		return nil, fmt.Errorf("not connected")
	}

	queryParams := map[string]string{
		"id":   s.client.token,
		"from": from,
		"to":   to,
	}

	var orders []Order
	err := s.client.get("/OrderHistory", queryParams, &orders)
	if err != nil {
		return nil, err
	}

	return orders, nil
}

// GetOrdersPagination ดึงประวัติคำสั่งแบบแบ่งหน้า
func (s *HistoryService) GetOrdersPagination(from, to string, page, pageSize int) ([]Order, error) {
	if s.client.token == "" {
		return nil, fmt.Errorf("not connected")
	}

	queryParams := map[string]string{
		"id":       s.client.token,
		"from":     from,
		"to":       to,
		"page":     fmt.Sprintf("%d", page),
		"pageSize": fmt.Sprintf("%d", pageSize),
	}

	var orders []Order
	err := s.client.get("/OrderHistoryPagination", queryParams, &orders)
	if err != nil {
		return nil, err
	}

	return orders, nil
}

// IsOrderHistoryDownloadComplete ตรวจสอบว่าดาวน์โหลดประวัติเสร็จหรือไม่
func (s *HistoryService) IsOrderHistoryDownloadComplete() (bool, error) {
	if s.client.token == "" {
		return false, fmt.Errorf("not connected")
	}

	queryParams := map[string]string{
		"id": s.client.token,
	}

	var result bool
	err := s.client.get("/OrderHistoryDownloadComplete", queryParams, &result)
	if err != nil {
		return false, err
	}

	return result, nil
}

// GetPositions ดึงตำแหน่งในประวัติ
func (s *HistoryService) GetPositions(from, to string) ([]HistoryPosition, error) {
	if s.client.token == "" {
		return nil, fmt.Errorf("not connected")
	}

	queryParams := map[string]string{
		"id":   s.client.token,
		"from": from,
		"to":   to,
	}

	var positions []HistoryPosition
	err := s.client.get("/HistoryPositions", queryParams, &positions)
	if err != nil {
		return nil, err
	}

	return positions, nil
}

// GetPositionsByCloseTime ดึงตำแหน่งตามเวลาปิด
func (s *HistoryService) GetPositionsByCloseTime(from, to string) ([]HistoryPosition, error) {
	if s.client.token == "" {
		return nil, fmt.Errorf("not connected")
	}

	queryParams := map[string]string{
		"id":   s.client.token,
		"from": from,
		"to":   to,
	}

	var positions []HistoryPosition
	err := s.client.get("/HistoryPositionsByCloseTime", queryParams, &positions)
	if err != nil {
		return nil, err
	}

	return positions, nil
}

// GetDealsByPositionId ดึงดีลตาม Position ID
func (s *HistoryService) GetDealsByPositionId(positionId int64) ([]Deal, error) {
	if s.client.token == "" {
		return nil, fmt.Errorf("not connected")
	}

	queryParams := map[string]string{
		"id":         s.client.token,
		"positionId": fmt.Sprintf("%d", positionId),
	}

	var deals []Deal
	err := s.client.get("/HistoryDealsByPositionId", queryParams, &deals)
	if err != nil {
		return nil, err
	}

	return deals, nil
}
