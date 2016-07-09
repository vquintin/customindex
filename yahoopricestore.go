package customindex

import (
	"fmt"
	"time"

	"github.com/doneland/yquotes"
)

type YahooPriceStore struct {
	next PriceStore
}

func (store YahooPriceStore) UnitPrice(asset interface{}, date time.Time) (MoneyAmount, error) {
	switch asset := asset.(type) {
	case Equity:
		return store.unitPriceForEquity(asset, date)
	default:
		return store.next.UnitPrice(asset, date)
	}
}

func (store YahooPriceStore) unitPriceForEquity(equity Equity, date time.Time) (MoneyAmount, error) {
	start := date.Add(-week)
	prices, err := yquotes.GetDailyHistory(string(equity.Symbol), start, date)
	if err != nil {
		return MoneyAmount{}, err
	}
	if len(prices) != 0 {
		return MoneyAmount{prices[0].Close, equity.Currency}, nil
	}
	return MoneyAmount{}, fmt.Errorf("No value found for %v before %v", equity.Symbol, date)
}
