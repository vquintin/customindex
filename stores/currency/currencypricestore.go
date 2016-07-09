package currency

import (
	"time"

	"bitbucket.org/virgilequintin/customindex/assets"
	"bitbucket.org/virgilequintin/customindex/stores"
)

type CurrencyPriceStore struct {
	Next stores.PriceStore
}

func (store CurrencyPriceStore) UnitPrice(asset interface{}, date time.Time) (assets.MoneyAmount, error) {
	switch asset := asset.(type) {
	case assets.Currency:
		return assets.MoneyAmount{1.0, asset}, nil
	default:
		return store.Next.UnitPrice(asset, date)
	}
}
