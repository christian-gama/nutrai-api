package value

// BMI represents a patient BMI, which is a ratio of the weight of a person to the square of the height.
type BMI float32

// Float32.
func (w BMI) Float32() float32 {
	return float32(w)
}

// Calculate calculates the BMI based on the weight and height.
func (w BMI) Calculate(weightKG WeightKG, heightM HeightM) BMI {
	return BMI(weightKG.Float32() / (heightM.Float32() * heightM.Float32()))
}
