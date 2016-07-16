package fx

import (
	"fmt"
	"time"

	"github.com/openprovider/ecbrates"
	"github.com/vquintin/customindex/assets"
)

// EcbChanger is a money amount changer which uses the rates found on the European
// Central Bank website
type EcbChanger struct {
}

// Change converts a money amount to another currency.
func (store EcbChanger) Change(p assets.MoneyAmount, other assets.Currency, date time.Time) (assets.MoneyAmount, error) {
	if p.Currency == other {
		return p, nil
	}
	rates, err := ecbrates.Load()
	var before ecbrates.Rates
	for _, v := range rates {
		rateDate, err := time.Parse("2006-01-02", v.Date)
		if err != nil {
			return assets.MoneyAmount{}, err
		}
		if rateDate.Before(date) {
			before = v
			break
		}
	}
	fmt.Println(before)
	amount, err := before.Convert(p.Amount, ecbrates.Currency(p.Currency), ecbrates.Currency(other))
	return assets.MoneyAmount{amount, other}, err
}
