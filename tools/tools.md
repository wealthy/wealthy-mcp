# Wealthy MCP Tools Documentation

This document provides an overview of the various tools available in the Wealthy MCP system.

## Watchlist Tools

### Create Watchlist (`create_watchlist`)
A tool for creating and managing watchlists of securities.

**Functionality:**
- Add securities to a watchlist
- Track multiple watchlists
- Manage watchlist entries

## Research Tool

### Research (`research`)
A tool for accessing trading ideas and research information.

**Functionality:**
- Get trade ideas
- Access research recommendations
- View market analysis

## Reports Tool

### Reports (`reports_tool`)
A comprehensive tool for generating various types of reports.

**Supported Report Types:**
- Holdings Report (`holdings`)
- Positions Report (`positions`)
- Order Book Report (`order_book`)

**Parameters:**
- `report`: Type of report to generate (holdings/positions/order_book)

## Orders Tool

### Place Order (`place_order`)
A tool for executing trading orders.

**Parameters:**
- `exchange_name`: Exchange identifier (1=NSE, 2=NFO, 3=BSE, 4=BFO)
- `token`: Trading token
- `trading_symbol`: Symbol to trade
- `quantity`: Quantity to trade
- `price`: Order price
- `trigger_price`: Trigger price for stop orders
- `order_type`: Type of order (1=CNC, 2=MIS, 3=NRML, 4=COVER, 5=BRACKET, 6=MTF)
- `transaction_type`: Buy (1) or Sell (2)
- `price_type`: Price type (1=LMT, 2=MKT, 3=SLLMT, 4=SLMKT, 5=DS, 6=TWOLEG, 7=THREEELEG)
- `validity`: Order validity (1=DAY, 2=IOC, 3=EOS, 4=GTT)
- `disc_quantity`: Disclosed quantity
- `is_amo`: After Market Order flag

**Protection Parameters:**
- `target_price`: Target price for the order
- `stop_loss_price`: Stop loss price
- `trail_price`: Trailing price

## Search Tool

### Search (`search`)
A tool for searching and finding security symbols.

**Parameters:**
- `query`: Search query for finding a security symbol

## Price Tool

### Get Price (`get_price`)
A tool for retrieving real-time price information for securities.

**Parameters:**
- `symbols`: Array of symbols to get prices for
  - Format: `exchange:trading_symbol`
  - Examples: 
    - `nse:RELIANCE-EQ`
    - `bse:RELIANCE-EQ`
    - `nse:INFY-EQ`
    - `bse:INFY`
    - `nfo:RELIANCE29MAY25F`

## Best Practices

1. Always validate input parameters before making requests
2. Handle errors appropriately
3. Use proper error handling mechanisms
4. Follow rate limiting guidelines
5. Implement proper timeout handling
6. Use logging for debugging and monitoring

## Error Handling

Each tool implements proper error handling and validation:
- Input parameter validation
- Business logic validation
- Error response formatting
- Proper error propagation

## Security Considerations

1. Handle API keys and credentials securely
2. Use environment variables for sensitive information
3. Implement proper authentication and authorization
4. Monitor and implement rate limiting
5. Follow security best practices

## Contributing

When contributing to these tools:
1. Follow the existing code structure
2. Add comprehensive tests
3. Update documentation
4. Follow security best practices
5. Add proper error handling
6. Include logging statements
