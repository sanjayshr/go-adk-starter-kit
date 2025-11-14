// Package logger provides structured logging utilities using slog
package logger

import (
	"context"
	"fmt"
	"io"
	"log/slog"
	"strings"

	"google.golang.org/adk/session"
)

// New creates a new structured logger with slog
func New(w io.Writer, level slog.Level) *slog.Logger {
	opts := &slog.HandlerOptions{
		Level: level,
	}
	handler := slog.NewJSONHandler(w, opts)
	return slog.New(handler)
}

// AgentLogger handles logging of agent outputs from session state
type AgentLogger struct {
	sessionService session.Service
	appName        string
	userID         string
	logger         *slog.Logger
}

// NewAgentLogger creates a new agent logger
func NewAgentLogger(sessionService session.Service, appName, userID string, logger *slog.Logger) *AgentLogger {
	return &AgentLogger{
		sessionService: sessionService,
		appName:        appName,
		userID:         userID,
		logger:         logger,
	}
}

// LogBlogOutputs retrieves and logs blog pipeline outputs from session state
func (al *AgentLogger) LogBlogOutputs(ctx context.Context, sessionID string) {
	al.logger.Info("Retrieving blog pipeline outputs", "sessionID", sessionID)
	al.logger.Debug("Agent logger config", "appName", al.appName, "userID", al.userID)

	fmt.Printf("\n%s\n", strings.Repeat("-", 80))
	fmt.Println("Agent Outputs (from Session State):")
	fmt.Printf("%s\n", strings.Repeat("-", 80))

	resp, err := al.sessionService.Get(ctx, &session.GetRequest{
		AppName:   al.appName,
		UserID:    al.userID,
		SessionID: sessionID,
	})
	if err != nil {
		al.logger.Error("Could not retrieve session", "error", err, "sessionID", sessionID)
		return
	}

	al.logger.Debug("Session retrieved successfully", "sessionID", sessionID)

	// Log blog outline
	if outline, err := resp.Session.State().Get("blog_outline"); err == nil {
		fmt.Printf("\n[OutlineAgent] Output:\n%v\n", outline)
		al.logger.Info("Retrieved blog outline", "length", len(fmt.Sprint(outline)))
		al.logger.Debug("Blog outline content", "outline", outline)
	} else {
		al.logger.Error("Blog outline not found in session state", "error", err)
	}

	// Log blog draft
	if draft, err := resp.Session.State().Get("blog_draft"); err == nil {
		fmt.Printf("\n[WriterAgent] Output:\n%v\n", draft)
		al.logger.Info("Retrieved blog draft", "length", len(fmt.Sprint(draft)))
		al.logger.Debug("Blog draft content", "draft", draft)
	} else {
		al.logger.Error("Blog draft not found in session state", "error", err)
	}

	// Log final blog
	if final, err := resp.Session.State().Get("final_blog"); err == nil {
		fmt.Printf("\n[EditorAgent] Output:\n%v\n", final)
		al.logger.Info("Retrieved final blog", "length", len(fmt.Sprint(final)))
		al.logger.Debug("Final blog content", "final", final)
	} else {
		al.logger.Error("Final blog not found in session state", "error", err)
	}
}

