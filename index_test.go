package customindex

import (
	"encoding/json"
	"testing"
	"time"
)

var index3 = Index{
	"Index Perso",
	MoneyAmount{1000, "EUR"},
	time.Date(2016, 1, 1, 0, 0, 0, 0, time.UTC),
	map[interface{}]float64{
		Equity{"AAPL", "USD"}: 1.0,
		Currency("EUR"):       1.0,
	},
}

func TestMarshallingDoesNotReturnAnError(t *testing.T) {
	bytes, err := json.Marshal(index3)
	if err != nil {
		t.Error(err)
	}
	t.Log(string(bytes))
	index := Index{}
	err = json.Unmarshal(bytes, &index)
	if err != nil {
		t.Error(err)
	}
}
