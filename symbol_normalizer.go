package mt5client

import (
	"fmt"
	"strings"
	"sync"
	"time"
)

// SymbolNormalizer จัดการการแปลง symbol ให้ตรงกับโบรกเกอร์
type SymbolNormalizer struct {
	client        *Client
	symbolCache   []string            // cache รายการ symbol ทั้งหมด
	symbolMap     map[string]string   // map: normalized → actual
	cacheTime     time.Time           // เวลาที่ cache ล่าสุด
	cacheDuration time.Duration       // ระยะเวลา cache (default: 1 ชั่วโมง)
	mu            sync.RWMutex        // lock สำหรับ concurrent access
	aliasMap      map[string][]string // alias mappings
}

// NewSymbolNormalizer สร้าง normalizer ใหม่
func (c *Client) NewSymbolNormalizer() *SymbolNormalizer {
	return &SymbolNormalizer{
		client:        c,
		symbolMap:     make(map[string]string),
		cacheDuration: time.Hour, // cache 1 ชั่วโมง
		aliasMap:      getDefaultAliasMap(),
	}
}

// getDefaultAliasMap สร้าง alias mappings
func getDefaultAliasMap() map[string][]string {
	return map[string][]string{
		// Gold
		"XAUUSD": {"GOLD", "XAUUSD", "GOLD."},
		"GOLD":   {"XAUUSD", "GOLD", "XAUUSD."},

		// Silver
		"XAGUSD": {"SILVER", "XAGUSD", "SILVER."},
		"SILVER": {"XAGUSD", "SILVER", "XAGUSD."},

		// Bitcoin (บางโบรกเกอร์)
		"BTCUSD": {"BITCOIN", "BTCUSD", "BTC"},

		// Ethereum
		"ETHUSD": {"ETHEREUM", "ETHUSD", "ETH"},

		// Oil
		"USOIL": {"CL", "WTI", "USOIL", "CRUDE"},
		"UKOIL": {"BRENT", "UKOIL"},

		// Indices
		"US30":   {"DOW", "US30", "DOWJONES"},
		"NAS100": {"NASDAQ", "NAS100", "NDX"},
		"SPX500": {"SP500", "SPX500", "US500"},
		"GER40":  {"DAX", "GER40", "DE40"},
	}
}

// SetCacheDuration กำหนดระยะเวลา cache
func (sn *SymbolNormalizer) SetCacheDuration(duration time.Duration) {
	sn.mu.Lock()
	defer sn.mu.Unlock()
	sn.cacheDuration = duration
}

// refreshCache อัพเดท cache symbol list
func (sn *SymbolNormalizer) refreshCache() error {
	sn.mu.Lock()
	defer sn.mu.Unlock()

	// ดึงรายการ symbol ทั้งหมด
	symbols, err := sn.client.Symbol.GetList()
	if err != nil {
		return fmt.Errorf("failed to get symbol list: %w", err)
	}

	sn.symbolCache = symbols
	sn.cacheTime = time.Now()

	// Clear map เดิม
	sn.symbolMap = make(map[string]string)

	return nil
}

// isCacheValid เช็คว่า cache ยังใช้ได้อยู่หรือไม่
func (sn *SymbolNormalizer) isCacheValid() bool {
	sn.mu.RLock()
	defer sn.mu.RUnlock()

	if len(sn.symbolCache) == 0 {
		return false
	}

	return time.Since(sn.cacheTime) < sn.cacheDuration
}

// Normalize แปลง symbol เป็นชื่อที่ใช้จริงในโบรกเกอร์
func (sn *SymbolNormalizer) Normalize(inputSymbol string) (string, error) {
	// ทำให้เป็นตัวพิมพ์ใหญ่
	inputSymbol = strings.ToUpper(strings.TrimSpace(inputSymbol))

	if inputSymbol == "" {
		return "", fmt.Errorf("input symbol is empty")
	}

	// 1. เช็คใน cache map ก่อน (fastest)
	sn.mu.RLock()
	if cached, found := sn.symbolMap[inputSymbol]; found {
		sn.mu.RUnlock()
		return cached, nil
	}
	sn.mu.RUnlock()

	// 2. Refresh cache ถ้าหมดอายุ
	if !sn.isCacheValid() {
		if err := sn.refreshCache(); err != nil {
			return "", err
		}
	}

	// 3. ค้นหา symbol
	result, err := sn.findSymbol(inputSymbol)
	if err != nil {
		return "", err
	}

	// 4. เก็บลง cache
	sn.mu.Lock()
	sn.symbolMap[inputSymbol] = result
	sn.mu.Unlock()

	return result, nil
}

// findSymbol ค้นหา symbol จาก input
func (sn *SymbolNormalizer) findSymbol(input string) (string, error) {
	sn.mu.RLock()
	symbolList := sn.symbolCache
	sn.mu.RUnlock()

	// Strategy 1: Exact match
	for _, symbol := range symbolList {
		if strings.EqualFold(symbol, input) {
			return symbol, nil
		}
	}

	// Strategy 2: Prefix match (สำหรับ suffix)
	// EURUSD → EURUSD.c, EURUSD#, EURUSD.iux
	matches := []string{}
	for _, symbol := range symbolList {
		// เช็คว่าขึ้นต้นด้วย input และตามด้วย suffix
		if strings.HasPrefix(strings.ToUpper(symbol), input) {
			// ตรวจสอบว่าส่วนที่เหลือเป็น suffix จริงๆ (ไม่ใช่ symbol อื่น)
			remaining := strings.ToUpper(symbol)[len(input):]
			if isSuffix(remaining) {
				matches = append(matches, symbol)
			}
		}
	}

	if len(matches) == 1 {
		return matches[0], nil
	} else if len(matches) > 1 {
		// เจอหลายตัว เลือกที่สั้นที่สุด (มี suffix น้อยที่สุด)
		shortest := matches[0]
		for _, m := range matches[1:] {
			if len(m) < len(shortest) {
				shortest = m
			}
		}
		return shortest, nil
	}

	// Strategy 3: Alias matching
	// XAUUSD → GOLD, GOLD → XAUUSD
	if aliases, found := sn.aliasMap[input]; found {
		for _, alias := range aliases {
			// ลอง exact match กับ alias
			for _, symbol := range symbolList {
				if strings.EqualFold(symbol, alias) {
					return symbol, nil
				}
			}

			// ลอง prefix match กับ alias
			for _, symbol := range symbolList {
				upperSymbol := strings.ToUpper(symbol)
				if strings.HasPrefix(upperSymbol, strings.ToUpper(alias)) {
					remaining := upperSymbol[len(alias):]
					if isSuffix(remaining) {
						return symbol, nil
					}
				}
			}
		}
	}

	// Strategy 4: Fuzzy matching (contains)
	// สำหรับกรณีพิเศษ เช่น GOLD.iux มี input เป็น GOLD
	for _, symbol := range symbolList {
		upperSymbol := strings.ToUpper(symbol)
		if strings.Contains(upperSymbol, input) {
			// ตรวจสอบว่าเป็นส่วนหนึ่งของ symbol จริงๆ
			if strings.HasPrefix(upperSymbol, input) || strings.HasSuffix(upperSymbol, input) {
				return symbol, nil
			}
		}
	}

	return "", fmt.Errorf("symbol '%s' not found in broker", input)
}

// isSuffix เช็คว่า string เป็น suffix ของโบรกเกอร์หรือไม่
func isSuffix(s string) bool {
	if s == "" {
		return false
	}

	// suffix มักขึ้นต้นด้วย . # _ -
	firstChar := s[0]
	if firstChar == '.' || firstChar == '#' || firstChar == '_' || firstChar == '-' {
		return true
	}

	// หรือเป็นชื่อสั้นๆ เช่น "c", "pro", "ecn"
	s = strings.ToLower(s)
	commonSuffixes := []string{"c", "pro", "ecn", "raw", "std", "prime", "iux", "uk", "us"}
	for _, suffix := range commonSuffixes {
		if s == suffix {
			return true
		}
	}

	return false
}

// NormalizeMany แปลงหลาย symbols พร้อมกัน
func (sn *SymbolNormalizer) NormalizeMany(inputSymbols []string) (map[string]string, error) {
	result := make(map[string]string)

	for _, input := range inputSymbols {
		normalized, err := sn.Normalize(input)
		if err != nil {
			result[input] = "" // ไม่เจอ
		} else {
			result[input] = normalized
		}
	}

	return result, nil
}

// GetAvailableSymbols ดึงรายการ symbol ทั้งหมดที่มี
func (sn *SymbolNormalizer) GetAvailableSymbols() ([]string, error) {
	if !sn.isCacheValid() {
		if err := sn.refreshCache(); err != nil {
			return nil, err
		}
	}

	sn.mu.RLock()
	defer sn.mu.RUnlock()

	// Return copy
	symbols := make([]string, len(sn.symbolCache))
	copy(symbols, sn.symbolCache)

	return symbols, nil
}

// ClearCache ล้าง cache
func (sn *SymbolNormalizer) ClearCache() {
	sn.mu.Lock()
	defer sn.mu.Unlock()

	sn.symbolCache = nil
	sn.symbolMap = make(map[string]string)
	sn.cacheTime = time.Time{}
}

// AddAlias เพิ่ม alias mapping เอง
func (sn *SymbolNormalizer) AddAlias(normalized string, aliases []string) {
	sn.mu.Lock()
	defer sn.mu.Unlock()

	normalized = strings.ToUpper(normalized)
	sn.aliasMap[normalized] = aliases
}
