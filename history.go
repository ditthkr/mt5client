package mt5client

import "fmt"

// HistoryService จัดการประวัติ
type HistoryService struct {
	client *Client
}

// GetOrders ดึงประวัติคำสั่งซื้อขาย
func (r *HistoryService) GetOrders(from, to string) ([]Order, error) {
	if r.client.token == "" {
		return nil, fmt.Errorf("not connected")
	}

	queryParams := map[string]string{
		"id":   r.client.token,
		"from": from,
		"to":   to,
	}

	var orders []Order
	err := r.client.get("/OrderHistory", queryParams, &orders)
	if err != nil {
		return nil, err
	}

	return orders, nil
}

// GetOrdersPagination ดึงประวัติคำสั่งแบบแบ่งหน้า
func (r *HistoryService) GetOrdersPagination(from, to string, page, pageSize int) ([]Order, error) {
	if r.client.token == "" {
		return nil, fmt.Errorf("not connected")
	}

	queryParams := map[string]string{
		"id":       r.client.token,
		"from":     from,
		"to":       to,
		"page":     fmt.Sprintf("%d", page),
		"pageSize": fmt.Sprintf("%d", pageSize),
	}

	var orders []Order
	err := r.client.get("/OrderHistoryPagination", queryParams, &orders)
	if err != nil {
		return nil, err
	}

	return orders, nil
}

// IsOrderHistoryDownloadComplete ตรวจสอบว่าดาวน์โหลดประวัติเสร็จหรือไม่
func (r *HistoryService) IsOrderHistoryDownloadComplete() (bool, error) {
	if r.client.token == "" {
		return false, fmt.Errorf("not connected")
	}

	queryParams := map[string]string{
		"id": r.client.token,
	}

	var result bool
	err := r.client.get("/OrderHistoryDownloadComplete", queryParams, &result)
	if err != nil {
		return false, err
	}

	return result, nil
}

// GetPositions ดึงตำแหน่งในประวัติ
func (r *HistoryService) GetPositions(from, to string) ([]HistoryPosition, error) {
	if r.client.token == "" {
		return nil, fmt.Errorf("not connected")
	}

	queryParams := map[string]string{
		"id":   r.client.token,
		"from": from,
		"to":   to,
	}

	var positions []HistoryPosition
	err := r.client.get("/HistoryPositions", queryParams, &positions)
	if err != nil {
		return nil, err
	}

	return positions, nil
}

// GetPositionsByCloseTime ดึงตำแหน่งตามเวลาปิด
func (r *HistoryService) GetPositionsByCloseTime(from, to string) ([]HistoryPosition, error) {
	if r.client.token == "" {
		return nil, fmt.Errorf("not connected")
	}

	queryParams := map[string]string{
		"id":   r.client.token,
		"from": from,
		"to":   to,
	}

	var positions []HistoryPosition
	err := r.client.get("/HistoryPositionsByCloseTime", queryParams, &positions)
	if err != nil {
		return nil, err
	}

	return positions, nil
}

// GetDealsByPositionId ดึงดีลตาม Position ID
func (r *HistoryService) GetDealsByPositionId(positionId int64) ([]Deal, error) {
	if r.client.token == "" {
		return nil, fmt.Errorf("not connected")
	}

	queryParams := map[string]string{
		"id":         r.client.token,
		"positionId": fmt.Sprintf("%d", positionId),
	}

	var deals []Deal
	err := r.client.get("/HistoryDealsByPositionId", queryParams, &deals)
	if err != nil {
		return nil, err
	}

	return deals, nil
}
