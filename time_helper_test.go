package mt5client

import (
	"encoding/json"
	"testing"
)

func TestParseTime(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
		wantErr  bool
	}{
		{
			name:     "MT5 format without timezone",
			input:    "2025-10-17T23:57:56",
			expected: "2025-10-17T23:57:56",
			wantErr:  false,
		},
		{
			name:     "ISO8601 with Z",
			input:    "2025-10-17T23:57:56Z",
			expected: "2025-10-17T23:57:56",
			wantErr:  false,
		},
		{
			name:     "ISO8601 with timezone",
			input:    "2025-10-17T23:57:56+07:00",
			expected: "2025-10-17T23:57:56",
			wantErr:  false,
		},
		{
			name:     "Date only",
			input:    "2025-10-17",
			expected: "2025-10-17T00:00:00",
			wantErr:  false,
		},
		{
			name:     "Empty string",
			input:    "",
			expected: "",
			wantErr:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := parseTime(tt.input)

			if (err != nil) != tt.wantErr {
				t.Errorf("parseTime() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !tt.wantErr && tt.expected != "" {
				formatted := result.Format("2006-01-02T15:04:05")
				if formatted != tt.expected {
					t.Errorf("Expected %s, got %s", tt.expected, formatted)
				}
			}
		})
	}
}

func TestQuoteUnmarshal(t *testing.T) {
	jsonData := `{
		"symbol": "EURUSD",
		"bid": 1.16504,
		"ask": 1.16522,
		"time": "2025-10-17T23:57:56",
		"last": 0.0,
		"volume": 0,
		"spread": 18
	}`

	var quote Quote
	err := json.Unmarshal([]byte(jsonData), &quote)
	if err != nil {
		t.Fatalf("Failed to unmarshal: %v", err)
	}

	if quote.Symbol != "EURUSD" {
		t.Errorf("Expected EURUSD, got %s", quote.Symbol)
	}

	if quote.Bid != 1.16504 {
		t.Errorf("Expected 1.16504, got %f", quote.Bid)
	}

	expectedTime := "2025-10-17T23:57:56"
	actualTime := quote.Time.Format("2006-01-02T15:04:05")
	if actualTime != expectedTime {
		t.Errorf("Expected %s, got %s", expectedTime, actualTime)
	}
}

func TestBarUnmarshal(t *testing.T) {
	jsonData := `{
		"time": "2025-10-17T23:57:00",
		"open": 1.16500,
		"high": 1.16550,
		"low": 1.16490,
		"close": 1.16520,
		"volume": 1000,
		"spread": 15
	}`

	var bar Bar
	err := json.Unmarshal([]byte(jsonData), &bar)
	if err != nil {
		t.Fatalf("Failed to unmarshal: %v", err)
	}

	if bar.Open != 1.16500 {
		t.Errorf("Expected 1.16500, got %f", bar.Open)
	}

	expectedTime := "2025-10-17T23:57:00"
	actualTime := bar.Time.Format("2006-01-02T15:04:05")
	if actualTime != expectedTime {
		t.Errorf("Expected %s, got %s", expectedTime, actualTime)
	}
}
