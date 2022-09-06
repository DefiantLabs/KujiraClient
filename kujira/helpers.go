package kujira

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

//GetCoinGeckoPairs Gets list of Kujira swappable pairs and the pair Pool.
//The Pool address is the smart contract address on chain.
//Current URL: https://api.kujira.app/api/coingecko/pairs
func GetCoinGeckoPairs(url string) ([]CoinGeckoPair, error) {
	pairs := struct {
		Pairs []CoinGeckoPair `json:"pairs"`
	}{}

	err := getJson(url, &pairs)
	if err != nil {
		return nil, err
	}

	return pairs.Pairs, nil
}

func getJson(url string, target interface{}) error {
	var client = &http.Client{Timeout: 10 * time.Second}

	r, err := client.Get(url)
	if err != nil {
		return err
	}
	defer r.Body.Close()

	return json.NewDecoder(r.Body).Decode(target)
}

func PrintOrders(orders BookResponse) {
	for _, order := range orders.Base {
		//Example: OSMO/axlUSDC pool
		//Selling 'TotalOfferAmount' OSMO at price 'X' each
		fmt.Printf("Offer denom: %s%s, offer price: %f, offer amount: %d\n",
			order.OfferDenom.CW20Type, order.OfferDenom.NativeType,
			order.QuotePrice, order.TotalOfferAmount)
	}
	for _, order := range orders.Quote {
		//Example: OSMO/axlUSDC pool
		//Buying 'TotalOfferAmount' OSMO at price 'X' each
		fmt.Printf("Offer denom: %s%s, offer price: %f, offer amount: %d\n",
			order.OfferDenom.CW20Type, order.OfferDenom.NativeType,
			order.QuotePrice, order.TotalOfferAmount)
	}
}
