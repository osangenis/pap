// Command line tool to render a GPT-3 prompt result into files
package main

import (
	"context"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/osangenis/pap/v2/cmd/openai-render/response"
	openai "github.com/sashabaranov/go-openai"
)

const openaiKeyEnv = "OPENAI_API_KEY"

func main() {
	pOutputDir := flag.String("output_dir", "", "The directory in where files/code blocks from the response will be saved")
	flag.Parse()

	if pOutputDir == nil || *pOutputDir == "" {
		panic("The flag output_dir is required")
	}

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
		panic(fmt.Sprintf("ChatCompletion error: %v\n", err))
	}

	files := response.FilesFromChat(resp.Choices[0].Message.Content, "go")
	err = files.Write(*pOutputDir)
	if err != nil {
		panic(err)
	}

	fmt.Println(resp.Choices[0].Message.Content)
}

func stdin() string {
	buf, err := io.ReadAll(os.Stdin)
	if err != nil {
		panic(err)
	}
	return strings.TrimSuffix(string(buf), "\n")
}
