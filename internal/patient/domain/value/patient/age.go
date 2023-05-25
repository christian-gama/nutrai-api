package value

// Age is a data type that signifies a patient's age. It's crucial for determining
// appropriate health and diet plans.
type Age int8

// Int8 returns the int8 representation of the age.
func (w Age) Int8() int8 {
	return int8(w)
}
