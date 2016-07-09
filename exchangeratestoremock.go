package customindex

import (
	"fmt"
	"time"
)

type currencyPairWithDate struct {
	date time.Time
	l    Currency
	r    Currency
}

type exchangeRateStoreMock struct {
	Rates map[currencyPairWithDate]float64
}

func (store *exchangeRateStoreMock) Convert(moneyAmount MoneyAmount, targetCurrency Currency, date time.Time) (MoneyAmount, error) {
	sourceCurrency := moneyAmount.Currency
	cp := currencyPairWithDate{date, sourceCurrency, targetCurrency}
	if sourceCurrency == targetCurrency {
		return moneyAmount, nil
	}
	v, ok := store.Rates[cp]
	if ok {
		return MoneyAmount{moneyAmount.Amount * v, targetCurrency}, nil
	}
	cp = currencyPairWithDate{date, targetCurrency, sourceCurrency}
	v, ok = store.Rates[cp]
	if ok {
		return MoneyAmount{moneyAmount.Amount / v, targetCurrency}, nil
	}
	return MoneyAmount{}, fmt.Errorf("No rates from %v to %v", sourceCurrency, targetCurrency)
}
