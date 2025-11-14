// Package blog provides a sequential agent pipeline for blog post creation
//
// AGENT PATTERN: Sequential Agent (Deterministic Pipeline)
// - Agents execute in guaranteed order: Outline → Writer → Editor
// - Each agent passes output to next via session state
// - Predictable, linear workflow with no branching
package blog

import (
	"fmt"

	"google.golang.org/adk/agent"
	"google.golang.org/adk/agent/llmagent"
	"google.golang.org/adk/agent/workflowagents/sequentialagent"
	"google.golang.org/adk/model"
)

const (
	// OutputKeyBlogOutline is the session state key for blog outline
	OutputKeyBlogOutline = "blog_outline"
	// OutputKeyBlogDraft is the session state key for blog draft
	OutputKeyBlogDraft = "blog_draft"
	// OutputKeyFinalBlog is the session state key for final blog
	OutputKeyFinalBlog = "final_blog"
)

// Build creates a sequential agent pipeline (Outline -> Writer -> Editor)
func Build(mdl model.LLM) (agent.Agent, error) {
	// Outline Agent: Creates blog outline
	outlineAgent, err := llmagent.New(llmagent.Config{
		Name:        "OutlineAgent",
		Model:       mdl,
		Description: "Creates blog post outlines",
		Instruction: `Create a blog outline for the given topic with:
1. A catchy headline
2. An introduction hook
3. 2-3 main sections with key points
4. A concluding thought`,
		OutputKey: OutputKeyBlogOutline,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to create outline agent: %w", err)
	}

	// Writer Agent: Writes the full blog post
	writerAgent, err := llmagent.New(llmagent.Config{
		Name:        "WriterAgent",
		Model:       mdl,
		Description: "Writes blog posts from outlines",
		Instruction: "Following the outline from the session state under 'blog_outline', write a short blog post (200-300 words) with an engaging and informative tone.",
		OutputKey:   OutputKeyBlogDraft,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to create writer agent: %w", err)
	}

	// Editor Agent: Edits and polishes the draft
	editorAgent, err := llmagent.New(llmagent.Config{
		Name:        "EditorAgent",
		Model:       mdl,
		Description: "Edits and polishes blog post drafts",
		Instruction: "Edit the draft from the session state under 'blog_draft'. Polish the text by fixing grammatical errors, improving flow and sentence structure, and enhancing overall clarity.",
		OutputKey:   OutputKeyFinalBlog,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to create editor agent: %w", err)
	}

	// Sequential Agent: Runs agents in guaranteed order
	sequentialPipeline, err := sequentialagent.New(sequentialagent.Config{
		AgentConfig: agent.Config{
			Name:      "BlogPipeline",
			SubAgents: []agent.Agent{outlineAgent, writerAgent, editorAgent},
		},
	})
	if err != nil {
		return nil, fmt.Errorf("failed to create sequential agent: %w", err)
	}

	return sequentialPipeline, nil
}

// DefaultPrompt returns a sample prompt for this agent system
func DefaultPrompt() string {
	return "Write a blog post about the benefits of multi-agent systems for software developers"
}

