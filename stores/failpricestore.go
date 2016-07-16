package stores

import (
	"fmt"
	"time"

	"github.com/vquintin/customindex/assets"
)

// FailPricer is an always failing Pricer. It is useful to terminate the chain of pricers.
type FailPricer struct {
}

// UnitPrice always returns an error.
func (store FailPricer) UnitPrice(asset interface{}, date time.Time) (assets.MoneyAmount, error) {
	return assets.MoneyAmount{}, fmt.Errorf("Can't find price for asset %v of type %T", asset, asset)
}
