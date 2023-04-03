package main

import (
	"context"
	"embed"
	"fmt"
	"github.com/sashabaranov/go-openai"
	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
)

//go:embed all:frontend/dist
var assets embed.FS

func (a *App) SendMessage(messages []openai.ChatCompletionMessage) (string, error) {
	fmt.Println("Request Messages:")
	for _, msg := range messages {
		fmt.Printf("Role: %s, Content: %s\n", msg.Role, msg.Content)
	}

	resp, err := a.client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model:    openai.GPT3Dot5Turbo,
			Messages: messages,
		},
	)

	if err != nil {
		return "", err
	}

	fmt.Printf("Response Message: %s\n", resp.Choices[0].Message.Content)

	return resp.Choices[0].Message.Content, nil
}

func main() {
	// Create an instance of the app structure
	app := NewApp()

	// Create application with options
	err := wails.Run(&options.App{
		Title:  "My OpenAI Client",
		Width:  1024,
		Height: 768,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		BackgroundColour: &options.RGBA{R: 27, G: 38, B: 54, A: 1},
		OnStartup:        app.startup,
		Bind: []interface{}{
			app,
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}
}
