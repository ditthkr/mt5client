package mt5client

import "fmt"

// OrderService จัดการคำสั่งซื้อขาย
type OrderService struct {
	client *Client
}

// GetOpened ดึงคำสั่งที่เปิดอยู่ทั้งหมด
func (r *OrderService) GetOpened() ([]Order, error) {
	if r.client.token == "" {
		return nil, fmt.Errorf("not connected")
	}

	queryParams := map[string]string{
		"id": r.client.token,
	}

	var orders []Order
	err := r.client.get("/OpenedOrders", queryParams, &orders)
	if err != nil {
		return nil, err
	}

	return orders, nil
}

// GetOpenedByTicket ดึงคำสั่งที่เปิดอยู่ตาม ticket
func (r *OrderService) GetOpenedByTicket(ticket int64) (*Order, error) {
	if r.client.token == "" {
		return nil, fmt.Errorf("not connected")
	}

	queryParams := map[string]string{
		"id":     r.client.token,
		"ticket": fmt.Sprintf("%d", ticket),
	}

	var order Order
	err := r.client.get("/OpenedOrder", queryParams, &order)
	if err != nil {
		return nil, err
	}

	return &order, nil
}

// GetOpenedTickets ดึง tickets ของคำสั่งที่เปิดอยู่
func (r *OrderService) GetOpenedTickets() ([]int64, error) {
	if r.client.token == "" {
		return nil, fmt.Errorf("not connected")
	}

	queryParams := map[string]string{
		"id": r.client.token,
	}

	var tickets []int64
	err := r.client.get("/OpenedOrdersTickets", queryParams, &tickets)
	if err != nil {
		return nil, err
	}

	return tickets, nil
}

// GetClosed ดึงคำสั่งที่ปิดแล้ว
func (r *OrderService) GetClosed(from, to string) ([]Order, error) {
	if r.client.token == "" {
		return nil, fmt.Errorf("not connected")
	}

	queryParams := map[string]string{
		"id":   r.client.token,
		"from": from,
		"to":   to,
	}

	var orders []Order
	err := r.client.get("/ClosedOrders", queryParams, &orders)
	if err != nil {
		return nil, err
	}

	return orders, nil
}

// GetPendingHistory ดึงประวัติคำสั่ง pending
func (r *OrderService) GetPendingHistory(from, to string) ([]Order, error) {
	if r.client.token == "" {
		return nil, fmt.Errorf("not connected")
	}

	queryParams := map[string]string{
		"id":   r.client.token,
		"from": from,
		"to":   to,
	}

	var orders []Order
	err := r.client.get("/PendingOrderHistory", queryParams, &orders)
	if err != nil {
		return nil, err
	}

	return orders, nil
}
