package rpc

import (
	"context"

	"github.com/gagliardetto/solana-go"
)

type GetTokenSupplyResult struct {
	RPCContext
	Value *UiTokenAmount `json:"value"`
}

// GetTokenSupply returns the total supply of an SPL Token type.
func (cl *Client) GetTokenSupply(
	ctx context.Context,
	tokenMint solana.PublicKey, // Pubkey of token Mint to query
	commitment CommitmentType,
) (out *GetTokenSupplyResult, err error) {
	params := []interface{}{tokenMint}
	if commitment != "" {
		params = append(params,
			M{"commitment": commitment},
		)
	}
	err = cl.rpcClient.CallFor(&out, "getTokenSupply", params)
	return
}
