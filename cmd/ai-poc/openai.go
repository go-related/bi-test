package main

import (
	"context"
	"fmt"
	"github.com/openai/openai-go"
	"github.com/openai/openai-go/option"
	"strings"
)

type AIClient struct {
	cl *openai.Client
}

func NewAIClient(apiKey string) *AIClient {
	client := openai.NewClient(
		option.WithAPIKey(apiKey),
	)
	return &AIClient{
		client,
	}
}

func (c *AIClient) Summarize(text string) (string, error) {
	if len(text) > 20000 {
		text = text[:20000]
	}

	chatCompletion, err := c.cl.Chat.Completions.New(context.Background(), openai.ChatCompletionNewParams{
		Messages: openai.F([]openai.ChatCompletionMessageParamUnion{
			openai.UserMessage(fmt.Sprintf("Create a summary for this investing pdf file: '%s'", text)),
		}),
		Model: openai.F(openai.ChatModelGPT4oMini),
	})
	if err != nil || chatCompletion == nil {
		return "summary not available", err
	}

	result := strings.Builder{}
	for _, ch := range chatCompletion.Choices {
		result.WriteString(ch.Message.Content)
	}

	return result.String(), nil
}
