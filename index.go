package customindex

import (
	"encoding/json"
	"time"
)

type Index struct {
	name         string
	initialValue MoneyAmount
	creation     time.Time
	weights      map[Asset]float64
}

type IndexSP struct {
	Name         string           `json:"name"`
	InitialValue MoneyAmount      `json:"initialValue"`
	Creation     time.Time        `json:"creation"`
	Currencies   []CurrencyWeight `json:"currencies"`
	Equities     []EquityWeight   `json:"equities"`
	Indexes      []IndexWeight    `json:"indexes"`
}

type CurrencyWeight struct {
	Currency
	Weight float64
}

type EquityWeight struct {
	Equity
	Weight float64
}

type IndexWeight struct {
	Index
	Weight float64
}

func (i Index) MarshalJSON() ([]byte, error) {
	sp := IndexSP{i.name, i.initialValue, i.creation, []CurrencyWeight{}, []EquityWeight{}, []IndexWeight{}}
	for k, v := range i.weights {
		switch k := k.(type) {
		case Currency:
			sp.Currencies = append(sp.Currencies, CurrencyWeight{k, v})
		case Equity:
			sp.Equities = append(sp.Equities, EquityWeight{k, v})
		case Index:
			sp.Indexes = append(sp.Indexes, IndexWeight{k, v})
		}
	}
	return json.Marshal(sp)
}

func (index *Index) UnmarshalJSON(data []byte) error {
	sp := IndexSP{}
	err := json.Unmarshal(data, &sp)
	if err != nil {
		return err
	}
	index.name = sp.Name
	index.creation = sp.Creation
	index.initialValue = sp.InitialValue
	index.weights = make(map[Asset]float64)
	for _, v := range sp.Currencies {
		index.weights[v.Currency] = v.Weight
	}
	for _, v := range sp.Equities {
		index.weights[v.Equity] = v.Weight
	}
	for _, v := range sp.Indexes {
		index.weights[v.Index] = v.Weight
	}
	return nil
}

type ratioAndErr struct {
	ratio float64
	err   error
}

func (index Index) UnitPrice(date time.Time) (MoneyAmount, error) {
	quit := make(chan interface{})
	c := make(chan ratioAndErr)
	var weightSum float64
	for k, v := range index.weights {
		go func(k Asset, v float64) {
			ratio, err := performanceRatio(k, index.creation, date, index.initialValue.Currency)
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

func performanceRatio(asset Asset, start time.Time, end time.Time, currency Currency) (float64, error) {
	initialValue := make(chan maAndErr)
	finalValue := make(chan maAndErr)
	go func() {
		initialValue <- capitalValueInCurrency(asset, start, currency)
	}()
	go func() {
		finalValue <- capitalValueInCurrency(asset, end, currency)
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

func capitalValueInCurrency(asset Asset, date time.Time, currency Currency) maAndErr {
	ma, err := asset.UnitPrice(date)
	if err != nil {
		return maAndErr{ma, err}
	}
	cma, err := ma.Convert(currency, date)
	return maAndErr{cma, err}
}
