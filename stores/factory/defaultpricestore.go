package factory

import (
	"bitbucket.org/virgilequintin/customindex/stores"
	"bitbucket.org/virgilequintin/customindex/stores/currency"
	"bitbucket.org/virgilequintin/customindex/stores/equity"
	"bitbucket.org/virgilequintin/customindex/stores/fx"
	"bitbucket.org/virgilequintin/customindex/stores/index"
)

func NewPriceStore() stores.PriceStore {
	store1 := stores.FailPriceStore{}
	store2 := currency.CurrencyPriceStore{store1}
	store3 := equity.YahooPriceStore{&store2}
	rateStore := fx.FixerExchangeRateStore{}
	store4 := index.IndexPriceStore{&store3, &store3, &rateStore}
	store4.Head = store4
	return store4
}
