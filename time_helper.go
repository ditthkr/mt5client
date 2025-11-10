package mt5client

import (
	"strings"
	"time"
)

// parseTime แปลง string เป็น time.Time รองรับหลาย format
func parseTime(str string) (time.Time, error) {
	// ลบ quotes ถ้ามี
	str = strings.Trim(str, `"`)

	if str == "" || str == "null" {
		return time.Time{}, nil
	}

	// รายการ format ที่ต้องการรองรับ
	formats := []string{
		"2006-01-02T15:04:05",           // MT5 format (ไม่มี timezone)
		"2006-01-02T15:04:05Z",          // ISO8601 with Z
		"2006-01-02T15:04:05Z07:00",     // ISO8601 with timezone
		"2006-01-02T15:04:05.999999999", // with nanoseconds
		"2006-01-02 15:04:05",           // space separator
		"2006-01-02",                    // date only
	}

	var lastErr error
	for _, format := range formats {
		t, err := time.Parse(format, str)
		if err == nil {
			return t, nil
		}
		lastErr = err
	}

	// ถ้าไม่ตรงกับ format ไหนเลย
	return time.Time{}, lastErr
}
