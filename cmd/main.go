package main

import (
	"currency-trader/internal/bot"
	"time"
)

func main() {
	pairs := []string{"GBP/USD"}
	// Set the candle close time to 10pm GMT which is 5pm EST
	location, _ := time.LoadLocation("GMT")
	now := time.Now().In(location)
	candleCloseTime := time.Date(now.Year(), now.Month(), now.Day(), 22, 0, 0, 0, location)

	tradingBot := bot.NewBot(pairs, 50, 200, candleCloseTime)
	tradingBot.Start()
}
