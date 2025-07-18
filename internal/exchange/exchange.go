package exchange

import (
	"crypto/tls"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

type ForexRate struct {
	Code6509 string `json:"6509"`
	ConidEx  string `json:"conidEx"`
	Updated  int64  `json:"_updated"`
	Conid    int    `json:"conid"`
	Code6119 string `json:"6119"`
	ServerID string `json:"server_id"`
	Price    string `json:"31"`
}

func PlaceOrder(orderType string, amount float64, price float64) (string, error) {
	// Implement logic to execute buy/sell orders on the exchange
	return "", nil
}

// getConidForPair returns the IBKR conid for a given forex pair (limited to common pairs)
func getConidForPair(pair string) string {
	conidMap := map[string]string{
		"GBP/USD": "12087792",
		"EUR/USD": "12087777",
		"USD/JPY": "12087817",
		"EUR/GBP": "12087782",
	}
	return conidMap[pair]
}

// GetForexRate fetches the forex rate for a currency pair using IBKR Web API and cookie authentication
func GetForexRate(baseURL, sessionCookie, pair string) (*ForexRate, error) {

	conid := getConidForPair(pair)
	if conid == "" {
		return nil, fmt.Errorf("unsupported pair: %s", pair)
	}

	client := &http.Client{}
	url := fmt.Sprintf("%s/v1/api/iserver/marketdata/snapshot?conids=%s&fields=31", baseURL, conid)
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Cookie", sessionCookie)
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 200 {
		return nil, errors.New("failed to fetch forex rate, exchange returned status: " + resp.Status)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	rate := []ForexRate{}
	if err := json.Unmarshal(body, &rate); err != nil {
		return nil, err
	}

	return &rate[0], nil
}
