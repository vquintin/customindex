package customindex

import "time"

type IndexPriceStore struct {
	next      PriceStore
	head      PriceStore
	rateStore ExchangeRateStore
}

func (store IndexPriceStore) UnitPrice(asset interface{}, date time.Time) (MoneyAmount, error) {
	switch asset := asset.(type) {
	case Index:
		return store.unitPriceForIndex(asset, date)
	default:
		return store.next.UnitPrice(asset, date)
	}
}

type ratioAndErr struct {
	ratio float64
	err   error
}

const week = 168 * time.Hour

func (store IndexPriceStore) unitPriceForIndex(index Index, date time.Time) (MoneyAmount, error) {
	quit := make(chan interface{})
	c := make(chan ratioAndErr)
	var weightSum float64
	for k, v := range index.weights {
		go func(asset interface{}, v float64) {
			ratio, err := store.performanceRatio(asset, index.creation, date, index.initialValue.Currency)
			select {
			case c <- ratioAndErr{v * ratio, err}:
			case <-quit:
			}
		}(k, v)
		weightSum += v
	}
	var ratioSum float64
	for i, n := 0, len(index.weights); i < n; i++ {
		rae := <-c
		if rae.err != nil {
			close(quit)
			return MoneyAmount{}, rae.err
		}
		ratioSum += rae.ratio
	}
	return index.initialValue.Mul(ratioSum / weightSum), nil
}

type maAndErr struct {
	ma  MoneyAmount
	err error
}

func (store IndexPriceStore) performanceRatio(asset interface{}, start time.Time, end time.Time, currency Currency) (float64, error) {
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

func (store IndexPriceStore) capitalValueInCurrency(asset interface{}, date time.Time, currency Currency) maAndErr {
	ma, err := store.head.UnitPrice(asset, date)
	if err != nil {
		return maAndErr{ma, err}
	}
	cma, err := store.rateStore.Convert(ma, currency, date)
	return maAndErr{cma, err}
}
