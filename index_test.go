package customindex

import (
	"encoding/json"
	"testing"
	"time"
)

var index1 = Index{
	"Index 1",
	MoneyAmount{1000, "EUR"},
	time.Date(2016, 1, 1, 0, 0, 0, 0, time.UTC),
	[]Weight{
		Weight{MoneyAmount{1000, "USD"}, 1.0},
	},
}

const magic = 996.9779865952726

func TestConvertToTargetCurrencyBeforeComputePerformanceRatio(t *testing.T) {
	actual, err := index1.Value(time.Date(2016, 2, 1, 0, 0, 0, 0, time.UTC))
	if err != nil {
		t.Error(err)
	}
	expected := MoneyAmount{magic, "EUR"}
	if actual != expected {
		t.Errorf("The value of the index is not as expected. Expected: %v. Got: %v", expected, actual)
	}
}

var index2 = Index{
	"Index 2",
	MoneyAmount{1000, "EUR"},
	time.Date(2016, 1, 1, 0, 0, 0, 0, time.UTC),
	[]Weight{
		Weight{MoneyAmount{1000, "USD"}, 1.0},
		Weight{MoneyAmount{1000, "EUR"}, 1.0},
	},
}

func TestThatIndexWeightingIsNotObviouslyWrong(t *testing.T) {
	actual, err := index2.Value(time.Date(2016, 2, 1, 0, 0, 0, 0, time.UTC))
	if err != nil {
		t.Error(err)
	}
	expected := MoneyAmount{magic/2.0 + 500.0, "EUR"}
	if actual != expected {
		t.Errorf("The value of the index is not as expected. Expected: %v. Got: %v", expected, actual)
	}
}

var index3 = Index{
	"Index Perso",
	MoneyAmount{1000, "EUR"},
	time.Date(2016, 1, 1, 0, 0, 0, 0, time.UTC),
	[]Weight{
		Weight{Equity{"AAPL", "USD"}, 1.0},
		Weight{MoneyAmount{1000, "EUR"}, 1.0},
	},
}

func TestMarshallingDoesNotReturnAnError(t *testing.T) {
	bytes, err := json.Marshal(index3)
	if err != nil {
		t.Error(err)
	}
	t.Log(string(bytes))
}
