package factory

import (
	"bitbucket.org/virgilequintin/customindex/stores"
	"bitbucket.org/virgilequintin/customindex/stores/cache"
	"bitbucket.org/virgilequintin/customindex/stores/currency"
	"bitbucket.org/virgilequintin/customindex/stores/equity"
	"bitbucket.org/virgilequintin/customindex/stores/fx"
	"bitbucket.org/virgilequintin/customindex/stores/index"
)

func NewPriceStore() stores.PriceStore {
	store1 := stores.FailPriceStore{}
	store2 := currency.CurrencyPriceStore{store1}
	store3 := equity.YahooPriceStore{store2}
	rateStore1 := fx.FixerExchangeRateStore{}
	rateStore2 := cache.NewExchangeRateCache(rateStore1)
	store4 := index.IndexPriceStore{store3, store3, rateStore2}
	store5 := cache.NewPriceCache(store4)
	store4.Head = store5
	return store5
}
