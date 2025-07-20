package main

import (
	"currency-trader/internal/bot"
	"flag"
	"time"
)

func main() {
	sessionCookie := flag.String("sessionCookie", "", "IBKR session cookie value")
	flag.Parse()
	if *sessionCookie == "" {
		panic("sessionCookie flag is required. Please provide the IBKR session cookie value.")
	}

	pairs := []string{"GBP/USD"}
	// Set the candle close time to 10pm GMT which is 5pm EST
	location, _ := time.LoadLocation("GMT")
	now := time.Now().In(location)
	candleCloseTime := time.Date(now.Year(), now.Month(), now.Day(), 22, 0, 0, 0, location)

	tradingBot := bot.NewBot(pairs, 50, 200, candleCloseTime, *sessionCookie)
	tradingBot.Start()
}
