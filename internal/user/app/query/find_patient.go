package query

type FindPatientInput struct {
	ID uint `form:"id"`
}

type FindPatientOutput struct {
	ID       uint            `json:"id"`
	Age      int             `json:"age"`
	HeightM  float32         `json:"heightM"`
	WeightKG float32         `json:"weightKG"`
	User     *FindUserOutput `json:"user,omitempty"`
}
