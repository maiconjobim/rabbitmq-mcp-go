# Tools

This directory contains all the MCP (Model Control Protocol) tools implemented in the project. Each tool is a separate package that implements the MCP tool interface, providing specific functionality to interact with RabbitMQ.

## Structure

Each tool follows these conventions:
- Located in its own subdirectory for clean organization
- Implements the MCP tool interface defined by `mcp.Tool`
- Contains both the tool definition and its handler implementation
- Includes appropriate tests in a `_test.go` file
- Uses consistent error handling and logging patterns

## Implementing a New Tool

To create a new tool:

1. Create a new directory under `internal/tools/`
2. Implement the tool using `mcp.NewTool()` with appropriate parameters
3. Define the tool handler function that processes the request
4. Register the tool in the main application
5. Add tests for your implementation

Example structure:
```
internal/tools/
├── mytool/
│   ├── mytool.go      # Main implementation
│   └── mytool_test.go # Tests
```

## Current Tools

- `publish`: Publishes messages to RabbitMQ queues or exchanges
  - Supports both direct queue and exchange publishing
  - Handles different content types (text/plain, application/json)
  - Provides message headers support

## Best Practices

1. **Input Validation**
   - Use MCP's built-in validation (Required, Enum, Pattern)
   - Add runtime validation for complex cases
   - Return clear error messages

2. **Error Handling**
   - Use `mcp.NewToolResultError` for user-facing errors
   - Use `mcp.NewToolResultErrorFromErr` for system errors
   - Include context in error messages

3. **Documentation**
   - Add clear descriptions to all tool parameters
   - Document any special requirements or behaviors
   - Include examples where helpful

## Testing

Each tool should include tests that cover:
- Happy path scenarios
- Error cases
- Edge cases
- Input validation
- Integration with RabbitMQ (where applicable) 