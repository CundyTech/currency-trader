package wallet

type Transaction struct {
	Type   string
	Amount float64
}

type Wallet struct {
	Transactions []Transaction
	Balance      float64
}

func (w Wallet) CalculateBalance(transactions []Transaction) float64 {
	for _, transaction := range transactions {
		switch transaction.Type {
		case "deposit":
			w.Balance += transaction.Amount
		case "withdrawal":
			w.Balance -= transaction.Amount
		}
	}
	return w.Balance
}

func AddTransaction(transactions []Transaction, transactionType string, amount float64) []Transaction {
	newTransaction := Transaction{
		Type:   transactionType,
		Amount: amount,
	}
	transactions = append(transactions, newTransaction)
	return transactions
}
