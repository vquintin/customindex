package mock

import (
	"fmt"
	"time"

	"bitbucket.org/virgilequintin/customindex/assets"
)

type AssetAndDate struct {
	Asset interface{}
	Date  time.Time
}

type PriceStoreMock struct {
	Called bool
	Values map[AssetAndDate]assets.MoneyAmount
}

func (mock *PriceStoreMock) UnitPrice(asset interface{}, date time.Time) (assets.MoneyAmount, error) {
	mock.Called = true
	aAD := AssetAndDate{asset, date}
	if v, ok := mock.Values[aAD]; ok {
		return v, nil
	}
	return assets.MoneyAmount{}, fmt.Errorf("No values in mock for %v at %v. Map: %v", asset, date, mock.Values)
}
