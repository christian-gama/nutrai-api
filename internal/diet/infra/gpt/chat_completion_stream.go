package gpt

import (
	"context"
	"errors"
	"io"

	"github.com/christian-gama/nutrai-api/internal/diet/domain/repo"
)

type ChatCompletionStreamOutput struct {
	Content string `json:"content"`
}

type ChatCompletionStream struct {
	Config repo.ChatCompletionConfigInput
	Client repo.GptClient
}

func NewChatCompletionStream(
	config repo.ChatCompletionConfigInput,
	client repo.GptClient,
) *ChatCompletionStream {
	return &ChatCompletionStream{
		Config: config,
		Client: client,
	}
}

func (c *ChatCompletionStream) Execute(ctx context.Context, messages []repo.ChatCompletionMessage, stream chan ChatCompletionStreamOutput) error {

	resp, err := c.Client.CreateChatCompletionStream(
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
			},
		},
	)

	defer resp.Close()

	if err != nil {
		return err
	}

	for {
		response, err := resp.Read()

		if errors.Is(err, io.EOF) {
			break
		}

		if err != nil {
			return err
		}

		content := response.Choices[0].Delta.Content

		r := ChatCompletionStreamOutput{
			Content: content,
		}

		stream <- r
	}

	return nil

}
