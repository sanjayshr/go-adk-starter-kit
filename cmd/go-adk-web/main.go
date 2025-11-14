package main

import (
	"context"
	"log"
	"os"

	_ "github.com/joho/godotenv/autoload"
	"google.golang.org/adk/cmd/launcher/adk"
	"google.golang.org/adk/cmd/launcher/full"
	"google.golang.org/adk/model/gemini"
	"google.golang.org/adk/server/restapi/services"
	"google.golang.org/genai"

	"github.com/sanjayshr/go-adk-starter-kit/internal/agents/blog"
	"github.com/sanjayshr/go-adk-starter-kit/internal/config"
)

func main() {
	ctx := context.Background()

	// Get API key from environment
	apiKey := config.GetAPIKey()
	if apiKey == "" {
		log.Fatal("GOOGLE_API_KEY environment variable not set")
	}

	// Create Gemini model
	model, err := gemini.NewModel(ctx, "gemini-2.5-flash-lite", &genai.ClientConfig{
		APIKey: apiKey,
	})
	if err != nil {
		log.Fatalf("Failed to create model: %v", err)
	}

	// Build the blog agent
	blogAgent, err := blog.Build(model)
	if err != nil {
		log.Fatalf("Failed to build agent: %v", err)
	}

	// Configure launcher with agent
	launcherConfig := &adk.Config{
		AgentLoader: services.NewSingleAgentLoader(blogAgent),
	}

	// Create and execute launcher
	l := full.NewLauncher()
	if err = l.Execute(ctx, launcherConfig, os.Args[1:]); err != nil {
		log.Fatalf("Run failed: %v\n\n%s", err, l.CommandLineSyntax())
	}
}
