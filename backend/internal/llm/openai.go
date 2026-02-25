package llm

import (
	"context"
	"encoding/json"
	"errors"

	"github.com/openai/openai-go/v3"
	"github.com/openai/openai-go/v3/responses"
)

// MockClient provides predictable responses for testing without an API key.
// It classifies tasks based on simple keyword matching in the title.
type OpenAIClient struct{}

func NewOpenAIClient() *OpenAIClient {
	return &OpenAIClient{}
}

// ClassifyTask sends the task to the OpenAI API and parses the response.
func (m *OpenAIClient) ClassifyTask(ctx context.Context, title string, description string) (*TaskClassification, error) {
	client := openai.NewClient()

	prompt := "Classify the following task with tags, priority (high, medium, low), category (bug, feature, improvement, research), and a one-line summary.\n\n" +
		"Title: " + title + "\nDescription: " + description + "\n\n" +
		"Respond in JSON with keys: tags, priority, category, summary."

	resp, err := client.Responses.New(ctx, responses.ResponseNewParams{
		Input: responses.ResponseNewParamsInputUnion{OfString: openai.String(prompt)},
		Model: openai.ChatModelGPT5Mini,
	})

	if err != nil {
		return nil, err
	}

	if resp.OutputText() == "" {
		return nil, errors.New("no response from OpenAI")
	}

	var result TaskClassification
	err = json.Unmarshal([]byte(resp.OutputText()), &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}
