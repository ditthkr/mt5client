package mt5client

import "fmt"

// ServiceFunctions ฟังก์ชันบริการทั่วไป
type ServiceFunctions struct {
	client *Client
}

// GetVersion ดึงเวอร์ชันของ API
func (s *ServiceFunctions) GetVersion() (string, error) {
	var version string
	err := s.client.get("/Version", nil, &version)
	return version, err
}

// Ping ทดสอบการเชื่อมต่อ
func (s *ServiceFunctions) Ping() (string, error) {
	var result string
	err := s.client.get("/Ping", nil, &result)
	return result, err
}

// PingHost ทดสอบการเชื่อมต่อไปยัง host
func (s *ServiceFunctions) PingHost(host string) (bool, error) {
	queryParams := map[string]string{
		"host": host,
	}

	var result bool
	err := s.client.get("/PingHost", queryParams, &result)
	return result, err
}

// PingHostMany ทดสอบหลาย hosts
func (s *ServiceFunctions) PingHostMany(hosts []string) (map[string]bool, error) {
	queryParams := map[string]string{}

	for i, host := range hosts {
		queryParams[fmt.Sprintf("hosts[%d]", i)] = host
	}

	var results map[string]bool
	err := s.client.get("/PingHostMany", queryParams, &results)
	return results, err
}

// Search ค้นหาสัญลักษณ์
func (s *ServiceFunctions) Search(keyword string) ([]string, error) {
	if s.client.token == "" {
		return nil, fmt.Errorf("not connected")
	}

	queryParams := map[string]string{
		"id":      s.client.token,
		"keyword": keyword,
	}

	var results []string
	err := s.client.get("/Search", queryParams, &results)
	return results, err
}

// GetServerTimezone ดึงข้อมูล timezone ของ server
func (s *ServiceFunctions) GetServerTimezone() (string, error) {
	if s.client.token == "" {
		return "", fmt.Errorf("not connected")
	}

	queryParams := map[string]string{
		"id": s.client.token,
	}

	var timezone string
	err := s.client.get("/ServerTimezone", queryParams, &timezone)
	return timezone, err
}

// GetClusterDetails ดึงข้อมูล cluster
func (s *ServiceFunctions) GetClusterDetails() (map[string]interface{}, error) {
	if s.client.token == "" {
		return nil, fmt.Errorf("not connected")
	}

	queryParams := map[string]string{
		"id": s.client.token,
	}

	var details map[string]interface{}
	err := s.client.get("/ClusterDetails", queryParams, &details)
	return details, err
}

// ChangePassword เปลี่ยนรหัสผ่าน
func (s *ServiceFunctions) ChangePassword(oldPassword, newPassword string) error {
	if s.client.token == "" {
		return fmt.Errorf("not connected")
	}

	queryParams := map[string]string{
		"id":          s.client.token,
		"oldPassword": oldPassword,
		"newPassword": newPassword,
	}

	var result string
	err := s.client.get("/ChangePassword", queryParams, &result)
	return err
}

// GetDemo ขอบัญชี demo
func (s *ServiceFunctions) GetDemo(server, name, email string) (map[string]interface{}, error) {
	queryParams := map[string]string{
		"server": server,
		"name":   name,
		"email":  email,
	}

	var result map[string]interface{}
	err := s.client.get("/GetDemo", queryParams, &result)
	return result, err
}

// GetRequiredMargin คำนวณ margin ที่ต้องการ
func (s *ServiceFunctions) GetRequiredMargin(symbol string, volume float64) (float64, error) {
	if s.client.token == "" {
		return 0, fmt.Errorf("not connected")
	}

	queryParams := map[string]string{
		"id":     s.client.token,
		"symbol": symbol,
		"volume": fmt.Sprintf("%.2f", volume),
	}

	var margin float64
	err := s.client.get("/RequiredMargin", queryParams, &margin)
	return margin, err
}

// GetMails ดึงอีเมล
func (s *ServiceFunctions) GetMails() ([]Mail, error) {
	if s.client.token == "" {
		return nil, fmt.Errorf("not connected")
	}

	queryParams := map[string]string{
		"id": s.client.token,
	}

	var mails []Mail
	err := s.client.get("/Mails", queryParams, &mails)
	return mails, err
}

// GetMarketWatchMany ดึง market watch หลายสัญลักษณ์
func (s *ServiceFunctions) GetMarketWatchMany(symbols []string) ([]MarketWatch, error) {
	if s.client.token == "" {
		return nil, fmt.Errorf("not connected")
	}

	queryParams := map[string]string{
		"id": s.client.token,
	}

	for i, symbol := range symbols {
		queryParams[fmt.Sprintf("symbols[%d]", i)] = symbol
	}

	var marketWatch []MarketWatch
	err := s.client.get("/MarketWatchMany", queryParams, &marketWatch)
	return marketWatch, err
}

// GetQuoteClient ดึงข้อมูล quote client
func (s *ServiceFunctions) GetQuoteClient() (map[string]interface{}, error) {
	if s.client.token == "" {
		return nil, fmt.Errorf("not connected")
	}

	queryParams := map[string]string{
		"id": s.client.token,
	}

	var client map[string]interface{}
	err := s.client.get("/QuoteClient", queryParams, &client)
	return client, err
}

// LoadServersDat โหลดไฟล์ servers.dat
func (s *ServiceFunctions) LoadServersDat(data []byte) error {
	var result string
	err := s.client.post("/LoadServersDat", nil, data, &result)
	return err
}

// GetMetricsApiKey ดึง API key สำหรับ metrics
func (s *ServiceFunctions) GetMetricsApiKey() (string, error) {
	if s.client.token == "" {
		return "", fmt.Errorf("not connected")
	}

	queryParams := map[string]string{
		"id": s.client.token,
	}

	var apiKey string
	err := s.client.get("/MetricsApiKey", queryParams, &apiKey)
	return apiKey, err
}

// GetReadMe ดึง README
func (s *ServiceFunctions) GetReadMe() (string, error) {
	var readme string
	err := s.client.get("/ReadMe", nil, &readme)
	return readme, err
}
