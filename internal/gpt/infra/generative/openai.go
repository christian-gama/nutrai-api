package generative

import (
	"context"

	gpt "github.com/christian-gama/nutrai-api/internal/gpt/domain/model/message"
	value "github.com/christian-gama/nutrai-api/internal/gpt/domain/value/message"
	"github.com/sashabaranov/go-openai"
)

type Generative struct {
	Client *openai.Client
}

func NewGenerative(client *openai.Client) *Generative {
	return &Generative{
		Client: client,
	}
}

func (g *Generative) ChatCompletion(
	ctx context.Context,
	input []*gpt.Message,
) (*gpt.Message, error) {

	var messages []openai.ChatCompletionMessage

	for _, message := range input {
		messages = append(messages, openai.ChatCompletionMessage{
			Role:    message.Role.String(),
			Content: message.Content.String(),
		})
	}

	model := input[0].Model

	stop := []string{}

	for _, s := range model.Stop {
		stop = append(stop, s.String())
	}

	resp, err := g.Client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model:            model.Name.String(),
			Messages:         messages,
			MaxTokens:        model.MaxTokens.Int(),
			Temperature:      model.Temperature.Float32(),
			TopP:             model.TopP.Float32(),
			PresencePenalty:  model.PresencePenalty.Float32(),
			FrequencyPenalty: model.FrequencyPenalty.Float32(),
			Stop:             stop,
		},
	)

	if err != nil {
		return nil, err
	}

	content := resp.Choices[0].Message.Content

	return &gpt.Message{
		Role:    value.Assistant,
		Content: value.Content(content),
	}, nil
}
