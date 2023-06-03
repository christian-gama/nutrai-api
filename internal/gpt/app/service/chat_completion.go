package service

import (
	"context"

	"github.com/christian-gama/nutrai-api/internal/core/domain/service"
	"github.com/christian-gama/nutrai-api/internal/gpt/domain/model/gpt"
	"github.com/christian-gama/nutrai-api/internal/gpt/domain/repo"
	value "github.com/christian-gama/nutrai-api/internal/gpt/domain/value/message"
)

type message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}
type ChatCompletionOutput struct {
	Message message `json:"message"`
}

type ChatCompletionInput struct {
	Messages []message `json:"messages"`
}

type ChatCompletion = service.Handler[*ChatCompletionInput, *ChatCompletionOutput]

type ChatCompletionImpl struct {
	Client repo.Generative
}

func NewChatCompletion(
	client repo.Generative,
) ChatCompletion {
	return &ChatCompletionImpl{
		Client: client,
	}
}

func (c *ChatCompletionImpl) Handle(
	ctx context.Context,
	input *ChatCompletionInput,
) (*ChatCompletionOutput, error) {
	var messages []*gpt.Message

	for _, message := range input.Messages {
		messages = append(messages, &gpt.Message{
			Role:    value.Role(message.Role),
			Content: value.Content(message.Content),
		})
	}

	resp, err := c.Client.ChatCompletion(
		context.Background(),
		&repo.ChatCompletionInput{
			Messages: messages,
			Model:    gpt.NewModel(),
		},
	)
	if err != nil {
		return nil, err
	}

	return &ChatCompletionOutput{
		Message: message{
			Role:    resp.Role.String(),
			Content: resp.Content.String(),
		},
	}, nil
}
