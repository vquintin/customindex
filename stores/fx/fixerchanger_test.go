package fx

import (
	"encoding/json"
	"io/ioutil"
	"testing"
	"time"

	"github.com/vquintin/customindex/assert"
	"github.com/vquintin/customindex/assets"
)

func TestFixerResponseUnmarshalling(t *testing.T) {
	jsonResponse, err := ioutil.ReadFile("fixerresponse.json")
	assert.AssertNoError(t, err)

	var fxResp fixerResponse
	err = json.Unmarshal(jsonResponse, &fxResp)

	assert.AssertNoError(t, err)
	assert.AssertEquals(t, "The value for DKK is not as expected", 7.438, fxResp.Rates.DKK)
}

func TestConnectivityToFixerAPI(t *testing.T) {
	store := FixerChanger{}
	date := time.Date(2016, 07, 01, 0, 0, 0, 0, time.UTC)
	ma := assets.MoneyAmount{100, "EUR"}

	actual, err := store.Change(ma, "CZK", date)

	assert.AssertNoError(t, err)
	expected := assets.MoneyAmount{2709.3, "CZK"}
	assert.AssertEquals(t, "The conversion is not as expected", expected, actual)
}
