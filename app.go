package main

import (
	"context"
	"fmt"
	"github.com/sashabaranov/go-openai"
	"os"
)

// App struct
type App struct {
	ctx    context.Context
	client *openai.Client
}

// NewApp creates a new App application struct
func NewApp() *App {
	apiKey := os.Getenv("OPENAI_API_KEY")
	if apiKey == "" {
		panic("OPENAI_API_KEY environment variable not set")
	}

	return &App{
		client: openai.NewClient(apiKey),
	}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

// Greet returns a greeting for the given name
func (a *App) Greet(name string) string {
	return fmt.Sprintf("Hello %s, It's show time!", name)
}
