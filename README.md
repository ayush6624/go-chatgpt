# Go-ChatGPT

Go-ChatGPT is an open-source GoLang client for ChatGPT, a large language model trained by OpenAI. With Go-ChatGPT, you can quickly and easily integrate ChatGPT's language processing capabilities into your Go applications.

## Features

- [x] Easy-to-use GoLang client for ChatGPT
- [x] Sends text to ChatGPT and receives a response
- [x] Support custom model and parameters
- [x] Supports GPT3.5 and GPT4 models


## Installation

You can install ChatGPT-Go using Go modules:

```bash
go get github.com/ayush6624/go-chatgpt
```

## Getting Started
Get your API key from the OpenAI Dashboard - [https://platform.openai.com/account/api-keys](https://platform.openai.com/account/api-keys) and export this either as an environment variable, or put this your `.bashrc` or `.zshrc`
```bash
export OPENAI_KEY=sk...
```

___

1. In your Go code, import the ChatGPT-Go package:
    ```go
    import (
        "github.com/ayush6624/go-chatgpt"
    )
    ```

2. Create a new ChatGPT client with your API key
    ```go
    key := os.Getenv("OPENAI_KEY")

    client, err := chatgpt.NewClient(key)
	if err != nil {
		log.Fatal(err)
	}
    ```
3. Use the `SimpleSend` API to send text to ChatGPT and get a response.
   ```go
    ctx := context.Background()

    res, err := c.SimpleSend(ctx, "Hello, how are you?")
	if err != nil {
		// handle error
	}
   ```
   The SimpleSend method sends the specified text to ChatGPT and returns a response. If an error occurs, it returns an error message.
4. To use a custom model/parameters, use the `Send` API.
   ```go
    ctx := context.Background()

    res, err = c.Send(ctx, &chatgpt.ChatCompletionRequest{
		Model: chatgpt.GPT35Turbo,
		Messages: []chatgpt.ChatMessage{
			{
				Role: chatgpt.ChatGPTModelRoleSystem,
				Content: "Hey, Explain GoLang to me in 2 sentences.",
			},
		},
	})
	if err != nil {
		// handle error
	}
   ```
## Contribute
If you want to contribute to this project, feel free to open a PR or an issue.


## License
This package is licensed under MIT license. See [LICENSE](./LICENSE) for details.


___
[![codecov](https://codecov.io/gh/ayush6624/go-chatgpt/branch/main/graph/badge.svg?token=2VXFP3238M)](https://codecov.io/gh/ayush6624/go-chatgpt)
[![Go](https://github.com/ayush6624/go-chatgpt/actions/workflows/go.yml/badge.svg)](https://github.com/ayush6624/go-chatgpt/actions/workflows/go.yml)
