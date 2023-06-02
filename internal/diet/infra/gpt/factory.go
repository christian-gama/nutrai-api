package gpt

import (
	"github.com/christian-gama/nutrai-api/config/env"
	"github.com/christian-gama/nutrai-api/internal/diet/domain/repo"
)

func MakeChatCompletion() *ChatCompletion {

	config := repo.ChatCompletionConfigInput{
		Model:            env.Gpt.Model,
		MaxTokens:        env.Gpt.MaxTokens,
		Temperature:      env.Gpt.Temperature,
		TopP:             env.Gpt.TopP,
		PresencePenalty:  env.Gpt.PresencePenalty,
		FrequencyPenalty: env.Gpt.FrequencyPenalty,
		N:                env.Gpt.N,
		Stop:             env.Gpt.Stop,
	}

	return NewChatCompletion()
}
