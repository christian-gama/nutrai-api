package value

// DurationInWeeks represents the duration of a diet in weeks.
type DurationInWeeks int16

// Int16 returns the int16 representation of the duration.
func (d DurationInWeeks) Int16() int16 {
	return int16(d)
}

// IsValid returns true if the duration is valid.
func (d DurationInWeeks) IsValid() bool {
	return d > 0 && d < 100
}
