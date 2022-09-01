package main

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/CosmWasm/wasmd/x/wasm/types"
	"github.com/DefiantLabs/KujiraClient/kujira"
)

func main() {
	//Note, this app has only been tested for queries, not transactions
	chain := "kaiyo-1"
	node := "https://rpc.kaiyo.kujira.setten.io:443"
	keyringDir := "/home/kyle/.kujira"

	clientCtx := kujira.GetTxClient(chain, node, keyringDir, "test", "kyle")
	queryClient := types.NewQueryClient(clientCtx)

	//You can get a full list of pairs at https://api.kujira.app/api/coingecko/pairs
	osmoAxlUsdcContractAddress := "kujira1aakfpghcanxtc45gpqlx8j3rq0zcpyf49qmhm9mdjrfx036h4z5sfmexun"
	query := "{\"book\": {}}"

	res, err := queryClient.SmartContractState(
		context.Background(),
		&types.QuerySmartContractStateRequest{
			Address:   osmoAxlUsdcContractAddress,
			QueryData: []byte(query),
		},
	)

	if err != nil {
		fmt.Printf("Error %f querying Kujira contract\n", err)
	} else {
		var orderbook kujira.BookResponse
		marshalErr := json.Unmarshal(res.Data, &orderbook)
		if marshalErr != nil {
			fmt.Printf("Unmarshal error %f \n", marshalErr)
		}
	}

}
