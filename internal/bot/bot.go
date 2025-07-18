package bot

import (
	exchange "currency-trader/internal/exchange"
	"currency-trader/internal/strategy"
	w "currency-trader/internal/wallet"
	"fmt"
	"strconv"
	"time"
)

type PriceFeed struct {
	Pair  string
	Price float64
}

type Bot struct {
	isTrading       bool
	strategies      map[string]*strategy.SMACrossover
	wallet          w.Wallet
	pairs           []string
	candleCloseTime time.Time // The time of day when the strategy should execute
}

func NewBot(pairs []string, shortPeriod int, longPeriod int, candleCloseTime time.Time) *Bot {
	strategies := make(map[string]*strategy.SMACrossover)
	for _, pair := range pairs {
		strategies[pair] = &strategy.SMACrossover{ShortPeriod: shortPeriod}
	}
	return &Bot{
		pairs:           pairs,
		isTrading:       false,
		strategies:      strategies,
		wallet:          w.Wallet{Transactions: []w.Transaction{}, Balance: 10.0},
		candleCloseTime: candleCloseTime,
	}
}

func (b *Bot) Start() {
	fmt.Println("Starting the trading bot...")
	b.isTrading = true
	for b.isTrading {
		duration := waitUntil(b.candleCloseTime)
		fmt.Printf("Sleeping for %v until next candle close time", duration)
		time.Sleep(duration)
		data, err := exchange.GetForexRate("https://localhost:5000", "x-sess-uuid=0.981e1202.1752786317.178fe338", b.pairs[0])
		if err != nil {
			fmt.Println("Error fetching forex rate:", err)
			return
		}

		strat := b.strategies[b.pairs[0]]

		//convert the price from string to float64
		price, err := strconv.ParseFloat(data.Price, 64)
		if err != nil {
			fmt.Println("Error converting price to float64:", err)
			return
		}
		signal := strat.GenerateSignal(price)
		if b.wallet.Balance <= 0 {
			println("Insufficient balance to continue trading.")
			b.Stop()
			return
		}
		if signal == "buy" {
			b.wallet.Transactions = w.AddTransaction(b.wallet.Transactions, "withdrawl", price)
		}
		if signal == "sell" {
			b.wallet.Transactions = w.AddTransaction(b.wallet.Transactions, "deposit", price)
		}
		println("Action:", signal)
		if signal == "sell" || signal == "buy" {
			println("Current Balance:", fmt.Sprintf("%.2f", b.wallet.CalculateBalance(b.wallet.Transactions)))
		}

	}
}

// waitUntil calculates the duration until the next 10 PM GMT
func waitUntil(next time.Time) time.Duration {
	now := time.Now().UTC()
	if now.After(next) {
		next = next.Add(24 * time.Hour)
	}
	return next.Sub(now)
}

func (b *Bot) Stop() {
	b.isTrading = false
}
