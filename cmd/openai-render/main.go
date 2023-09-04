// Command line tool to render a GPT-3 prompt result into files
package main

import (
	"context"
	"fmt"
	"io"
	"os"
	"strings"

	openai "github.com/sashabaranov/go-openai"
)

const openaiKeyEnv = "OPENAI_API_KEY"

func main() {
	apiKey := os.Getenv(openaiKeyEnv)
	if apiKey == "" {
		panic(fmt.Sprintf("%v is not set", openaiKeyEnv))
	}

	client := openai.NewClient(apiKey)
	resp, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: openai.GPT3Dot5Turbo,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleUser,
					Content: stdin(),
				},
			},
		},
	)

	if err != nil {
		fmt.Printf("ChatCompletion error: %v\n", err)
		return
	}

	// files := filesFromChatCompletion(resp)
	fmt.Println(resp.Choices[0].Message.Content)
}

func stdin() string {
	buf, err := io.ReadAll(os.Stdin)
	if err != nil {
		panic(err)
	}
	return strings.TrimSuffix(string(buf), "\n")
}
