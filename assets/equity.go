package assets

import "fmt"

type YahooSymbol string

type Equity struct {
	Symbol   YahooSymbol `json:"symbol"`
	Currency Currency    `json:"currency"`
}

func (e Equity) String() string {
	return fmt.Sprintf("%v (%v)", e.Symbol, e.Currency)
}
