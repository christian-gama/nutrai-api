package value

// MonthlyCostUSD represents the monthly cost of a diet in USD.
type MonthlyCostUSD float32

// Float32 returns the float32 representation of the monthly cost.
func (c MonthlyCostUSD) Float32() float32 {
	return float32(c)
}

// IsValid returns true if the monthly cost is valid.
func (c MonthlyCostUSD) IsValid() bool {
	return c > 0 && c < 10_000
}
