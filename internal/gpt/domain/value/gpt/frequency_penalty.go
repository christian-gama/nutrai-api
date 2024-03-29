package value

// FrequencyPenalty is a data type that represents the frequency penalty.
type FrequencyPenalty float32

// Float32 returns the string representation of the FrequencyPenalty.
func (n FrequencyPenalty) Float32() float32 {
	return float32(n)
}
