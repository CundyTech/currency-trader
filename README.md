# Currency Trading Bot

This project is a currency trading bot that automates trading strategies on various exchanges. It is designed to be modular and extensible, allowing for easy integration of new trading strategies and exchange APIs. The bot currently supports multi-currency forex pair trading with strategies 50/200 SMA crossover and standard SMA.

## Project Structure

```
currency-trader
├── cmd
│   └── main.go          # Entry point of the application
├── internal
│   ├── bot
│   │   └── bot.go       # Trading bot implementation
│   ├── exchange
│   │   └── exchange.go  # Exchange interaction logic (IBKR HTTP API)
│   ├── strategy
│   │   └── strategy.go  # Trading strategies (SMA, SMA Crossover)
│   └── utils
│       └── utils.go     # Utility functions
├── go.mod               # Module definition
├── go.sum               # Dependency checksums
└── README.md            # Project documentation
```

## Features

- Multi-currency forex pair trading (e.g., GBP/USD, EUR/USD)
- Simple Moving Average (SMA) and 50/200 SMA crossover strategies
- Integration with Interactive Brokers Web API for real-time forex rates
- CI/CD with SonarQube/SonarCloud code quality and coverage

## Setup Instructions

1. **Clone the repository:**
   ```
   git clone https://github.com/CundyTech/currency-trader
   cd currency-trader
   ```

2. **Install dependencies:**
   ```
   go mod tidy
   ```

3. **Build the application (cross-platform):**
   - **For Windows:**
     ```
     go build -o currency-trader.exe ./cmd
     ```
   - **For Linux:**
     ```
     go build -o currency-trader ./cmd
     ```
   - **For macOS:**
     ```
     go build -o currency-trader ./cmd
     ```

4. **Run the application:**
   - On your development machine:
     ```
     go run cmd/main.go
     ```
   - Or on any target machine:
     ```
     ./currency-trader      # Linux/macOS
     currency-trader.exe    # Windows
     ```

## Usage

- Configure the bot to trade specific forex pairs by editing `cmd/main.go`.
- The bot supports both Simple Moving Average and SMA Crossover strategies. You can select or implement strategies in `internal/strategy/strategy.go`.
- To fetch real-time forex rates, configure your Interactive Brokers Web API credentials and session cookie in `internal/exchange/exchange.go`.
- The bot is scheduled to fetch rates and make trading decisions once per day at 10pm GMT.

## Examples

- **Implement a new trading strategy:** Create a struct that satisfies the `Strategy` interface in `internal/strategy/strategy.go`.
- **Fetch GBP/USD rate from IBKR:** Use `GetForexRate(baseURL, sessionCookie, "GBP/USD")` in `internal/exchange/exchange.go`.

## CI/CD and Quality Gates

- The project uses GitHub Actions and SonarQube/SonarCloud for code quality and coverage.
- Quality gate thresholds (e.g., minimum code coverage) are enforced via SonarQube/SonarCloud settings.

## Contributing

Contributions are welcome! Please open an issue or submit a pull request for any enhancements or bug fixes.

## License

This project is licensed under the MIT License. You are free to use, modify, and distribute this software for any purpose. See the [LICENSE](LICENSE) file for details.
