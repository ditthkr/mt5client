package mt5client

import "fmt"

// PriceService จัดการข้อมูลราคา
type PriceService struct {
	client *Client
}

// GetHistory ดึงประวัติราคา
func (r *PriceService) GetHistory(symbol, timeframe string, count int) ([]Bar, error) {
	if r.client.token == "" {
		return nil, fmt.Errorf("not connected")
	}

	queryParams := map[string]string{
		"id":        r.client.token,
		"symbol":    symbol,
		"timeframe": timeframe,
		"count":     fmt.Sprintf("%d", count),
	}

	var bars []Bar
	err := r.client.get("/PriceHistory", queryParams, &bars)
	if err != nil {
		return nil, err
	}

	return bars, nil
}

// GetHistoryEx ดึงประวัติราคาแบบ Extended
func (r *PriceService) GetHistoryEx(symbol, timeframe, from, to string) ([]Bar, error) {
	if r.client.token == "" {
		return nil, fmt.Errorf("not connected")
	}

	queryParams := map[string]string{
		"id":        r.client.token,
		"symbol":    symbol,
		"timeframe": timeframe,
		"from":      from,
		"to":        to,
	}

	var bars []Bar
	err := r.client.get("/PriceHistoryEx", queryParams, &bars)
	if err != nil {
		return nil, err
	}

	return bars, nil
}

// GetHistoryMany ดึงประวัติหลายสัญลักษณ์
func (r *PriceService) GetHistoryMany(symbols []string, timeframe string, count int) (map[string][]Bar, error) {
	if r.client.token == "" {
		return nil, fmt.Errorf("not connected")
	}

	queryParams := map[string]string{
		"id":        r.client.token,
		"timeframe": timeframe,
		"count":     fmt.Sprintf("%d", count),
	}

	for i, symbol := range symbols {
		queryParams[fmt.Sprintf("symbols[%d]", i)] = symbol
	}

	var result map[string][]Bar
	err := r.client.get("/PriceHistoryMany", queryParams, &result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// GetHistoryExMany ดึงประวัติแบบ Extended หลายสัญลักษณ์
func (r *PriceService) GetHistoryExMany(symbols []string, timeframe, from, to string) (map[string][]Bar, error) {
	if r.client.token == "" {
		return nil, fmt.Errorf("not connected")
	}

	queryParams := map[string]string{
		"id":        r.client.token,
		"timeframe": timeframe,
		"from":      from,
		"to":        to,
	}

	for i, symbol := range symbols {
		queryParams[fmt.Sprintf("symbols[%d]", i)] = symbol
	}

	var result map[string][]Bar
	err := r.client.get("/PriceHistoryExMany", queryParams, &result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// GetHistoryHighLow ดึง High/Low ในช่วงเวลา
func (r *PriceService) GetHistoryHighLow(symbol, from, to string) (map[string]float64, error) {
	if r.client.token == "" {
		return nil, fmt.Errorf("not connected")
	}

	queryParams := map[string]string{
		"id":     r.client.token,
		"symbol": symbol,
		"from":   from,
		"to":     to,
	}

	var result map[string]float64
	err := r.client.get("/PriceHistoryHighLow", queryParams, &result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// GetHistoryToday ดึงประวัติวันนี้
func (r *PriceService) GetHistoryToday(symbol, timeframe string) ([]Bar, error) {
	if r.client.token == "" {
		return nil, fmt.Errorf("not connected")
	}

	queryParams := map[string]string{
		"id":        r.client.token,
		"symbol":    symbol,
		"timeframe": timeframe,
	}

	var bars []Bar
	err := r.client.get("/PriceHistoryToday", queryParams, &bars)
	if err != nil {
		return nil, err
	}

	return bars, nil
}

// GetHistoryTodayMany ดึงประวัติวันนี้หลายสัญลักษณ์
func (r *PriceService) GetHistoryTodayMany(symbols []string, timeframe string) (map[string][]Bar, error) {
	if r.client.token == "" {
		return nil, fmt.Errorf("not connected")
	}

	queryParams := map[string]string{
		"id":        r.client.token,
		"timeframe": timeframe,
	}

	for i, symbol := range symbols {
		queryParams[fmt.Sprintf("symbols[%d]", i)] = symbol
	}

	var result map[string][]Bar
	err := r.client.get("/PriceHistoryTodayMany", queryParams, &result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// GetHistoryMonth ดึงประวัติรายเดือน
func (r *PriceService) GetHistoryMonth(symbol, timeframe string, year, month int) ([]Bar, error) {
	if r.client.token == "" {
		return nil, fmt.Errorf("not connected")
	}

	queryParams := map[string]string{
		"id":        r.client.token,
		"symbol":    symbol,
		"timeframe": timeframe,
		"year":      fmt.Sprintf("%d", year),
		"month":     fmt.Sprintf("%d", month),
	}

	var bars []Bar
	err := r.client.get("/PriceHistoryMonth", queryParams, &bars)
	if err != nil {
		return nil, err
	}

	return bars, nil
}

// GetHistoryMonthMany ดึงประวัติรายเดือนหลายสัญลักษณ์
func (r *PriceService) GetHistoryMonthMany(symbols []string, timeframe string, year, month int) (map[string][]Bar, error) {
	if r.client.token == "" {
		return nil, fmt.Errorf("not connected")
	}

	queryParams := map[string]string{
		"id":        r.client.token,
		"timeframe": timeframe,
		"year":      fmt.Sprintf("%d", year),
		"month":     fmt.Sprintf("%d", month),
	}

	for i, symbol := range symbols {
		queryParams[fmt.Sprintf("symbols[%d]", i)] = symbol
	}

	var result map[string][]Bar
	err := r.client.get("/PriceHistoryMonthMany", queryParams, &result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// RequestTickHistory ขอประวัติ tick
func (r *PriceService) RequestTickHistory(symbol, from, to string) error {
	if r.client.token == "" {
		return fmt.Errorf("not connected")
	}

	queryParams := map[string]string{
		"id":     r.client.token,
		"symbol": symbol,
		"from":   from,
		"to":     to,
	}

	var result string
	err := r.client.get("/TickHistoryRequest", queryParams, &result)
	return err
}

// StopTickHistory หยุดการขอประวัติ tick
func (r *PriceService) StopTickHistory() error {
	if r.client.token == "" {
		return fmt.Errorf("not connected")
	}

	queryParams := map[string]string{
		"id": r.client.token,
	}

	var result string
	err := r.client.get("/TickHistoryStop", queryParams, &result)
	return err
}
