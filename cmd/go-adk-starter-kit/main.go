package main

import (
	"context"
	"fmt"
	"os"
	"strings"

	_ "github.com/joho/godotenv/autoload"
	"google.golang.org/adk/agent"
	"google.golang.org/adk/model/gemini"
	"google.golang.org/adk/runner"
	"google.golang.org/adk/session"
	"google.golang.org/genai"

	"github.com/sanjayshr/go-adk-starter-kit/internal/agents/blog"
	"github.com/sanjayshr/go-adk-starter-kit/internal/config"
	"github.com/sanjayshr/go-adk-starter-kit/internal/logger"
)

func main() {
	ctx := context.Background()

	// Parse configuration flags
	cfg := config.ParseFlags()

	// Initialize structured logger with configured level
	logLevel := cfg.GetLogLevel()
	log := logger.New(os.Stdout, logLevel)
	log.Info("Starting ADK application", "logLevel", cfg.LogLevel, "agentLogger", cfg.AgentLogger)
	log.Debug("Debug logging enabled")

	// Create Gemini model
	apiKey := config.GetAPIKey()
	if apiKey == "" {
		log.Error("GOOGLE_API_KEY environment variable not set")
		os.Exit(1)
	}

	model, err := gemini.NewModel(ctx, "gemini-2.5-flash-lite", &genai.ClientConfig{
		APIKey: apiKey,
	})
	if err != nil {
		log.Error("Failed to create model", "error", err)
		os.Exit(1)
	}

	// Build the blog agent
	log.Info("Building blog agent pipeline")
	blogAgent, err := blog.Build(model)
	if err != nil {
		log.Error("Failed to build agent", "error", err)
		os.Exit(1)
	}

	// Create session service and runner
	sessionService := session.InMemoryService()
	r, err := runner.New(runner.Config{
		AppName:        "adk_starter_kit",
		Agent:          blogAgent,
		SessionService: sessionService,
	})
	if err != nil {
		log.Error("Failed to create runner", "error", err)
		os.Exit(1)
	}

	// Create session
	sess, err := sessionService.Create(ctx, &session.CreateRequest{
		AppName: "adk_starter_kit",
		UserID:  "demo_user",
	})
	if err != nil {
		log.Error("Failed to create session", "error", err)
		os.Exit(1)
	}

	// Get prompt from flag or use default
	prompt := cfg.Prompt
	if prompt == "" {
		prompt = blog.DefaultPrompt()
	}

	log.Info("Running blog pipeline", "prompt", prompt)
	log.Debug("Session details", "sessionID", sess.Session.ID(), "userID", "demo_user")
	fmt.Printf("\n%s\n", strings.Repeat("=", 80))
	fmt.Printf("Blog Pipeline\n")
	fmt.Printf("%s\n", strings.Repeat("=", 80))
	fmt.Printf("> %s\n\n", prompt)

	// Run the agent
	eventIter := r.Run(
		ctx,
		"demo_user",
		sess.Session.ID(),
		genai.NewContentFromText(prompt, genai.RoleUser),
		agent.RunConfig{
			StreamingMode: agent.StreamingModeNone,
		},
	)

	// Process events
	for event, err := range eventIter {
		if err != nil {
			log.Error("Error during agent execution", "error", err)
			return
		}

		if event == nil {
			log.Debug("Received nil event")
			continue
		}

		if event.Author != "" {
			fmt.Printf("Agent: %s\n", event.Author)
			log.Debug("Event from agent", "author", event.Author)
		}

		if event.Content != nil && len(event.Content.Parts) > 0 && event.Content.Parts[0].Text != "" {
			fmt.Printf("Response:\n%s\n", event.Content.Parts[0].Text)
			log.Debug("Response content length", "length", len(event.Content.Parts[0].Text))
		}
	}

	// Log agent outputs if enabled
	if cfg.AgentLogger {
		log.Info("Agent logger enabled, retrieving session outputs")
		agentLogger := logger.NewAgentLogger(sessionService, "adk_starter_kit", "demo_user", log)
		agentLogger.LogBlogOutputs(ctx, sess.Session.ID())
	} else {
		log.Info("Agent logger disabled, skipping session output logging")
	}

	fmt.Printf("\n%s\n", strings.Repeat("=", 80))
	fmt.Println("Completed!")
	fmt.Printf("%s\n\n", strings.Repeat("=", 80))
	log.Info("Application finished successfully")
}

