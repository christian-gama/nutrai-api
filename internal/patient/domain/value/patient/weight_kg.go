package value

// WeightKG is a data type that denotes a patient's weight in kilograms. It is a vital measure in
// health assessments and planning diet routines.
type WeightKG float32

// Float32 returns the float32 representation of the weight.
func (w WeightKG) Float32() float32 {
	return float32(w)
}
