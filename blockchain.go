package blockchain

import (
	"game_server/pkg/blockchain/implement"
	"game_server/pkg/http"
	"game_server/pkg/identity/CurrencyType"
)

type BlockchainInfo interface {
	CurrentBlockHeight() (int, error)
	BlockInfo(blockHeightNumber int) (string, error)
}

func NewBlockchainInfo(currencyType CurrencyType.CurrencyType) BlockchainInfo {
	switch currencyType {
	case CurrencyType.Arbis:
		return implement.NewArbis(http.NewArbisHttpClient())
	case CurrencyType.Solana:
		return implement.NewSolana(http.NewSolanaHttpClient())
	default:
		return nil
	}
}
