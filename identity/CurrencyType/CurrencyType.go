package CurrencyType

type CurrencyType int

const (
	Arbis CurrencyType = iota + 1
	Solana
)

func (c CurrencyType) Int() int {
	return int(c)
}
