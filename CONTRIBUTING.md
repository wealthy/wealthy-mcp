# Contributing to Wealthy MCP Server

Thank you for your interest in contributing to the Wealthy MCP Server! This document provides guidelines and instructions for contributing to this project.

## Table of Contents

- [Code of Conduct](#code-of-conduct)
- [Getting Started](#getting-started)
- [Development Environment](#development-environment)
- [Contribution Workflow](#contribution-workflow)
- [Pull Request Guidelines](#pull-request-guidelines)
- [Coding Standards](#coding-standards)
- [Testing](#testing)
- [Documentation](#documentation)
- [Release Process](#release-process)

## Code of Conduct

Please be respectful and considerate of others when contributing to this project. We aim to foster an inclusive and welcoming community.

## Getting Started

1. Fork the repository on GitHub
2. Clone your fork locally:
   ```bash
   git clone https://github.com/YOUR-USERNAME/wealthy-mcp.git
   cd wealthy-mcp
   ```
3. Add the upstream repository as a remote:
   ```bash
   git remote add upstream https://github.com/wealthy/wealthy-mcp.git
   ```
4. Create a new branch for your changes:
   ```bash
   git checkout -b feature/your-feature-name
   ```

## Development Environment

### Prerequisites

- Go 1.23 or later
- A Wealthy trading account for testing
- Claude or Cursor for MCP client testing

### Setup

1. Install dependencies:
   ```bash
   go mod download
   ```
2. Build the project:
   ```bash
   go build ./cmd/*.go
   ```

## Contribution Workflow

1. Make sure you have the latest changes from the upstream repository:
   ```bash
   git fetch upstream
   git rebase upstream/main
   ```
2. Implement your changes
3. Write or update tests as needed
4. Ensure all tests pass:
   ```bash
   go test ./...
   ```
5. Commit your changes with a descriptive message:
   ```bash
   git commit -m "Add feature: description of your changes"
   ```
6. Push to your fork:
   ```bash
   git push origin feature/your-feature-name
   ```
7. Create a pull request from your fork to the main repository

## Pull Request Guidelines

- Fill in the required pull request template
- Include a clear description of the changes and their purpose
- Reference any related issues using the GitHub issue linking syntax (e.g., "Fixes #123")
- Ensure all CI checks pass
- Update documentation as needed
- Keep pull requests focused on a single concern

## Coding Standards

- Follow standard Go coding conventions
- Use meaningful variable and function names
- Write clear comments for complex logic
- Format your code using `gofmt` or `go fmt`
- Use `golint` and `go vet` to ensure code quality

## Testing

- Write tests for new features and bug fixes
- Ensure existing tests pass with your changes
- Focus on both unit tests and integration tests
- Use the `testing` package for writing tests
- Run tests locally before submitting a pull request:
  ```bash
  go test ./...
  ```

## Documentation

- Update README.md if your changes affect usage
- Document new features or API changes
- Update code documentation (comments) as needed
- Consider adding examples for complex features

## Release Process

The project follows semantic versioning. Releases are created automatically through GitHub Actions when a new tag is pushed:

1. Tags should follow the format `vX.Y.Z` (e.g., `v1.0.0`)
2. The GitHub workflow will build binaries for Linux, macOS, and Windows
3. Releases are published to the GitHub releases page

## Questions?

If you have any questions about contributing, please open an issue in the repository and we'll be happy to help!

Thank you for contributing to the Wealthy MCP Server project!