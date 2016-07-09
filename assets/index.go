package assets

import (
	"encoding/json"
	"time"
)

type Index struct {
	Name         string
	InitialValue MoneyAmount
	Creation     time.Time
	Weights      map[interface{}]float64
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
	sp := IndexSP{i.Name, i.InitialValue, i.Creation, []CurrencyWeight{}, []EquityWeight{}, []IndexWeight{}}
	for k, v := range i.Weights {
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
	index.Name = sp.Name
	index.Creation = sp.Creation
	index.InitialValue = sp.InitialValue
	index.Weights = make(map[interface{}]float64)
	for _, v := range sp.Currencies {
		index.Weights[v.Currency] = v.Weight
	}
	for _, v := range sp.Equities {
		index.Weights[v.Equity] = v.Weight
	}
	for _, v := range sp.Indexes {
		index.Weights[v.Index] = v.Weight
	}
	return nil
}
