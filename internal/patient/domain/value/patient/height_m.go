package value

// HeightM is a data type that represents a patient's height in meters. It is an essential
// factor in calculating BMI and determining health and diet plans.
type HeightM float32

// Float32 returns the float32 representation of the height.
func (w HeightM) Float32() float32 {
	return float32(w)
}
