package customindex

import (
	"encoding/json"
	"io/ioutil"
	"testing"
	"time"
)

func TestFixerResponseUnmarshalling(t *testing.T) {
	jsonResponse, err := ioutil.ReadFile("fixerresponse.json")
	assertNoError(t, err)

	var fxResp fixerResponse
	err = json.Unmarshal(jsonResponse, &fxResp)

	assertNoError(t, err)
	assertEquals(t, "The value for DKK is not as expected", 7.438, fxResp.Rates.DKK)
}

func TestConnectivityToFixerAPI(t *testing.T) {
	store := FixerExchangeRateStore{}
	date := time.Date(2016, 07, 01, 0, 0, 0, 0, time.UTC)
	ma := MoneyAmount{100, "EUR"}

	actual, err := store.Convert(ma, "CZK", date)

	assertNoError(t, err)
	expected := MoneyAmount{2709.3, "CZK"}
	assertEquals(t, "The conversion is not as expected", expected, actual)
}
