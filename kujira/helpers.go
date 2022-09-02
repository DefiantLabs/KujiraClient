package kujira

import "fmt"

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
