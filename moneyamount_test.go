package customindex

import (
	"testing"
	"time"
)

func TestThatValueIsConstant(t *testing.T) {
	ma := MoneyAmount{1000.0, "EUR"}
	actual, _ := ma.Value(time.Date(2030, 1, 1, 0, 0, 0, 0, time.UTC))
	expected := ma
	if actual != expected {
		t.Errorf("The money amount is not as expected. Got %v. Expected %v", actual, expected)
	}
}

func TestConvertConnectivity(t *testing.T) {
	ma := MoneyAmount{1000.0, "EUR"}
	actual, _ := ma.Convert("USD", time.Date(2015, 12, 01, 0, 0, 0, 0, time.UTC))
	expected := MoneyAmount{1057.9, "USD"}
	if actual != expected {
		t.Errorf("The dollar amount is not as expected. Got %v. Expected %v", actual, expected)
	}
}
