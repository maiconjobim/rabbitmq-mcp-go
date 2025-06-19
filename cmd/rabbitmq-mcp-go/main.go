package main

import (
	"fmt"

	"github.com/mark3labs/mcp-go/server"

	"github.com/maiconjobim/rabbitmq-mcp-go/internal/config"
	"github.com/maiconjobim/rabbitmq-mcp-go/internal/tools/publish"
)

func main() {
	cfg := config.Load()

	s := server.NewMCPServer(
		cfg.MCPServer.Name,
		cfg.MCPServer.Version,
		server.WithToolCapabilities(true),
		server.WithInstructions("A server that does amazing things"),
	)

	s.AddTool(publish.NewPublishTool(), publish.PublishHandler)

	if err := server.ServeStdio(s); err != nil {
		fmt.Printf("Server error: %v\n", err)
	}
}
