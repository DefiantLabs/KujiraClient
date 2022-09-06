package kujira

type BookResponse struct {
	Base  []PoolResponse `json:"base"`
	Quote []PoolResponse `json:"quote"`
}

type PoolResponse struct {
	QuotePrice       float64    `json:"quote_price,string"`
	OfferDenom       OfferDenom `json:"offer_denom"`
	TotalOfferAmount int64      `json:"total_offer_amount,string"`
}

//Note: Will always be one of 'NativeType' or 'CW20Type'
type OfferDenom struct {
	NativeType string `json:"native"`
	CW20Type   string `json:"cw20"`
}

type BookQuery struct {
	Book BookQueryOptions `json:"book"`
}

//For use with 'QueryOrders' (Kujira's Book QueryMsg)
//Leave as zero values to ignore a query option
type BookQueryOptions struct {
	Limit  uint8 `json:"limit,omitempty"`
	Offset uint8 `json:"offset,omitempty"`
}

// type CoinGeckoPairs struct {
// 	Pairs []CoinGeckoPair `json:"pairs"`
// }

type CoinGeckoPair struct {
	Base     string `json:"base"`
	PoolID   string `json:"pool_id"`
	Target   string `json:"target"`
	TickerID string `json:"ticker_id"`
}
