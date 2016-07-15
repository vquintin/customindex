package cache

import (
	"time"

	"bitbucket.org/virgilequintin/customindex/assets"
	"bitbucket.org/virgilequintin/customindex/stores"
)

type ExchangeRateCache struct {
	memo
}

type exchangeRateCacheKey struct {
	ma   assets.MoneyAmount
	cur  assets.Currency
	date time.Time
}

func (erc ExchangeRateCache) Convert(moneyAmount assets.MoneyAmount, targetCurrency assets.Currency, date time.Time) (assets.MoneyAmount, error) {
	key := exchangeRateCacheKey{assets.MoneyAmount{1.0, moneyAmount.Currency}, targetCurrency, date}
	val, err := erc.get(key)
	return val.(assets.MoneyAmount).Mul(moneyAmount.Amount), err
}

func NewExchangeRateCache(cached stores.ExchangeRateStore) ExchangeRateCache {
	memo := memo{cache: make(map[interface{}]*entry)}
	erc := ExchangeRateCache{memo}
	erc.f = func(key interface{}) (interface{}, error) {
		k := key.(exchangeRateCacheKey)
		return cached.Convert(k.ma, k.cur, k.date)
	}
	return erc
}
