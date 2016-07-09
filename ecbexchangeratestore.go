package customindex

import (
	"fmt"
	"time"

	"github.com/openprovider/ecbrates"
)

type EcbExchangeRateStore struct {
}

func (store EcbExchangeRateStore) Convert(p MoneyAmount, other Currency, date time.Time) (MoneyAmount, error) {
	if p.Currency == other {
		return p, nil
	}
	rates, err := ecbrates.Load()
	var before ecbrates.Rates
	for _, v := range rates {
		rateDate, err := time.Parse("2006-01-02", v.Date)
		if err != nil {
			return MoneyAmount{}, err
		}
		if rateDate.Before(date) {
			before = v
			break
		}
	}
	fmt.Println(before)
	amount, err := before.Convert(p.Amount, ecbrates.Currency(p.Currency), ecbrates.Currency(other))
	return MoneyAmount{amount, other}, err
}
