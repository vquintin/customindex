package cache

import (
	"time"

	"github.com/vquintin/customindex/assets"
	"github.com/vquintin/customindex/stores"
)

// ChangerCache is a cache for an exchange rate store.
// It is based on the non-blocking concurrent cache found in
// "The Go Programming Language" book.
type ChangerCache struct {
	memo
}

type changerCacheKey struct {
	ma   assets.MoneyAmount
	cur  assets.Currency
	date time.Time
}

// Change converts a money amount in a source currency to a target currency at
// the given date
func (erc *ChangerCache) Change(moneyAmount assets.MoneyAmount, targetCurrency assets.Currency, date time.Time) (assets.MoneyAmount, error) {
	key := changerCacheKey{assets.MoneyAmount{Amount: 1.0, Currency: moneyAmount.Currency}, targetCurrency, date}
	val, err := erc.get(key)
	return val.(assets.MoneyAmount).Mul(moneyAmount.Amount), err
}

// NewExchangeRateCache makes an ExchangeRateCache caching an exchange rate store
func NewExchangeRateCache(cached stores.Changer) ChangerCache {
	memo := memo{cache: make(map[interface{}]*entry)}
	erc := ChangerCache{memo}
	erc.f = func(key interface{}) (interface{}, error) {
		k := key.(changerCacheKey)
		return cached.Change(k.ma, k.cur, k.date)
	}
	return erc
}
