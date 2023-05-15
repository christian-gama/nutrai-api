package value

// BMI, or Body Mass Index, is a data type that represents a patient's body mass index.
// It is a key indicator of a patient's health status.
type BMI float32

// Float32.
func (w BMI) Float32() float32 {
	return float32(w)
}

// Calculate calculates the BMI based on the weight and height.
func (w BMI) Calculate(weightKG WeightKG, heightM HeightM) BMI {
	return BMI(weightKG.Float32() / (heightM.Float32() * heightM.Float32()))
}
