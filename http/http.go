package http

import "github.com/go-resty/resty/v2"

type ArbisHttpClient = *resty.Client
type SolanaHttpClient = *resty.Client

const (
	ArbisBaseUrl  = "https://api.arbiscan.io/"
	SolanaBaseUrl = "https://api.mainnet-beta.solana.com"
)

func NewArbisHttpClient() ArbisHttpClient {
	return resty.New().SetBaseURL(ArbisBaseUrl)
}

func NewSolanaHttpClient() SolanaHttpClient {
	return resty.New().SetBaseURL(SolanaBaseUrl)
}
