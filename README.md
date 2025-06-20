# RabbitMQ MCP Go

A Go implementation of Model Control Protocol (MCP) server for RabbitMQ integration.

## Project Structure

```
.
├── cmd/                    # Application entry points
├── internal/              # Private application code
│   ├── config/           # Configuration management
│   ├── tools/            # MCP tools implementations
│   ├── prompts/          # Prompt templates and configurations
│   └── resources/        # Shared resources and utilities
├── pkg/                   # Public library code
├── api/                   # API definitions and documentation
└── scripts/              # Build and utility scripts
```

## Directory Overview

### Internal Directory Structure

- **tools/**: Contains all MCP tool implementations. Each tool is a separate package that implements the MCP tool interface.
- **prompts/**: Stores prompt templates and configurations used by the MCP tools.
- **resources/**: Houses shared utilities, helper functions, and reusable components.
- **config/**: Manages application configuration.

## Requirements

- Go 1.24.3 or higher
- RabbitMQ server
- Cursor IDE or Claude Desktop for AI integration

## Getting Started

1. Install Go 1.24.3:
```bash
asdf install golang 1.24.3
```

2. Clone the repository:
```bash
git clone https://github.com/yourusername/rabbitmq-mcp-go.git
cd rabbitmq-mcp-go
```

3. Install dependencies:
```bash
go mod download
```

4. Build the project:
```bash
go build ./cmd/...
```

## MCP Client Integration

### Basic Configuration

Add this configuration to your MCP client settings:

```json
{
  "mcpServers": {
    "rabbitmq": {
      "command": "rabbitmq-mcp-server",
      "env": {
        "RABBITMQ_URL": "amqp://guest:guest@localhost:5672/",
        "MCP_TRANSPORT": "stdio"
      }
    }
  }
}
```

For secured environments:

```json
{
  "mcpServers": {
    "rabbitmq": {
      "command": "rabbitmq-mcp-server",
      "env": {
        "RABBITMQ_URL": "amqps://user:password@rabbitmq-host:5671/",
        "MCP_TRANSPORT": "stdio",
        "RABBITMQ_TLS_ENABLE": "true",
        "RABBITMQ_TLS_VERIFY": "true"
      }
    }
  }
}
```

### Cursor Integration

To use with [Cursor](https://cursor.sh/), create or edit `~/.cursor/mcp.json`:

```json
{
  "mcpServers": {
    "rabbitmq": {
      "command": "rabbitmq-mcp-server",
      "env": {
        "RABBITMQ_URL": "amqp://guest:guest@localhost:5672/",
        "MCP_TRANSPORT": "stdio"
      }
    }
  }
}
```

### Claude Desktop Integration

To use with Claude Desktop, edit your configuration file:
- macOS: `~/Library/Application Support/Claude/claude_desktop_config.json`
- Windows: `%APPDATA%\Claude\claude_desktop_config.json`

```json
{
  "mcpServers": {
    "rabbitmq": {
      "command": "rabbitmq-mcp-server",
      "env": {
        "RABBITMQ_URL": "amqp://guest:guest@localhost:5672/",
        "MCP_TRANSPORT": "stdio"
      }
    }
  }
}
```

## Using the RabbitMQ Publish Tool

The publish tool allows you to send messages to RabbitMQ queues or exchanges through the MCP interface.

### Tool Parameters

- `queue` (string, optional): Queue name to publish to
- `exchange` (string, optional): Exchange name to publish to
- `message` (string, required): Message content to publish
- `content_type` (string, optional): Content type of the message
  - Supported values: "text/plain" (default), "application/json"
- `headers` (string, optional): Message headers in JSON format

### Example Interactions

1. Publishing to a Queue:
```
User: "Send a message 'Hello World' to the queue 'my_queue'"

AI Assistant will use the publish tool:
{
  "queue": "my_queue",
  "message": "Hello, World!",
  "content_type": "text/plain"
}
```

2. Publishing JSON to an Exchange:
```
User: "Publish order status update to the 'orders' exchange"

AI Assistant will use the publish tool:
{
  "exchange": "orders",
  "message": "{\"order_id\": \"12345\", \"status\": \"completed\"}",
  "content_type": "application/json"
}
```

3. Publishing with Headers:
```
User: "Send a high-priority message to the notifications queue"

AI Assistant will use the publish tool:
{
  "queue": "notifications",
  "message": "Important system update",
  "headers": "{\"priority\": \"high\", \"timestamp\": \"2024-03-20T12:00:00Z\"}"
}
```

### Error Handling

The tool will return an error in the following cases:
- Neither queue nor exchange is specified
- Required message parameter is missing
- Invalid JSON format when content_type is "application/json"
- RabbitMQ connection or publishing errors

## Development

To add a new MCP tool:

1. Create a new directory under `internal/tools/`
2. Implement the MCP tool interface
3. Register the tool in the main application
4. Add corresponding prompts in `internal/prompts/` if needed
5. Document the tool in the tools README
