package customindex

import "time"

type CurrencyPriceStore struct {
	next PriceStore
}

func (store *CurrencyPriceStore) UnitPrice(asset interface{}, date time.Time) (MoneyAmount, error) {
	switch asset := asset.(type) {
	case Currency:
		return MoneyAmount{1.0, asset}, nil
	default:
		return store.next.UnitPrice(asset, date)
	}
}
