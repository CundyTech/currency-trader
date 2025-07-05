# Currency Trading Bot

This project is a currency trading bot that automates trading strategies on various exchanges. It is designed to be modular and extensible, allowing for easy integration of new trading strategies and exchange APIs.

## Project Structure

```
currency-trader
├── cmd
│   └── main.go          # Entry point of the application
├── internal
│   ├── bot
│   │   └── bot.go      # Trading bot implementation
│   ├── exchange
│   │   └── exchange.go  # Exchange interaction logic
│   ├── strategy
│   │   └── strategy.go  # Trading strategies definitions
│   └── utils
│       └── utils.go     # Utility functions
├── go.mod               # Module definition
├── go.sum               # Dependency checksums
└── README.md            # Project documentation
```

## Setup Instructions

1. **Clone the repository:**
   ```
   git clone <repository-url>
   cd currency-trader
   ```

2. **Install dependencies:**
   ```
   go mod tidy
   ```

3. **Run the application:**
   ```
   go run cmd/main.go
   ```

## Usage

- The bot can be configured to use different trading strategies by implementing the `Strategy` interface in the `internal/strategy/strategy.go` file.
- To start trading, call the `Start()` method on the `Bot` struct defined in `internal/bot/bot.go`.

## Examples

- Implement a new trading strategy by creating a new struct that satisfies the `Strategy` interface.
- Use the `Exchange` struct in `internal/exchange/exchange.go` to interact with your preferred trading exchange.

## Contributing

Contributions are welcome! Please open an issue or submit a pull request for any enhancements or bug fixes.