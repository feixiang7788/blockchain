package blockchain

import (
	"github.com/feixiang7788/blockchain/http"
	"github.com/feixiang7788/blockchain/identity/CurrencyType"
	"github.com/feixiang7788/blockchain/implement"
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
