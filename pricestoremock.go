package customindex

import (
	"fmt"
	"time"
)

type assetAndDate struct {
	asset interface{}
	date  time.Time
}

type priceStoreMock struct {
	Called bool
	Values map[assetAndDate]MoneyAmount
}

func (mock *priceStoreMock) UnitPrice(asset interface{}, date time.Time) (MoneyAmount, error) {
	mock.Called = true
	aAD := assetAndDate{asset, date}
	if v, ok := mock.Values[aAD]; ok {
		return v, nil
	}
	return MoneyAmount{}, fmt.Errorf("No values in mock for %v at %v. Map: %v", asset, date, mock.Values)
}
