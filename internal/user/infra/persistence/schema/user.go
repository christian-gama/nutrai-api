package schema

// User is the database schema for users.
type User struct {
	ID       uint `gorm:"primaryKey"`
	Email    string
	Name     string
	Password string
}

// TableName returns the table name for the User schema.
func (u User) TableName() string {
	return "users"
}
