package equity

import (
	"fmt"
	"time"

	"bitbucket.org/virgilequintin/customindex/assets"
	"bitbucket.org/virgilequintin/customindex/stores"

	"github.com/doneland/yquotes"
)

type YahooPriceStore struct {
	Next stores.PriceStore
}

func (store YahooPriceStore) UnitPrice(asset interface{}, date time.Time) (assets.MoneyAmount, error) {
	switch asset := asset.(type) {
	case assets.Equity:
		return store.unitPriceForEquity(asset, date)
	default:
		return store.Next.UnitPrice(asset, date)
	}
}

const week = 168 * time.Hour

func (store YahooPriceStore) unitPriceForEquity(equity assets.Equity, date time.Time) (assets.MoneyAmount, error) {
	start := date.Add(-week)
	prices, err := yquotes.GetDailyHistory(string(equity.Symbol), start, date)
	if err != nil {
		return assets.MoneyAmount{}, err
	}
	if len(prices) != 0 {
		return assets.MoneyAmount{prices[0].Close, equity.Currency}, nil
	}
	return assets.MoneyAmount{}, fmt.Errorf("No value found for %v before %v", equity.Symbol, date)
}
