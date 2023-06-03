package value

// TopP is a data type that represents the top p of a model.
type TopP float32

// Float32 returns the string representation of the TopP.
func (n TopP) Float32() float32 {
	return float32(n)
}
