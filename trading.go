package mt5client

import "fmt"

// TradingService จัดการการเทรด
type TradingService struct {
	client *Client
}

// Send ส่งคำสั่งซื้อขาย
func (s *TradingService) Send(req OrderRequest) (*TradeResult, error) {
	if s.client.token == "" {
		return nil, fmt.Errorf("not connected")
	}

	queryParams := map[string]string{
		"id":     s.client.token,
		"symbol": req.Symbol,
		"type":   req.Type,
		"volume": fmt.Sprintf("%.2f", req.Volume),
	}

	if req.Price > 0 {
		queryParams["price"] = fmt.Sprintf("%.5f", req.Price)
	}
	if req.StopLoss > 0 {
		queryParams["stopLoss"] = fmt.Sprintf("%.5f", req.StopLoss)
	}
	if req.TakeProfit > 0 {
		queryParams["takeProfit"] = fmt.Sprintf("%.5f", req.TakeProfit)
	}
	if req.Comment != "" {
		queryParams["comment"] = req.Comment
	}
	if req.MagicNumber > 0 {
		queryParams["magicNumber"] = fmt.Sprintf("%d", req.MagicNumber)
	}

	var result TradeResult
	err := s.client.get("/OrderSend", queryParams, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

// Buy ซื้อทันที
func (s *TradingService) Buy(symbol string, volume float64, sl, tp float64) (*TradeResult, error) {
	return s.Send(OrderRequest{
		Symbol:     symbol,
		Type:       "Buy",
		Volume:     volume,
		StopLoss:   sl,
		TakeProfit: tp,
	})
}

// Sell ขายทันที
func (s *TradingService) Sell(symbol string, volume float64, sl, tp float64) (*TradeResult, error) {
	return s.Send(OrderRequest{
		Symbol:     symbol,
		Type:       "Sell",
		Volume:     volume,
		StopLoss:   sl,
		TakeProfit: tp,
	})
}

// Modify แก้ไขคำสั่ง
func (s *TradingService) Modify(ticket int64, price, sl, tp float64) error {
	if s.client.token == "" {
		return fmt.Errorf("not connected")
	}

	queryParams := map[string]string{
		"id":     s.client.token,
		"ticket": fmt.Sprintf("%d", ticket),
	}

	if price > 0 {
		queryParams["price"] = fmt.Sprintf("%.5f", price)
	}
	if sl > 0 {
		queryParams["stopLoss"] = fmt.Sprintf("%.5f", sl)
	}
	if tp > 0 {
		queryParams["takeProfit"] = fmt.Sprintf("%.5f", tp)
	}

	var result string
	err := s.client.get("/OrderModify", queryParams, &result)
	return err
}

// Close ปิดคำสั่ง
func (s *TradingService) Close(ticket int64, volume float64) error {
	if s.client.token == "" {
		return fmt.Errorf("not connected")
	}

	queryParams := map[string]string{
		"id":     s.client.token,
		"ticket": fmt.Sprintf("%d", ticket),
	}

	if volume > 0 {
		queryParams["volume"] = fmt.Sprintf("%.2f", volume)
	}

	var result string
	err := s.client.get("/OrderClose", queryParams, &result)
	return err
}
