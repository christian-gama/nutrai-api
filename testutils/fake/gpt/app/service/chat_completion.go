package fake

import (
	"github.com/christian-gama/nutrai-api/internal/gpt/app/service"
	. "github.com/christian-gama/nutrai-api/testutils/fake"
	"github.com/go-faker/faker/v4"
)

func ChatCompletionInput() *service.ChatCompletionInput {
	data := new(service.ChatCompletionInput)

	err := faker.FakeData(data)
	if err != nil {
		ErrGenerating(err)
	}

	return data
}

func ChatCompletionOutput() *service.ChatCompletionOutput {
	data := new(service.ChatCompletionOutput)

	err := faker.FakeData(data)
	if err != nil {
		ErrGenerating(err)
	}

	return data
}
