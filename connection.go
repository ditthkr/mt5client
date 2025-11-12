package mt5client

import "fmt"

// ConnectionService จัดการการเชื่อมต่อ
type ConnectionService struct {
	client *Client
}

// Connect เชื่อมต่อ MT5
func (r *ConnectionService) Connect(params ConnectParams) (string, error) {
	queryParams := map[string]string{
		"user":     fmt.Sprintf("%d", params.User),
		"password": params.Password,
		"host":     params.Host,
		"port":     fmt.Sprintf("%d", params.Port),
	}

	var token string
	err := r.client.get("/Connect", queryParams, &token)
	if err != nil {
		return "", err
	}

	r.client.SetToken(token)
	return token, nil
}

// ConnectEx เชื่อมต่อแบบ Extended (ใช้ชื่อ server แทน host/port)
func (r *ConnectionService) ConnectEx(user int64, password, server string) (string, error) {
	queryParams := map[string]string{
		"user":     fmt.Sprintf("%d", user),
		"password": password,
		"server":   server,
	}

	var token string
	err := r.client.get("/ConnectEx", queryParams, &token)
	if err != nil {
		return "", err
	}

	r.client.SetToken(token)
	return token, nil
}

// ConnectProxy เชื่อมต่อผ่าน Proxy
func (r *ConnectionService) ConnectProxy(params ConnectParams, proxyType, proxyHost string, proxyPort int) (string, error) {
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
	err := r.client.get("/ConnectProxy", queryParams, &token)
	if err != nil {
		return "", err
	}

	r.client.SetToken(token)
	return token, nil
}

// Disconnect ตัดการเชื่อมต่อ
func (r *ConnectionService) Disconnect() error {
	if r.client.token == "" {
		return nil
	}

	queryParams := map[string]string{
		"id": r.client.token,
	}

	var result string
	err := r.client.get("/Disconnect", queryParams, &result)
	if err != nil {
		return err
	}

	r.client.SetToken("")
	return nil
}

// IsConnected ตรวจสอบสถานะการเชื่อมต่อ
func (r *ConnectionService) IsConnected() (bool, error) {
	if r.client.token == "" {
		return false, nil
	}

	queryParams := map[string]string{
		"id": r.client.token,
	}

	var result string
	err := r.client.get("/CheckConnect", queryParams, &result)
	if err != nil {
		return false, err
	}

	return result == "OK", nil
}
