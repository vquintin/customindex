package mock

import (
	"fmt"
	"time"

	"github.com/vquintin/customindex/assets"
)

type CurrencyPairWithDate struct {
	Date time.Time
	L    assets.Currency
	R    assets.Currency
}

type ChangerMock struct {
	Rates map[CurrencyPairWithDate]float64
}

func (store *ChangerMock) Change(moneyAmount assets.MoneyAmount, targetCurrency assets.Currency, date time.Time) (assets.MoneyAmount, error) {
	sourceCurrency := moneyAmount.Currency
	cp := CurrencyPairWithDate{date, sourceCurrency, targetCurrency}
	if sourceCurrency == targetCurrency {
		return moneyAmount, nil
	}
	v, ok := store.Rates[cp]
	if ok {
		return assets.MoneyAmount{moneyAmount.Amount * v, targetCurrency}, nil
	}
	cp = CurrencyPairWithDate{date, targetCurrency, sourceCurrency}
	v, ok = store.Rates[cp]
	if ok {
		return assets.MoneyAmount{moneyAmount.Amount / v, targetCurrency}, nil
	}
	return assets.MoneyAmount{}, fmt.Errorf("No rates from %v to %v", sourceCurrency, targetCurrency)
}
