package mt5client

import "fmt"

// ConnectionService จัดการการเชื่อมต่อ
type ConnectionService struct {
	client *Client
}

// Connect เชื่อมต่อ MT5
func (s *ConnectionService) Connect(params ConnectParams) (string, error) {
	queryParams := map[string]string{
		"user":     fmt.Sprintf("%d", params.User),
		"password": params.Password,
		"host":     params.Host,
		"port":     fmt.Sprintf("%d", params.Port),
	}

	var token string
	err := s.client.get("/Connect", queryParams, &token)
	if err != nil {
		return "", err
	}

	s.client.SetToken(token)
	return token, nil
}

// ConnectEx เชื่อมต่อแบบ Extended (ใช้ชื่อ server แทน host/port)
func (s *ConnectionService) ConnectEx(user int64, password, server string) (string, error) {
	queryParams := map[string]string{
		"user":     fmt.Sprintf("%d", user),
		"password": password,
		"server":   server,
	}

	var token string
	err := s.client.get("/ConnectEx", queryParams, &token)
	if err != nil {
		return "", err
	}

	s.client.SetToken(token)
	return token, nil
}

// ConnectProxy เชื่อมต่อผ่าน Proxy
func (s *ConnectionService) ConnectProxy(params ConnectParams, proxyType, proxyHost string, proxyPort int) (string, error) {
	queryParams := map[string]string{
		"user":      fmt.Sprintf("%d", params.User),
		"password":  params.Password,
		"host":      params.Host,
		"port":      fmt.Sprintf("%d", params.Port),
		"proxyType": proxyType,
		"proxyHost": proxyHost,
		"proxyPort": fmt.Sprintf("%d", proxyPort),
	}

	var token string
	err := s.client.get("/ConnectProxy", queryParams, &token)
	if err != nil {
		return "", err
	}

	s.client.SetToken(token)
	return token, nil
}

// Disconnect ตัดการเชื่อมต่อ
func (s *ConnectionService) Disconnect() error {
	if s.client.token == "" {
		return nil
	}

	queryParams := map[string]string{
		"id": s.client.token,
	}

	var result string
	err := s.client.get("/Disconnect", queryParams, &result)
	if err != nil {
		return err
	}

	s.client.SetToken("")
	return nil
}

// IsConnected ตรวจสอบสถานะการเชื่อมต่อ
func (s *ConnectionService) IsConnected() (bool, error) {
	if s.client.token == "" {
		return false, nil
	}

	queryParams := map[string]string{
		"id": s.client.token,
	}

	var result string
	err := s.client.get("/CheckConnect", queryParams, &result)
	if err != nil {
		return false, err
	}

	return result == "OK", nil
}
