package mt5client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"
)

// Client โครงสร้างหลักสำหรับเชื่อมต่อกับ MT5 REST API
type Client struct {
	baseURL    string
	token      string
	httpClient *http.Client

	// Services
	Connection   *ConnectionService
	Account      *AccountService
	Trading      *TradingService
	Order        *OrderService
	History      *HistoryService
	Quote        *QuoteService
	Symbol       *SymbolService
	Price        *PriceService
	Stats        *StatsService
	Subscription *SubscriptionService
	Service      *ServiceFunctions
}

// NewClient สร้าง Client ใหม่
func NewClient(baseURL string) *Client {
	if baseURL == "" {
		baseURL = "http://localhost:5000"
	}

	c := &Client{
		baseURL: baseURL,
		httpClient: &http.Client{
			Timeout: 30 * time.Second,
		},
	}

	c.Connection = &ConnectionService{client: c}
	c.Account = &AccountService{client: c}
	c.Trading = &TradingService{client: c}
	c.Order = &OrderService{client: c}
	c.History = &HistoryService{client: c}
	c.Quote = &QuoteService{client: c}
	c.Symbol = &SymbolService{client: c}
	c.Price = &PriceService{client: c}
	c.Stats = &StatsService{client: c}
	c.Subscription = &SubscriptionService{client: c}
	c.Service = &ServiceFunctions{client: c}

	return c
}

// SetToken ตั้งค่า token
func (r *Client) SetToken(token string) {
	r.token = token
}

// GetToken รับค่า token
func (r *Client) GetToken() string {
	return r.token
}

// doRequest ส่ง HTTP request
func (r *Client) doRequest(method, endpoint string, params map[string]string, body interface{}, result interface{}) error {
	fullURL := r.baseURL + endpoint

	if len(params) > 0 {
		values := url.Values{}
		for k, v := range params {
			values.Add(k, v)
		}
		fullURL += "?" + values.Encode()
	}

	var bodyReader io.Reader
	if body != nil {
		jsonData, err := json.Marshal(body)
		if err != nil {
			return fmt.Errorf("failed to marshal request body: %w", err)
		}
		bodyReader = bytes.NewBuffer(jsonData)
	}

	req, err := http.NewRequest(method, fullURL, bodyReader)
	if err != nil {
		return fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := r.httpClient.Do(req)
	if err != nil {
		return fmt.Errorf("failed to send request: %w", err)
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("failed to read response: %w", err)
	}

	if resp.StatusCode != 200 {
		return fmt.Errorf("unexpected status %d: %s", resp.StatusCode, string(respBody))
	}

	if result != nil {
		if strResult, ok := result.(*string); ok {
			*strResult = string(respBody)
			return nil
		}

		if err := json.Unmarshal(respBody, result); err != nil {
			return fmt.Errorf("failed to parse response: %w, body: %s", err, string(respBody))
		}
	}

	return nil
}

// get ส่ง GET request
func (r *Client) get(endpoint string, params map[string]string, result interface{}) error {
	return r.doRequest("GET", endpoint, params, nil, result)
}

// post ส่ง POST request
func (r *Client) post(endpoint string, params map[string]string, body interface{}, result interface{}) error {
	return r.doRequest("POST", endpoint, params, body, result)
}
