# Wealthy MCP(Model Context Protocol) Server

This repository contains official Wealthy mcp server to help users with trading platform features

## Prerequisites

- Wealthy trading account - https://www.wealthy.in/broking
- MCP clients - Claude or Cursor

## Getting Started

You can either install from source or download pre-built artifacts:

### Option 1: Install from Source
- Install using golang(go 1.23 or later):
   ```bash
   go install github.com/wealthy/wealthy-mcp/cmd/wealthy-mcp@latest
   ```

### Option 2: Download Pre-built Artifacts
- Visit our [releases page](https://github.com/wealthy/wealthy-mcp/releases) to download the latest pre-built binary for your platform
- Extract the downloaded archive and place the executable in a directory that's in your system PATH
- Run below commands to give permissions on mac:
```bash
chmod +x wealthy-mcp-macos-arm64-<version>
xattr -d com.apple.quarantine wealthy-mcp-macos-arm64-<version>
```

## MCP Configuration

To configure MCP (Multi-Cursor Protocol) servers, create/update the appropriate configuration file based on your client:

### For Cursor
Create/update `.cursor/mcp.json` in your home directory:
```json
{
    "mcpServers": {
      "wealthy-mcp": {
        "command": "<path to downloaded/installed binary>"
      }
    }
}
```

### For Claude
Create/update `claude_desktop_config.json`:
```json
{
    "mcpServers": {
      "wealthy-mcp": {
        "command": "<path to downloaded/installed binary>"
      }
    }
}
```

To use a custom port, add the `args` parameter:
```json
{
    "mcpServers": {
      "wealthy-mcp": {
        "command": "<path to downloaded/installed binary>",
        "args": ["addr=localhost:8006"]
      }
    }
}
```

- Restart your MCP client (Claude/Cursor)
- Wealthy login page will open - enter your credentials
- After successful login, return to your MCP client
- You're now ready to start smart trading! 🎉

## Usage

Here are the available query types and their purposes:

| Query Type | Purpose |
|------------|---------|
| `get_user_margin` | Retrieves user margin information |
| `get_price` | Retrieves the current market price for a specified trading symbol |
| `get_holdings` | Shows your current portfolio holdings and their details |
| `get_positions` | Displays your open trading positions |
| `get_order_book` | Lists all your orders (open, executed, and cancelled) |
| `get_trade_ideas` | Provides trading suggestions and market insights |
| `get_security_info` | Fetches detailed information about a specific security/stock |
| `place_order` | Places a new buy/sell order with specified parameters |
| `create_watchlist` | Creates a new watchlist of securities |
| `get_watchlist` | Retrieves existing watchlists |
| `update_watchlist` | Updates an existing watchlist with new securities |
| `search` | Searches and finds security symbols |
| `research` | Accesses trading ideas and research information |
| `reports_tool` | Generates various types of reports (holdings/positions/order_book) |

You can interact with these queries through natural language in Claude/Cursor. For example:
- "What is the price of RELIANCE?"
- "Show me my current holdings"
- "I want to buy 100 shares of TATAMOTORS at market price"
- "Create a watchlist for tech stocks"
- "Search for stocks related to renewable energy"
- "Get research recommendations for banking sector"

## Roadmap

- Real-time Market Data Integration
  - Implement WebSocket connections for live price streaming
  - Add support for multiple symbols subscription
  - Enhance price data formatting and display

- Price Alerts System
  - Enable setting price alerts through natural language
  - Support multiple alert conditions (above/below/crossing)
  - Implement alert notifications within chat
  - Add alert management (list/modify/delete alerts)

- Advanced Features
  - Portfolio analytics with real-time updates
  - Custom watchlists with real-time updates
  - Advanced alert conditions (technical indicators)
  - Enhanced reporting and analytics

- Platform Enhancements
  - Advanced order types
  - Automated trading strategies
  - Support for additional financial products

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

For support, please open an issue in the GitHub repository, contact the maintainers, or email Parth at parath.singh@wealthy.in.

