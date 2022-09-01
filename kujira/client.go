package kujira

import (
	"os"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	authTypes "github.com/cosmos/cosmos-sdk/x/auth/types"
)

func GetTxClient(chain string, node string, keyringDir string, keyringBackend string, fromFlag string) *client.Context {
	encodingConfig := MakeEncodingConfig()
	clientCtx := client.Context{
		ChainID:      chain,
		NodeURI:      node,
		KeyringDir:   keyringDir,
		GenerateOnly: false,
	}

	ctxKeyring, krErr := client.NewKeyringFromBackend(clientCtx, keyringBackend)
	if krErr != nil {
		return nil
	}

	clientCtx = clientCtx.WithKeyring(ctxKeyring)

	//Where node is the node RPC URI
	rpcClient, rpcErr := client.NewClientFromNode(node)

	if rpcErr != nil {
		return nil
	}

	fromAddr, fromName, _, err := client.GetFromFields(clientCtx.Keyring, fromFlag, clientCtx.GenerateOnly)
	if err != nil {
		return nil
	}

	clientCtx = clientCtx.WithCodec(encodingConfig.Marshaler).
		WithChainID(chain).
		WithFrom(fromFlag).
		WithFromAddress(fromAddr).
		WithFromName(fromName).
		WithInterfaceRegistry(encodingConfig.InterfaceRegistry).
		WithTxConfig(encodingConfig.TxConfig).
		WithLegacyAmino(encodingConfig.Amino).
		WithInput(os.Stdin).
		WithAccountRetriever(authTypes.AccountRetriever{}).
		WithBroadcastMode(flags.BroadcastAsync).
		WithHomeDir(keyringDir).
		WithNodeURI(node).
		WithClient(rpcClient).
		WithSkipConfirmation(true)
		//WithViper("KUJIRA").

	return &clientCtx
}
