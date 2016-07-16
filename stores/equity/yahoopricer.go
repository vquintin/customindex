package equity

import (
	"fmt"
	"time"

	"github.com/doneland/yquotes"
	"github.com/vquintin/customindex/assets"
	"github.com/vquintin/customindex/stores"
)

// YahooPricer is a pricer able to price equities using the Yahoo finance API
type YahooPricer struct {
	Next stores.Pricer
}

// UnitPrice gives the price of an Equity at the given date. If the asset is not an equity, it calls the
// next pricer in chain.
func (store YahooPricer) UnitPrice(asset interface{}, date time.Time) (assets.MoneyAmount, error) {
	switch asset := asset.(type) {
	case assets.Equity:
		return store.unitPriceForEquity(asset, date)
	default:
		return store.Next.UnitPrice(asset, date)
	}
}

const week = 168 * time.Hour

func (store YahooPricer) unitPriceForEquity(equity assets.Equity, date time.Time) (assets.MoneyAmount, error) {
	start := date.Add(-week)
	prices, err := yquotes.GetDailyHistory(string(equity.Symbol), start, date)
	if err != nil {
		return assets.MoneyAmount{}, err
	}
	if len(prices) != 0 {
		return assets.MoneyAmount{Amount: prices[0].Close, Currency: equity.Currency}, nil
	}
	return assets.MoneyAmount{}, fmt.Errorf("No value found for %v before %v", equity.Symbol, date)
}
