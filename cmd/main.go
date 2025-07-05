package main

import (
	"currency-trader/internal/bot"
	"currency-trader/internal/strategy"
	"fmt"
)

func main() {
	smaStrategy := &strategy.SimpleMovingAverage{Period: 10}
	tradingBot := bot.NewBot(smaStrategy)
	fmt.Println("Starting the trading bot...")
	tradingBot.Start()
}
