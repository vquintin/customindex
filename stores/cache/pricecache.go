package cache

import (
	"sync"
	"time"

	"bitbucket.org/virgilequintin/customindex/assets"
	"bitbucket.org/virgilequintin/customindex/stores"
)

type PriceCache struct {
	cached stores.PriceStore
	mutex  sync.Mutex
	cache  map[key]*entry
}

type key struct {
	asset interface{}
	date  time.Time
}

type entry struct {
	res   result
	ready chan struct{}
}

type result struct {
	ma  assets.MoneyAmount
	err error
}

func (pc PriceCache) UnitPrice(asset interface{}, date time.Time) (assets.MoneyAmount, error) {
	key := key{asset, date}
	pc.mutex.Lock()
	e := pc.cache[key]
	if e == nil {
		e = &entry{ready: make(chan struct{})}
		pc.cache[key] = e
		pc.mutex.Unlock()
		e.res.ma, e.res.err = pc.cached.UnitPrice(asset, date)
		close(e.ready)
	} else {
		pc.mutex.Unlock()
		<-e.ready
	}
	return e.res.ma, e.res.err
}

func NewPriceCache(cached stores.PriceStore) PriceCache {
	return PriceCache{cached: cached, cache: make(map[key]*entry)}
}
