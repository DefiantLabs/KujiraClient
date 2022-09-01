package kujira

type Book struct { //[]BookResponse
}

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
