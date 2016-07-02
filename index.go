package customindex

import "time"

type Weight struct {
	Capital Capital `json:"capital"`
	Weight  float64 `json:"weight"`
}

type Index struct {
	Name         string      `json:"name"`
	InitialValue MoneyAmount `json:"initialValue"`
	Creation     time.Time   `json:"creation"`
	Weights      []Weight    `json:"weights"`
}

type ratioAndErr struct {
	ratio float64
	err   error
}

func (index Index) Value(date time.Time) (MoneyAmount, error) {
	quit := make(chan interface{})
	c := make(chan ratioAndErr)
	weightSum := 0.0
	for _, v := range index.Weights {
		go func(k Capital, v float64) {
			ratio, err := performanceRatio(k, index.Creation, date, index.InitialValue.Currency)
			select {
			case c <- ratioAndErr{v * ratio, err}:
			case <-quit:
			}
		}(v.Capital, v.Weight)
		weightSum += v.Weight
	}
	ratioSum := 0.0
	for i, n := 0, len(index.Weights); i < n; i++ {
		rae := <-c
		if rae.err != nil {
			close(quit)
			return MoneyAmount{}, rae.err
		}
		ratioSum += rae.ratio
	}
	return index.InitialValue.Mul(ratioSum / weightSum), nil
}

type maAndErr struct {
	ma  MoneyAmount
	err error
}

func performanceRatio(capital Capital, start time.Time, end time.Time, currency Currency) (float64, error) {
	initialValue := make(chan maAndErr)
	finalValue := make(chan maAndErr)
	go func() {
		initialValue <- capitalValueInCurrency(capital, start, currency)
	}()
	go func() {
		finalValue <- capitalValueInCurrency(capital, end, currency)
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

func capitalValueInCurrency(capital Capital, date time.Time, currency Currency) maAndErr {
	ma, err := capital.Value(date)
	if err != nil {
		return maAndErr{ma, err}
	}
	cma, err := ma.Convert(currency, date)
	return maAndErr{cma, err}
}
