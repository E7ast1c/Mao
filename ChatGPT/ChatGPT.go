package ChatGPT

import (
	"context"
	"fmt"
	"time"

	goopenai "github.com/sashabaranov/go-openai"
)

type Client struct {
	client goopenai.Client
}

func New(token string) Client {
	return Client{client: *goopenai.NewClient(token)}
}

func (c Client) ChatCompletionRequest(ctx context.Context, content string) (string, error) {
	resp, err := c.client.CreateChatCompletion(
		ctx,
		goopenai.ChatCompletionRequest{
			Model: goopenai.GPT3Dot5Turbo,
			Messages: []goopenai.ChatCompletionMessage{
				{
					Role:    goopenai.ChatMessageRoleUser,
					Content: content,
				},
			},
		},
	)

	if err != nil {
		fmt.Printf("ChatCompletion error: [%v] %v\n", time.Now(), err)
		return "", err
	}

	respContent := resp.Choices[0].Message.Content
	fmt.Printf("Response: [%v] %s\n", time.Now(), respContent)

	return respContent, nil
}
