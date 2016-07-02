package customindex

import (
	"testing"
	"time"
)

func TestEquityValueReturnsFridayPriceWhenDateIsSunday(t *testing.T) {
	e := Equity{"AAPL", "USD"}
	actual, err := e.Value(time.Date(2016, 06, 26, 0, 0, 0, 0, time.UTC))
	if err != nil {
		t.Error(err)
	}
	expected := MoneyAmount{93.4, "USD"}
	if actual != expected {
		t.Errorf("Stock value is not as expected. Expected: %v. Got: %v", expected, actual)
	}
}
