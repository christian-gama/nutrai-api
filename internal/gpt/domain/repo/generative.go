package repo

import (
	"context"

	"github.com/christian-gama/nutrai-api/internal/gpt/domain/model/gpt"
)

type ChatCompletionInput struct {
	Messages []*gpt.Message
	Model    *gpt.Model
}

type Generative interface {
	ChatCompletion(
		ctx context.Context,
		input *ChatCompletionInput,
	) (*gpt.Message, error)
}
