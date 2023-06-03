package value

// PresencePenalty is a data type that represents the presence penalty.
type PresencePenalty float32

// Float32 returns the string representation of the PresencePenalty.
func (n PresencePenalty) Float32() float32 {
	return float32(n)
}
