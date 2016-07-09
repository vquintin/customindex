package stores

import (
	"fmt"
	"time"

	"bitbucket.org/virgilequintin/customindex/assets"
)

type FailPriceStore struct {
}

func (store FailPriceStore) UnitPrice(asset interface{}, date time.Time) (assets.MoneyAmount, error) {
	return assets.MoneyAmount{}, fmt.Errorf("Can't find price for asset %v of type %T", asset, asset)
}
