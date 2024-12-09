package blockchain

import (
	"blockchain/http"
	"blockchain/identity/CurrencyType"
	"blockchain/implement"
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
