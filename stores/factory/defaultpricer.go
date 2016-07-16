package factory

import (
	"github.com/vquintin/customindex/stores"
	"github.com/vquintin/customindex/stores/cache"
	"github.com/vquintin/customindex/stores/currency"
	"github.com/vquintin/customindex/stores/equity"
	"github.com/vquintin/customindex/stores/fx"
	"github.com/vquintin/customindex/stores/index"
)

// NewPricer makes a default pricer for Currency, Equity and Index, with caching.
func NewPricer() stores.Pricer {
	store1 := stores.FailPricer{}
	store2 := currency.CurrencyPricer{Next: store1}
	store3 := equity.YahooPricer{Next: store2}
	rateStore1 := fx.FixerChanger{}
	rateStore2 := cache.NewExchangeRateCache(rateStore1)
	store4 := cache.NewPricerCache(store3)
	store5 := index.IndexPricer{Next: &store4, Head: &store4, Changer: &rateStore2}
	return &store5
}
