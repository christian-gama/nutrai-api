package service

// HashPasswordInput represents the input data for the HashPassword service.
type HashPasswordInput struct {
	Password string `json:"password" faker:"len=8"`
}

// HashPasswordOutput represents the output data for the HashPassword service.
type HashPasswordOutput struct {
	HashedPassword string `json:"hashedPassword"`
}
