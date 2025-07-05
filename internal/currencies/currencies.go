package currencies

type Currency struct{}

func (c Currency) GetSupportedCurrencies() []string {
	return []string{
		"USD", "EUR", "GBP", "JPY", "AUD", "CAD", "CHF", "CNY", "SEK", "NZD",
		"MXN", "SGD", "HKD", "NOK", "KRW", "TRY", "RUB", "INR", "BRL", "ZAR",
	}
}
func (c Currency) IsSupportedCurrency(currency string) bool {
	supportedCurrencies := c.GetSupportedCurrencies()
	for _, c := range supportedCurrencies {
		if c == currency {
			return true
		}
	}
	return false
}

func (c Currency) GetCurrencySymbol(currency string) string {
	switch currency {
	case "USD":
		return "$"
	case "EUR":
		return "€"
	case "GBP":
		return "£"
	case "JPY":
		return "¥"
	case "AUD":
		return "A$"
	case "CAD":
		return "C$"
	case "CHF":
		return "CHF"
	case "CNY":
		return "¥"
	case "SEK":
		return "kr"
	case "NZD":
		return "NZ$"
	default:
		return currency // Return the currency code if no symbol is defined
	}
}

func (c Currency) GetCurrencyName(currency string) string {
	switch currency {
	case "USD":
		return "United States Dollar"
	case "EUR":
		return "Euro"
	case "GBP":
		return "British Pound Sterling"
	case "JPY":
		return "Japanese Yen"
	case "AUD":
		return "Australian Dollar"
	case "CAD":
		return "Canadian Dollar"
	case "CHF":
		return "Swiss Franc"
	case "CNY":
		return "Chinese Yuan Renminbi"
	case "SEK":
		return "Swedish Krona"
	case "NZD":
		return "New Zealand Dollar"
	default:
		return currency // Return the currency code if no name is defined
	}
}
