package value

import (
	"github.com/christian-gama/nutrai-api/pkg/errutil/errors"
)

// MonthlyCostUSD is a data type that denotes the monthly cost of a specific diet plan, measured in
// US dollars. It provides a financial aspect of the diet plan.
type MonthlyCostUSD float32

// Float32 returns the float32 representation of the monthly cost.
func (c MonthlyCostUSD) Float32() float32 {
	return float32(c)
}

// Validate returns an error if the monthly cost is invalid.
func (c MonthlyCostUSD) Validate() error {
	const fieldName = "MonthlyCostUSD"

	if c == 0 {
		return errors.Required(fieldName)
	}

	if c < 0 {
		return errors.Invalid(fieldName, "cannot be negative")
	}

	return nil
}
