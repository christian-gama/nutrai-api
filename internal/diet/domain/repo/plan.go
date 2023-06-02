package repo

import (
	"context"

	"github.com/christian-gama/nutrai-api/internal/core/domain/queryer"
	"github.com/christian-gama/nutrai-api/internal/core/domain/value"
	"github.com/christian-gama/nutrai-api/internal/diet/domain/model/plan"
)

// SavePlanInput is the input for the Save method.
type SavePlanInput struct {
	Plan *plan.Plan
}

// AllPlansInput is the input for the All method.
type AllPlansInput struct {
	queryer.Filterer
	queryer.Sorter
	queryer.Paginator
	queryer.Preloader
}

// FindPlanInput is the input for the Find method.
type FindPlanInput struct {
	ID value.ID
	queryer.Filterer
	queryer.Preloader
}

// DeletePlanInput is the input for the Delete method.
type DeletePlanInput struct {
	IDs []value.ID
}

// UpdatePlanInput is the input for the Update method.
type UpdatePlanInput struct {
	Plan *plan.Plan
	ID   value.ID
}

// Plan is the interface that wraps the basic Plan methods.
type Plan interface {
	// All returns all plans.
	All(ctx context.Context, input AllPlansInput) (*queryer.PaginationOutput[*plan.Plan], error)

	// Delete deletes the plan with the given id.
	Delete(ctx context.Context, input DeletePlanInput) error

	// Find returns the plan with the given id.
	Find(ctx context.Context, input FindPlanInput) (*plan.Plan, error)

	// Save saves the given plan.
	Save(ctx context.Context, input SavePlanInput) (*plan.Plan, error)

	// Update updates the given plan.
	Update(ctx context.Context, input UpdatePlanInput) error
}

// --- GPT ---

// ChatCompletionMessage represents a message to be sent to the OpenAI API.
type ChatCompletionMessage struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

// ChatCompletionConfigInput represents the configuration for the OpenAI API.
type ChatCompletionConfigInput struct {
	Model            string         `json:"model"`
	MaxTokens        int            `json:"max_tokens,omitempty"`
	Temperature      float32        `json:"temperature,omitempty"`
	TopP             float32        `json:"top_p,omitempty"`
	N                int            `json:"n,omitempty"`
	Stop             []string       `json:"stop,omitempty"`
	PresencePenalty  float32        `json:"presence_penalty,omitempty"`
	FrequencyPenalty float32        `json:"frequency_penalty,omitempty"`
	LogitBias        map[string]int `json:"logit_bias,omitempty"`
	User             string         `json:"user,omitempty"`
}

// ChatCompletionInput represents the input to the OpenAI API.
type ChatCompletionInput struct {
	Messages []ChatCompletionMessage `json:"messages"`
	Config   ChatCompletionConfigInput
}

// ChatCompletionChoice represents a choice from the OpenAI API.
type ChatCompletionChoice struct {
	Index        int                   `json:"index"`
	Message      ChatCompletionMessage `json:"message"`
	FinishReason string                `json:"finish_reason"`
}

// Usage Represents the total token usage per request to OpenAI.
type Usage struct {
	PromptTokens     int `json:"prompt_tokens"`
	CompletionTokens int `json:"completion_tokens"`
	TotalTokens      int `json:"total_tokens"`
}

// ChatCompletionOutput represents the output from the OpenAI API.
type ChatCompletionOutput struct {
	ID      string                 `json:"id"`
	Object  string                 `json:"object"`
	Created int64                  `json:"created"`
	Model   string                 `json:"model"`
	Choices []ChatCompletionChoice `json:"choices"`
	Usage   Usage                  `json:"usage"`
}

// --- GPT Stream ---

type ChatCompletionStreamChoice struct {
	Index        int                   `json:"index"`
	Delta        ChatCompletionMessage `json:"delta"`
	FinishReason string                `json:"finish_reason"`
}

type ChatCompletionStreamResponse struct {
	ID      string                       `json:"id"`
	Object  string                       `json:"object"`
	Created int64                        `json:"created"`
	Model   string                       `json:"model"`
	Choices []ChatCompletionStreamChoice `json:"choices"`
}

// ChatCompletionStream
// Note: Perhaps it is more elegant to abstract Stream using generics.
type ChatCompletionStream struct {
	streamReader
}

type streamReader interface {
	Read() (*ChatCompletionStreamResponse, error)
	Close() error
}

type GptClient interface {
	CreateChatCompletion(ctx context.Context, input ChatCompletionInput) (*ChatCompletionOutput, error)
	CreateChatCompletionStream(ctx context.Context, input ChatCompletionInput) (*ChatCompletionStream, error)
}
