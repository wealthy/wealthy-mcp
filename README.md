# Wealthy MCP(Model Context Protocol) Server

This repository contains official Wealthy mcp server to help users with trading platform features


## Prerequisites

- Wealthy trading account - https://www.wealthy.in/broking
- Golang(go) 1.23 or later - https://go.dev/doc/install
- MCP clients - Claude or Cursor


## Getting Started

- Clone the repository:
   ```bash
   go install github.com/wealthy/wealthy-mcp@latest
   ```
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

