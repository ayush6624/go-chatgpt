package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	"github.com/ayush6624/go-chatgpt"
)

func main() {
	key := os.Getenv("OPENAI_KEY")
	c, err := chatgpt.NewClient(key)
	if err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()
	res, err := c.SimpleSend(ctx, "Hey, Explain GoLang to me in 2 sentences.")
	if err != nil {
		log.Fatal(err)
	}

	a, _ := json.MarshalIndent(res, "", "  ");
	log.Println(string(a))

	res, err = c.Send(ctx, &chatgpt.ChatCompletionRequest{
		Model: chatgpt.GPT4,
		Messages: []chatgpt.ChatMessage{
			{
				Role: chatgpt.ChatGPTModelRoleSystem,
				Content: "Hey, Explain GoLang to me in 2 sentences.",
			},
		},
	})
	if err != nil {
		log.Fatal(err)
	}

	a, _ = json.MarshalIndent(res, "", "  ");
	log.Println(string(a))
}