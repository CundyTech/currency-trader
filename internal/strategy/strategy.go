package strategy

import (
	"currency-trader/internal/utils"
	"fmt"
)

type Strategy interface {
	Execute(price float64) (string, error) // Execute the trading strategy and return the action (buy/sell) and any error
}

// Simple Moving Average strategy that calculates the average price over a specified period
type SimpleMovingAverage struct {
	Period     int // The period for the moving average
	Prices     []float64
	LastSignal string // "buy", "sell", or "do nothing""
}

func (s *SimpleMovingAverage) AddPrice(price float64) {
	s.Prices = append(s.Prices, price)
	if len(s.Prices) > s.Period {
		s.Prices = s.Prices[1:]
	}
}

// CalculateSMA calculates the Simple Moving Average based on the stored prices
// and returns the average price.
// If there are no prices, it returns 0.
func (s *SimpleMovingAverage) CalculateSMA() float64 {
	if len(s.Prices) == 0 {
		return 0
	}
	sum := 0.0
	for _, p := range s.Prices {
		sum += p
	}
	var avg = sum / float64(len(s.Prices))
	utils.LogInfo("Calculated SMA: " + utils.FormatCurrency(avg))
	return avg
}

// GenerateSignal generates a trading signal based on the current price and the Simple Moving Average.
func (s *SimpleMovingAverage) GenerateSignal(price float64) string {
	s.AddPrice(price)
	if len(s.Prices) < s.Period {
		return "do nothing" // Not enough data
	}
	sma := s.CalculateSMA()
	if price < sma && s.LastSignal != "buy" {
		s.LastSignal = "buy"
		return "buy"
	}
	if price > sma && s.LastSignal != "sell" {
		s.LastSignal = "sell"
		return "sell"
	}
	return "do nothing"
}

func (sma *SimpleMovingAverage) Execute(price float64) (string, error) {
	signal := sma.GenerateSignal(price)
	if signal != "do nothing" {
		println("Signal:", signal, "at price", fmt.Sprintf("%.2f", price))
	}
	return signal, nil
}

// SMA Crossover strategy (50/200)
type SMACrossover struct {
	ShortPeriod int
	LongPeriod  int
	ShortPrices []float64
	LongPrices  []float64
	LastSignal  string // "buy", "sell", or "do nothing"
}

// AddPrice adds a new price to the Short and Long SMA lists.
func (s *SMACrossover) AddPrice(price float64) {
	s.ShortPrices = append(s.ShortPrices, price)
	if len(s.ShortPrices) > s.ShortPeriod {
		s.ShortPrices = s.ShortPrices[1:]
	}
	s.LongPrices = append(s.LongPrices, price)
	if len(s.LongPrices) > s.LongPeriod {
		s.LongPrices = s.LongPrices[1:]
	}
}

// CalculateSMA calculates the Simple Moving Average based on the provided prices
func (s *SMACrossover) CalculateSMA(prices []float64) float64 {
	if len(prices) == 0 {
		return 0
	}
	sum := 0.0
	for _, p := range prices {
		sum += p
	}
	return sum / float64(len(prices))
}

// GenerateSignal generates a trading signal based on the current
// price and the Simple Moving Average Crossover.
func (s *SMACrossover) GenerateSignal(price float64) string {
	s.AddPrice(price)
	if len(s.ShortPrices) < s.ShortPeriod || len(s.LongPrices) < s.LongPeriod {
		return "do nothing" // Not enough data
	}
	shortSMA := s.CalculateSMA(s.ShortPrices)
	longSMA := s.CalculateSMA(s.LongPrices)
	if shortSMA > longSMA && s.LastSignal != "buy" {
		s.LastSignal = "buy"
		return "buy"
	}
	if shortSMA < longSMA && s.LastSignal != "sell" {
		s.LastSignal = "sell"
		return "sell"
	}
	return "do nothing"
}

// Execute executes the SMACrossover strategy and returns the trading signal.
func (s *SMACrossover) Execute(price float64) (string, error) {
	signal := s.GenerateSignal(price)
	if signal != "do nothing" {
		fmt.Println("SMA Crossover Signal:", signal, "at price", fmt.Sprintf("%.2f", price))
	}
	return signal, nil
}
