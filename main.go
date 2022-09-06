package main

import (
	"fmt"
	"os"

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

	//Full list of Kujira supported Pools and their swappable tokens.
	//Note that this can be queried on chain, but this is a little easier.
	pairs, err := kujira.GetCoinGeckoPairs("https://api.kujira.app/api/coingecko/pairs")
	if err != nil {
		fmt.Println("Oh no!")
		os.Exit(1)
	}

	for _, pair := range pairs {
		fmt.Printf("Contract pool: %s, base: %s, target: %s\n", pair.PoolID, pair.Base, pair.Target)
		orders := kujira.QueryOrders(pair.PoolID, kujira.BookQueryOptions{Limit: 1}, queryClient)
		kujira.PrintOrders(orders)
	}
}
