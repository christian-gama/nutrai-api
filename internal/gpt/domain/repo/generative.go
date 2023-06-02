package repo

import (
	"context"

	gpt "github.com/christian-gama/nutrai-api/internal/gpt/domain/model/message"
)

type Generative interface {
	ChatCompletion(
		ctx context.Context,
		input []*gpt.Message,
	) (*gpt.Message, error)
}
