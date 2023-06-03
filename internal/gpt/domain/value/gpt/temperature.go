package value

// Temperature is a data type that represents the temperature of a model.
type Temperature float32

// Float32 returns the string representation of the Temperature.
func (n Temperature) Float32() float32 {
	return float32(n)
}
