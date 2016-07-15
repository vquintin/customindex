package cache

import (
	"time"

	"bitbucket.org/virgilequintin/customindex/assets"
	"bitbucket.org/virgilequintin/customindex/stores"
)

type PriceCache struct {
	memo
}

type priceCacheKey struct {
	asset interface{}
	date  time.Time
}

func (pc PriceCache) UnitPrice(asset interface{}, date time.Time) (assets.MoneyAmount, error) {
	key := priceCacheKey{asset, date}
	val, err := pc.Get(key)
	return val.(assets.MoneyAmount), err
}

func NewPriceCache(cached stores.PriceStore) PriceCache {
	memo := memo{cache: make(map[interface{}]*entry)}
	pc := PriceCache{memo}
	pc.f = func(key interface{}) (interface{}, error) {
		k := key.(priceCacheKey)
		return cached.UnitPrice(k.asset, k.date)
	}
	return pc
}
