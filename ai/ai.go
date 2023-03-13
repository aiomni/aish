package ai

import (
	"context"
	"fmt"
	"time"

	"github.com/aiomni/aish/config"
	"github.com/aiomni/aish/prompts"
	openai "github.com/sashabaranov/go-openai"
)

var client *openai.Client

func init() {
	apiKey := config.GetAPIKey()
	openAIconfig := openai.DefaultConfig(apiKey)
	openAIconfig.OrgID = config.GetOrganizationID()
	proxyDomain := config.GetProxyDomain()
	openAIconfig.BaseURL = fmt.Sprintf("https://%s/v1", proxyDomain)
	openAIconfig.HTTPClient.Timeout = 30 * time.Second

	client = openai.NewClientWithConfig(openAIconfig)
}

func AskChatGPT(content string) (string, error) {
	body := openai.ChatCompletionRequest{
		Model:           openai.GPT3Dot5Turbo,
		Temperature:     0.2,
		PresencePenalty: 0.5,
		Messages: []openai.ChatCompletionMessage{
			{
				Role:    openai.ChatMessageRoleSystem,
				Content: prompts.SYSTEM_BASH_EXPERT,
			},
			{
				Role:    openai.ChatMessageRoleUser,
				Content: content,
			},
		},
	}

	resp, err := client.CreateChatCompletion(context.Background(), body)
	if err != nil {
		fmt.Println("Error:", err)
		return "", err
	}

	return resp.Choices[0].Message.Content, nil
}

func MockAskChatGPT(content string) (string, error) {
	r := `{
    "code": 1,
    "command": "ls -lh",
    "note": "This command lists all files in the current directory with their sizes in a human-readable format.",
    "warning": "A test warning"
}`

	return r, nil
}
