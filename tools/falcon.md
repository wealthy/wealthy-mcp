# Falcon Trading Platform Integration

This document describes the integration with the Falcon trading platform through the MCP server.

## Overview

The Falcon tool provides a seamless interface to interact with the Falcon trading platform, allowing you to:
- Place trading orders
- Get holdings information
- Get positions
- Get security information
- Get order book details
- Get trade ideas
- Get real-time price information

## Query Types

### 1. Place Order (`place_order`)
Places a new trading order on the Falcon platform.

**Parameters:**
- `account_id`: Trading account identifier
- `trading_symbol`: Symbol to trade
- `transaction_type`: Buy (1) or Sell (2)
- `quantity`: Number of units to trade
- `price`: Order price
- `trigger_price`: Price at which to trigger the order (for stop orders)
- `stop_loss_price`: Stop loss price
- `target_price`: Target price
- `trailing_price`: Trailing stop price
- `disclosed_quantity`: Quantity to disclose
- `validity`: Order validity period
- `is_amo`: After Market Order flag
- `order_type`: Type of order
- `price_type`: Price type specification

### 2. Get Holdings (`get_holdings`)
Retrieves current holdings information.

**Parameters:**
- `account_id`: Trading account identifier

### 3. Get Positions (`get_positions`)
Retrieves current positions information.

**Parameters:**
- `account_id`: Trading account identifier

### 4. Get Security Info (`get_security_info`)
Retrieves detailed information about a security.

**Parameters:**
- `trading_symbol`: Symbol to query
- `exchange_name`: Exchange identifier

### 5. Get Order Book (`get_order_book`)
Retrieves the current order book.

**Parameters:** None required

### 6. Get Trade Ideas (`get_trade_ideas`)
Retrieves trading ideas and recommendations.

**Parameters:** None required

### 7. Get Price (`get_price`)
Retrieves real-time price information.

**Parameters:**
- `trading_symbol`: Symbol to query price for

## Error Handling

The tool returns detailed error messages in the following cases:
- Invalid query type
- Missing required parameters
- Network connectivity issues
- Platform-specific errors

## Security Considerations

1. Always handle API keys and credentials securely
2. Use environment variables for sensitive information
3. Implement proper authentication and authorization
4. Monitor and implement rate limiting as per platform guidelines

## Example Usage

```go
args := falcon.FalconRequest{
    QueryType: "place_order",
    AccountID: "ACC123",
    TradingSymbol: "AAPL",
    TransactionType: 1, // Buy
    Quantity: 100,
    Price: "150.00",
}

result, err := falconService.PlaceOrder(ctx, args)
if err != nil {
    // Handle error
}
```

## Best Practices

1. Always use proper error handling
2. Implement retry mechanisms for transient failures
3. Use logging for debugging and monitoring
4. Validate input parameters before making requests
5. Follow rate limiting guidelines
6. Implement proper timeout handling

## Rate Limiting

The Falcon tool implements rate limiting to prevent overloading the trading platform. Default timeout is set to 60 seconds.

## Contributing

When contributing to the Falcon tool:
1. Follow the existing code structure
2. Add comprehensive tests
3. Update documentation
4. Follow security best practices
5. Add proper error handling
6. Include logging statements

## Testing

Run tests using:
```bash
go test ./tools -v
```

## Support

For issues and support:
1. Check the existing issues in the repository
2. Create a new issue with detailed information
3. Include relevant logs and error messages
4. Provide steps to reproduce the issue 