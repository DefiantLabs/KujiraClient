package kujira

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/CosmWasm/wasmd/x/wasm/types"
)

func QueryOrders(contract string, options BookQueryOptions, client types.QueryClient) (orders BookResponse) {
	query, e := json.Marshal(BookQuery{Book: options})
	if e != nil {
		fmt.Printf("Query options error %f\n", e)
	}
	// query := "{\"book\": {}}"
	// if options.Limit != 0 {
	// 	query = fmt.Sprintf("{\"book\": {\"limit\":%d}}", options.Limit)
	// }

	res, err := client.SmartContractState(
		context.Background(),
		&types.QuerySmartContractStateRequest{
			Address:   contract,
			QueryData: []byte(query),
		},
	)

	if err != nil {
		fmt.Printf("Error %f querying Kujira contract\n", err)
	} else {
		marshalErr := json.Unmarshal(res.Data, &orders)
		if marshalErr != nil {
			fmt.Printf("Unmarshal error %f \n", marshalErr)
		}
	}

	return
}
