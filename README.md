# Wealthy MCP(Model Context Protocol) Server

This repository contains official Wealthy mcp server to help users with trading platform features
# MCP (Mark3 Labs Control Plane) Tools

This repository contains MCP server that provides integration with multiple tools. It also provides a framework for building and integrating various tools that can be used with the MCP server.


## Prerequisites

- Wealthy trading account - https://www.wealthy.in/broking
- Golang(go) 1.23 or later - https://go.dev/doc/install
- MCP clients - Claude or Cursor
- Go 1.23 or later


## Getting Started

You can either install from source or download pre-built artifacts:

### Option 1: Install from Source
- Install using golang(go 1.23 or later):
   ```bash
   go install github.com/wealthy/wealthy-mcp@latest
   ```

### Option 2: Download Pre-built Artifacts
- Visit our [releases page](https://github.com/wealthy/wealthy-mcp/releases) to download the latest pre-built binary for your platform
- Extract the downloaded archive and place the executable in a directory that's in your system PATH

## MCP Cursor Configuration

To configure MCP (Multi-Cursor Protocol) servers, create a `.cursor/mcp.json` file in your home directory with the following structure:

```json
{
    "mcpServers": {
      "wealthy-mcp": {
        "command": "mcp_server"
      }
    }
}
``` 


-  Restart Claude/Cursor
- Wealthy login page will be opened, enter wealthy credentials and after successful login, return to Claude/Cursor
- We have setup wealthy mcp server now you are ready to do some smart trading 🎉

1. Clone the repository:
   ```bash
   git clone https://github.com/wealthy/wealthy-mcp.git
   cd mcp
   ```

2. Install dependencies:
   ```bash
   go mod download
   ```

3. Set up your environment variables:
   ```bash
   cp .env.example .env
   # Edit .env with your configuration
   ```

## Creating a New Tool

To create a new tool, follow these steps:

1. Create a new file in the `tools` directory for your tool (e.g., `tools/my_tool.go`):
   ```go
   package tools

   import (
       "context"
       "github.com/mark3labs/mcp-go/server"
       "github.com/wealthy/wealthy-mcp"
   )

   // Define your tool's parameters
   type MyToolParams struct {
       // Add your parameters with jsonschema tags for documentation
       Param1 string `json:"param1" jsonschema:"description=Description of param1"`
       Param2 int    `json:"param2" jsonschema:"description=Description of param2"`
   }

   // Implement your tool's handler function
   func handleMyTool(ctx context.Context, args MyToolParams) (any, error) {
       // Implement your tool's logic here
       return nil, nil
   }

   // Define your tool using MustTool
   var MyTool = mcp.MustTool(
       "my_tool",                    // Tool name
       "Description of my tool",     // Tool description
       handleMyTool,                 // Handler function
   )

   // Create a registration function
   func AddMyTool(mcp *server.MCPServer) {
       MyTool.Register(mcp)
   }
   ```

2. Register your tool in your main application (typically in main.go):
   ```go
   func main() {
       server := mcp.NewServer()
       tools.AddMyTool(server)  // Register your new tool
       // ... rest of your server setup
   }
   ```

## Tool Definition Best Practices

1. **Parameter Documentation**: Always provide clear descriptions for your tool's parameters using the `jsonschema` tag.

2. **Error Handling**: Return meaningful errors that can help users understand what went wrong.

3. **Type Safety**: Use strongly typed parameter structures.

4. **Context Usage**: Always accept and use the context parameter for cancellation and timeout support.

5. **Tool Registration**: Follow the naming convention `Add<ToolName>` for your registration function.

6. **Package Structure**: Place all tool-related code in the `tools` package.

## Example: Plane.io Integration

The repository includes an example integration with Plane.io in the `plane` directory. This demonstrates how to:

- Define tool models with proper JSON schema tags
- Implement tool handlers with proper error handling
- Structure your code for maintainability
- Write tests for your tools

## Testing

Write tests for your tools using Go's testing package:

```go
func TestMyTool(t *testing.T) {
    ctx := context.Background()
    req := MyToolRequest{
        Param1: "test",
        Param2: 42,
    }
    
    resp, err := MyToolHandler(ctx, req)
    if err != nil {
        t.Fatalf("unexpected error: %v", err)
    }
    
    // Add your assertions here
}
```

## Contributing

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

### Trading Disclaimer

The software is provided for informational purposes only and should not be construed as investment advice. Trading in financial instruments involves high risks including the risk of losing some, or all, of your investment amount. Before trading, you should carefully consider your investment objectives, level of experience, and risk appetite.

### Legal Notice

By using this software, you acknowledge that:
1. You understand the risks associated with trading
2. You will seek professional financial advice when needed
3. The authors and contributors are not liable for any trading losses
4. You will comply with all applicable financial regulations

## Support

For support, please open an issue in the GitHub repository or contact the maintainers. 

