package assets

import "fmt"

// YahooSymbol represents the ticker symbol of a stock on yahoo finance
// (e.g. "FP.PA" for Total S.A.).
type YahooSymbol string

// Equity represents a stock whose price is available on yahoo finance.
// The currency has to be given because the currency of the stock is not
// available on yahoo finance API.
type Equity struct {
	Symbol   YahooSymbol `json:"symbol"`
	Currency Currency    `json:"currency"`
}

func (e Equity) String() string {
	return fmt.Sprintf("%v (%v)", e.Symbol, e.Currency)
}
