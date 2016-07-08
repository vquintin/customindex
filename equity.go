package customindex

import (
	"fmt"
	"time"

	"github.com/doneland/yquotes"
)

type YahooSymbol string

type Equity struct {
	Symbol   YahooSymbol `json:"symbol"`
	Currency Currency    `json:"currency"`
}

func (e Equity) String() string {
	return fmt.Sprintf("%v (%v)", e.Symbol, e.Currency)
}

const week = 168 * time.Hour

func (e Equity) UnitPrice(date time.Time) (MoneyAmount, error) {
	start := date.Add(-week)
	prices, err := yquotes.GetDailyHistory(string(e.Symbol), start, date)
	if err != nil {
		return MoneyAmount{}, err
	}
	if len(prices) != 0 {
		return MoneyAmount{prices[0].Close, e.Currency}, nil
	}
	return MoneyAmount{}, fmt.Errorf("No value found for %v before %v", e.Symbol, date)
}
