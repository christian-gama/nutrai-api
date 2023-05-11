package command

// UpdatePatientInput represents the input data for the UpdatePatient command.
type UpdatePatientInput struct {
	ID uint `form:"id"`

	Age      int              `json:"age" faker:"boundary_start=18,boundary_end=100"`
	HeightM  float32          `json:"heightM" faker:"boundary_start=1,boundary_end=2"`
	WeightKG float32          `json:"weightKG" faker:"boundary_start=30,boundary_end=100"`
	User     *UpdateUserInput `json:"user" faker:"-"`
}
