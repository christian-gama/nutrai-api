package generative

import (
	"github.com/christian-gama/nutrai-api/config/env"
	"github.com/sashabaranov/go-openai"
)

func MakeGenerative() *Generative {
	client := openai.NewClient(env.Gpt.ApiKey)
	return NewGenerative(client)
}
