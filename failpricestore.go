package customindex

import (
	"fmt"
	"time"
)

type FailPriceStore struct {
}

func (store FailPriceStore) UnitPrice(asset interface{}, date time.Time) (MoneyAmount, error) {
	return MoneyAmount{}, fmt.Errorf("Can't find price for asset %v of type %T", asset, asset)
}
