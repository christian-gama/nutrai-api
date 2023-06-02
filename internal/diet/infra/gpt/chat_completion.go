package gpt

import (
	"context"

	"github.com/christian-gama/nutrai-api/internal/diet/domain/repo"
)

type ChatCompletionOutput struct {
	Content string `json:"content"`
}

type ChatCompletion struct {
	Config repo.ChatCompletionConfigInput
	Client repo.GptClient
}

func NewChatCompletion(
	config repo.ChatCompletionConfigInput,
	client repo.GptClient,
) *ChatCompletion {
	return &ChatCompletion{
		Config: config,
		Client: client,
	}
}

func (c *ChatCompletion) Execute(
	ctx context.Context,
	messages []repo.ChatCompletionMessage,
) (*ChatCompletionOutput, error) {

	resp, err := c.Client.CreateChatCompletion(
		context.Background(),
		repo.ChatCompletionInput{
			Messages: messages,
			Config: repo.ChatCompletionConfigInput{
				Model:            c.Config.Model,
				Temperature:      c.Config.Temperature,
				TopP:             c.Config.TopP,
				PresencePenalty:  c.Config.PresencePenalty,
				FrequencyPenalty: c.Config.FrequencyPenalty,
				MaxTokens:        c.Config.MaxTokens,
				N:                c.Config.N,
				Stop:             c.Config.Stop,
				LogitBias:        c.Config.LogitBias,
				User:             c.Config.User,
			},
		},
	)

	if err != nil {
		return nil, err
	}

	return &ChatCompletionOutput{
		Content: resp.Choices[0].Message.Content,
	}, nil

}
