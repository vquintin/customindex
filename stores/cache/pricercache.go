package cache

import (
	"time"

	"bitbucket.org/virgilequintin/customindex/assets"
	"bitbucket.org/virgilequintin/customindex/stores"
)

// PricerCache is a cache for an exchange rate store.
// It is based on the non-blocking concurrent cache found in
// "The Go Programming Language" book.
type PricerCache struct {
	memo
}

type pricerCacheKey struct {
	asset interface{}
	date  time.Time
}

// UnitPrice gives the price of an asset at the closest date before the given date.
// The cached store is guaranteed to be called only once on an asset/date pair except
// when the asset is non hashable.
func (pc PricerCache) UnitPrice(asset interface{}, date time.Time) (assets.MoneyAmount, error) {
	key := pricerCacheKey{asset, date}
	val, err := pc.get(key)
	return val.(assets.MoneyAmount), err
}

// NewPricerCache makes a price store cache caching the given store.
func NewPricerCache(cached stores.Pricer) PricerCache {
	memo := newGoPLCache(func(key interface{}) (interface{}, error) {
		k := key.(pricerCacheKey)
		return cached.UnitPrice(k.asset, k.date)
	})
	pc := PricerCache{*memo}
	return pc
}
