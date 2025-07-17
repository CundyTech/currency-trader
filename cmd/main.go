package main

import (
	"currency-trader/internal/bot"
)

func main() {
	pairs := []string{"GBP/USD"}
	tradingBot := bot.NewBot(pairs, 50, 200)
	tradingBot.Start()
}
