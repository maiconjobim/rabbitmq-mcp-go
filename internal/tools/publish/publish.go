package publish

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/maiconjobim/rabbitmq-mcp-go/internal/config"
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/rabbitmq/amqp091-go"
)

func NewPublishTool() mcp.Tool {
	return mcp.NewTool("rabbitmq_publish",
		mcp.WithDescription("Publish a message to a RabbitMQ queue or exchange"),
		mcp.WithString("queue",
			mcp.Description("Queue name to publish to (optional if exchange is set)"),
		),
		mcp.WithString("exchange",
			mcp.Description("Exchange name to publish to (optional if queue is set)"),
		),
		mcp.WithString("message",
			mcp.Required(),
			mcp.Description("Message content to publish"),
		),
		mcp.WithString("headers",
			mcp.Description("Optional headers in JSON format"),
		),
		mcp.WithString("content_type",
			mcp.Description("Content-Type of the message (e.g., text/plain, application/json)"),
			mcp.Enum("text/plain", "application/json"),
		),
	)
}

func PublishHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	cfg := config.Load()

	queue := request.GetString("queue", "")
	exchange := request.GetString("exchange", "")
	message, err := request.RequireString("message")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}
	headers := request.GetString("headers", "")
	contentType := request.GetString("content_type", "text/plain")

	if contentType == "application/json" {
		// Validate JSON
		var js interface{}
		if err := json.Unmarshal([]byte(message), &js); err != nil {
			return mcp.NewToolResultError("message is not valid JSON"), nil
		}
	}

	if queue == "" && exchange == "" {
		return mcp.NewToolResultError("either queue or exchange must be provided"), nil
	}

	conn, err := amqp091.Dial(cfg.GetRabbitMQURL())
	if err != nil {
		return mcp.NewToolResultError(fmt.Sprintf("failed to connect to RabbitMQ: %v", err)), nil
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		return mcp.NewToolResultError(fmt.Sprintf("failed to open channel: %v", err)), nil
	}
	defer ch.Close()

	var amqpHeaders amqp091.Table
	if headers != "" {
		amqpHeaders = amqp091.Table{}
	}

	var routingKey string
	if queue != "" {
		routingKey = queue
	} else {
		routingKey = ""
	}

	err = ch.PublishWithContext(ctx,
		exchange,
		routingKey,
		false,
		false,
		amqp091.Publishing{
			ContentType: contentType,
			Body:        []byte(message),
			Headers:     amqpHeaders,
		})

	if err != nil {
		return mcp.NewToolResultError(fmt.Sprintf("failed to publish message: %v", err)), nil
	}

	destination := queue
	if destination == "" {
		destination = exchange
	}

	return mcp.NewToolResultText(fmt.Sprintf("Message published successfully to %s", destination)), nil
}
