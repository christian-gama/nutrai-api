package schema

// Patient is the database schema for a patient.
type Patient struct {
	ID       uint `gorm:"primaryKey"`
	UserID   uint
	User     User `gorm:"foreignKey:UserID"`
	WeightKG float32
	HeightM  float32
	Age      int8
}

// TableName returns the table name for the Patient schema.
func (u Patient) TableName() string {
	return "patients"
}
