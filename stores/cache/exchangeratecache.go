package cache

import (
	"time"

	"bitbucket.org/virgilequintin/customindex/assets"
	"bitbucket.org/virgilequintin/customindex/stores"
)

// ExchangeRateCache is a cache for an exchange rate store.
// It is based on the non-blocking concurrent cache found in
// "The Go Programming Language" book.
type ExchangeRateCache struct {
	memo
}

type exchangeRateCacheKey struct {
	ma   assets.MoneyAmount
	cur  assets.Currency
	date time.Time
}

// Convert convert a money amount in a source currency to a target currency at
// the given date
func (erc ExchangeRateCache) Convert(moneyAmount assets.MoneyAmount, targetCurrency assets.Currency, date time.Time) (assets.MoneyAmount, error) {
	key := exchangeRateCacheKey{assets.MoneyAmount{Amount: 1.0, Currency: moneyAmount.Currency}, targetCurrency, date}
	val, err := erc.get(key)
	return val.(assets.MoneyAmount).Mul(moneyAmount.Amount), err
}

// NewExchangeRateCache makes an ExchangeRateCache caching an exchange rate store
func NewExchangeRateCache(cached stores.ExchangeRateStore) ExchangeRateCache {
	memo := memo{cache: make(map[interface{}]*entry)}
	erc := ExchangeRateCache{memo}
	erc.f = func(key interface{}) (interface{}, error) {
		k := key.(exchangeRateCacheKey)
		return cached.Convert(k.ma, k.cur, k.date)
	}
	return erc
}
