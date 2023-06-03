package repo

import (
	"context"

	"github.com/christian-gama/nutrai-api/internal/gpt/domain/model/gpt"
)

type Generative interface {
	ChatCompletion(
		ctx context.Context,
		input []*gpt.Message,
	) (*gpt.Message, error)
}
