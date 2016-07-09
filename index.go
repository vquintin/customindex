package customindex

import (
	"encoding/json"
	"time"
)

type Index struct {
	name         string
	initialValue MoneyAmount
	creation     time.Time
	weights      map[interface{}]float64
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
	index.weights = make(map[interface{}]float64)
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
