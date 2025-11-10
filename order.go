package mt5client

import "fmt"

// OrderService จัดการคำสั่งซื้อขาย
type OrderService struct {
	client *Client
}

// GetOpened ดึงคำสั่งที่เปิดอยู่ทั้งหมด
func (s *OrderService) GetOpened() ([]Order, error) {
	if s.client.token == "" {
		return nil, fmt.Errorf("not connected")
	}

	queryParams := map[string]string{
		"id": s.client.token,
	}

	var orders []Order
	err := s.client.get("/OpenedOrders", queryParams, &orders)
	if err != nil {
		return nil, err
	}

	return orders, nil
}

// GetOpenedByTicket ดึงคำสั่งที่เปิดอยู่ตาม ticket
func (s *OrderService) GetOpenedByTicket(ticket int64) (*Order, error) {
	if s.client.token == "" {
		return nil, fmt.Errorf("not connected")
	}

	queryParams := map[string]string{
		"id":     s.client.token,
		"ticket": fmt.Sprintf("%d", ticket),
	}

	var order Order
	err := s.client.get("/OpenedOrder", queryParams, &order)
	if err != nil {
		return nil, err
	}

	return &order, nil
}

// GetOpenedTickets ดึง tickets ของคำสั่งที่เปิดอยู่
func (s *OrderService) GetOpenedTickets() ([]int64, error) {
	if s.client.token == "" {
		return nil, fmt.Errorf("not connected")
	}

	queryParams := map[string]string{
		"id": s.client.token,
	}

	var tickets []int64
	err := s.client.get("/OpenedOrdersTickets", queryParams, &tickets)
	if err != nil {
		return nil, err
	}

	return tickets, nil
}

// GetClosed ดึงคำสั่งที่ปิดแล้ว
func (s *OrderService) GetClosed(from, to string) ([]Order, error) {
	if s.client.token == "" {
		return nil, fmt.Errorf("not connected")
	}

	queryParams := map[string]string{
		"id":   s.client.token,
		"from": from,
		"to":   to,
	}

	var orders []Order
	err := s.client.get("/ClosedOrders", queryParams, &orders)
	if err != nil {
		return nil, err
	}

	return orders, nil
}

// GetPendingHistory ดึงประวัติคำสั่ง pending
func (s *OrderService) GetPendingHistory(from, to string) ([]Order, error) {
	if s.client.token == "" {
		return nil, fmt.Errorf("not connected")
	}

	queryParams := map[string]string{
		"id":   s.client.token,
		"from": from,
		"to":   to,
	}

	var orders []Order
	err := s.client.get("/PendingOrderHistory", queryParams, &orders)
	if err != nil {
		return nil, err
	}

	return orders, nil
}
