package gpt

import (
	"context"

	"github.com/christian-gama/nutrai-api/internal/diet/domain/repo"
)

type ChatCompletionConfigInput struct {
	Model                string
	Temperature          float32 // 0.0 to 1.0
	TopP                 float32 // 0.0 to 1.0 - to a low value, like 0.1, the model will be very conservative in its word choices, and will tend to generate relatively predictable prompts
	N                    int     // number of messages to generate
	PresencePenalty      float32 // -2.0 to 2.0 - Number between -2.0 and 2.0. Positive values penalize new tokens based on whether they appear in the text so far, increasing the model's likelihood to talk about new topics.
	FrequencyPenalty     float32 // -2.0 to 2.0 - Number between -2.0 and 2.0. Positive values penalize new tokens based on their existing frequency in the text so far, increasing the model's likelihood to talk about new topics.
	InitialSystemMessage string
}

type ChatCompletionInput struct {
	Message string `json:"message"`
}

type ChatCompletionOutput struct {
	Content string `json:"content"`
}

type ChatCompletion struct {
	Config ChatCompletionConfigInput
	Client repo.GptClient
}

func NewChatCompletion(
	config ChatCompletionConfigInput,
	client repo.GptClient,
) *ChatCompletion {
	return &ChatCompletion{
		Config: config,
		Client: client,
	}
}

func (c *ChatCompletion) Execute(ctx context.Context, input ChatCompletionInput) (*ChatCompletionOutput, error) {
	messages := []repo.ChatCompletionMessage{
		{
			Role:    "System",
			Content: c.Config.InitialSystemMessage,
		},
		{
			Role:    "User",
			Content: input.Message,
		},
	}

	resp, err := c.Client.CreateChatCompletion(
		context.Background(),
		repo.ChatCompletionConfigInput{
			Model:            c.Config.Model,
			Messages:         messages,
			Temperature:      c.Config.Temperature,
			TopP:             c.Config.TopP,
			PresencePenalty:  c.Config.PresencePenalty,
			FrequencyPenalty: c.Config.FrequencyPenalty,
		},
	)

	if err != nil {
		return nil, err
	}

	return &ChatCompletionOutput{
		Content: resp.Choices[0].Message.Content,
	}, nil

}
