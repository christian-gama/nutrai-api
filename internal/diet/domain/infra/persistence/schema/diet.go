package schema

// Diet is the database schema for diets.
type Diet struct {
	ID              uint `gorm:"primaryKey"`
	Description     string
	RestrictedFood  []string
	DurationInWeeks int16
	Goal            string
	MealPlan        string
	MonthlyCostUSD  float64
}

// TableName returns the table name for the Diet schema.
func (d Diet) TableName() string {
	return "diets"
}
