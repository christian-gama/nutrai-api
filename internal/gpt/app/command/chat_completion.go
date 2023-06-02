package gpt

import (
	"context"

	gpt "github.com/christian-gama/nutrai-api/internal/gpt/domain/model/message"
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

type ChatCompletion struct {
	Client repo.Generative
}

func NewChatCompletion(
	client repo.Generative,
) *ChatCompletion {
	return &ChatCompletion{
		Client: client,
	}
}

func (c *ChatCompletion) Execute(
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
		messages,
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
