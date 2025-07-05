package bot

import (
	s "currency-trader/internal/strategy"
	w "currency-trader/internal/wallet"
	"fmt"
)

type Bot struct {
	isTrading bool
	strategy  s.Strategy
	wallet    w.Wallet
}

func NewBot(s s.Strategy) *Bot {
	return &Bot{
		isTrading: false,
		strategy:  s,
		wallet:    w.Wallet{Transactions: []w.Transaction{}, Balance: 10.0},
	}
}

func (b *Bot) Start() {
	b.isTrading = true
	for b.isTrading {
		// Example: Simulate price feed
		prices := []float64{1.10, 1.12, 1.11, 1.15, 1.13, 1.18, 1.16, 1.20, 1.22, 1.19, 1.23}
		for _, price := range prices {
			action, _ := b.strategy.Execute(price)
			if b.wallet.Balance <= 0 {
				println("Insufficient balance to continue trading.")
				b.Stop()
				return
			}
			if action == "buy" {
				b.wallet.Transactions = w.AddTransaction(b.wallet.Transactions, "withdrawl", price)
			}
			if action == "sell" {
				b.wallet.Transactions = w.AddTransaction(b.wallet.Transactions, "deposit", price)
			}
			println("Action:", action)
			if action == "sell" || action == "buy" {
				println("Current Balance:", fmt.Sprintf("%.2f", b.wallet.CalculateBalance(b.wallet.Transactions)))
			}
		}
	}
}

func (b *Bot) Stop() {
	b.isTrading = false
	// Logic to stop trading
}
