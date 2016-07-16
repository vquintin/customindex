package currency

import (
	"time"

	"bitbucket.org/virgilequintin/customindex/assets"
	"bitbucket.org/virgilequintin/customindex/stores"
)

// CurrencyPriceStore is a store used to retrieve the price of a currency.
type CurrencyPriceStore struct {
	Next stores.PriceStore
}

// UnitPrice gives the unit price of a currency.
// It is tautological. The unit price of USD is 1 USD and the unit price of
// EUR is 1 EUR. If the argument is not a Currency, it calls the next store in chain.
func (store CurrencyPriceStore) UnitPrice(asset interface{}, date time.Time) (assets.MoneyAmount, error) {
	switch asset := asset.(type) {
	case assets.Currency:
		return assets.MoneyAmount{Amount: 1.0, Currency: asset}, nil
	default:
		return store.Next.UnitPrice(asset, date)
	}
}
