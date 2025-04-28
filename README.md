# Wealthy MCP(Model Context Protocol) Server

This repository contains official Wealthy mcp server to help users with trading platform features


## Prerequisites

- Wealthy trading account - https://www.wealthy.in/broking
- Golang(go) 1.23 or later - https://go.dev/doc/install
- MCP clients - Claude or Cursor


## Getting Started

You can either install from source or download pre-built artifacts:

### Option 1: Install from Source
- Clone the repository:
   ```bash
   go install github.com/wealthy/wealthy-mcp@latest
   ```

### Option 2: Download Pre-built Artifacts
- Visit our [releases page](https://github.com/wealthy/wealthy-mcp/releases) to download the latest pre-built binary for your platform
- Extract the downloaded archive and place the executable in a directory that's in your system PATH

### Configure MCP Client
-  open Claude/Cursor mcp config
    add below json 
     ```
    {
     "mcpServers": {
      "falcon": {
                "command": "mcp_server"
            }
        }
    }
     ```
-  Restart Claude/Cursor
- Wealthy login page will be opened, enter wealthy credentials and after successful login, return to Claude/Cursor
- We have setup wealthy mcp server now you are ready to do some smart trading 🎉

