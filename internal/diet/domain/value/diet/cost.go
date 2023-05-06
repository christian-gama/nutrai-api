package value

import (
	"fmt"

	"github.com/christian-gama/nutrai-api/pkg/errutil"
)

// MonthlyCostUSD represents the monthly cost of a diet in USD.
type MonthlyCostUSD float32

// Float32 returns the float32 representation of the monthly cost.
func (c MonthlyCostUSD) Float32() float32 {
	return float32(c)
}

// Validate returns an error if the monthly cost is invalid.
func (c MonthlyCostUSD) Validate() error {
	const fieldName = "monthlyCostUSD"
	const maxCost = 9_999

	if c == 0 {
		return errutil.NewErrRequired(fieldName)
	}

	if c < 0 {
		return errutil.NewErrInvalid(fieldName, "cannot be negative")
	}

	if c > maxCost {
		return errutil.NewErrInvalid(fieldName, fmt.Sprintf("cannot be greater than %d", maxCost))
	}

	return nil
}
