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
	store2 := currency.CurrencyPricer{Next: &store1}
	store3 := equity.YahooPricer{Next: &store2}
	rateStore1 := fx.FixerChanger{}
	rateStore2 := cache.NewExchangeRateCache(&rateStore1)
	store4 := index.IndexPricer{Next: &store3, Head: &store1, Changer: &rateStore2}
	store5 := cache.NewPricerCache(&store4)
	store4.Head = &store5
	return &store5
}
