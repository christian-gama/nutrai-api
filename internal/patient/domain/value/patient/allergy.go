package value

type Allergy string

func (a Allergy) String() string {
	return string(a)
}
