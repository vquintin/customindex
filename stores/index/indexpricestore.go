package index

import (
	"time"

	"bitbucket.org/virgilequintin/customindex/assets"
	"bitbucket.org/virgilequintin/customindex/stores"
)

type IndexPriceStore struct {
	Next      stores.PriceStore
	Head      stores.PriceStore
	RateStore stores.ExchangeRateStore
}

func (store IndexPriceStore) UnitPrice(asset interface{}, date time.Time) (assets.MoneyAmount, error) {
	switch asset := asset.(type) {
	case assets.Index:
		return store.unitPriceForIndex(asset, date)
	default:
		return store.Next.UnitPrice(asset, date)
	}
}

type ratioAndErr struct {
	ratio float64
	err   error
}

const week = 168 * time.Hour

func (store IndexPriceStore) unitPriceForIndex(index assets.Index, date time.Time) (assets.MoneyAmount, error) {
	quit := make(chan interface{})
	c := make(chan ratioAndErr)
	var weightSum float64
	for k, v := range index.Weights {
		go func(asset interface{}, v float64) {
			ratio, err := store.performanceRatio(asset, index.Creation, date, index.InitialValue.Currency)
			select {
			case c <- ratioAndErr{v * ratio, err}:
			case <-quit:
			}
		}(k, v)
		weightSum += v
	}
	var ratioSum float64
	for i, n := 0, len(index.Weights); i < n; i++ {
		rae := <-c
		if rae.err != nil {
			close(quit)
			return assets.MoneyAmount{}, rae.err
		}
		ratioSum += rae.ratio
	}
	return index.InitialValue.Mul(ratioSum / weightSum), nil
}

type maAndErr struct {
	ma  assets.MoneyAmount
	err error
}

func (store IndexPriceStore) performanceRatio(asset interface{}, start time.Time, end time.Time, currency assets.Currency) (float64, error) {
	initialValue := make(chan maAndErr)
	finalValue := make(chan maAndErr)
	go func() {
		initialValue <- store.capitalValueInCurrency(asset, start, currency)
	}()
	go func() {
		finalValue <- store.capitalValueInCurrency(asset, end, currency)
	}()
	initialResult := <-initialValue
	finalResult := <-finalValue
	if initialResult.err != nil {
		return 0, initialResult.err
	} else if finalResult.err != nil {
		return 0, finalResult.err
	} else {
		return finalResult.ma.Div(initialResult.ma)
	}
}

func (store IndexPriceStore) capitalValueInCurrency(asset interface{}, date time.Time, currency assets.Currency) maAndErr {
	ma, err := store.Head.UnitPrice(asset, date)
	if err != nil {
		return maAndErr{ma, err}
	}
	cma, err := store.RateStore.Convert(ma, currency, date)
	return maAndErr{cma, err}
}
