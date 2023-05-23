package jwt

// TokenType represents the type of the JWT token, which can be either access or refresh.
type TokenType string

const (
	AccessTokenType  TokenType = "access"
	RefreshTokenType TokenType = "refresh"
)
