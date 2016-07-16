package index

import (
	"testing"
	"time"

	"github.com/vquintin/customindex/assert"
	"github.com/vquintin/customindex/assets"
	"github.com/vquintin/customindex/stores/mock"
)

var startDate = time.Date(2016, 1, 1, 0, 0, 0, 0, time.UTC)
var endDate = time.Date(2016, 2, 1, 0, 0, 0, 0, time.UTC)

var index1 = assets.Index{
	Name:         "Index 1",
	InitialValue: assets.MoneyAmount{Amount: 1000, Currency: "EUR"},
	Creation:     startDate,
	Weights: map[interface{}]float64{
		assets.Currency("USD"): 1.0,
	},
}

func TestIndexStoreCallsNextInChainWhenAssetIsNotIndex(t *testing.T) {
	headMock := mock.PricerMock{}
	expected := assets.MoneyAmount{Amount: 19.0, Currency: "USD"}
	nextMock := mock.PricerMock{Values: map[mock.AssetAndDate]assets.MoneyAmount{
		mock.AssetAndDate{Asset: 42, Date: startDate}: expected,
	}}
	rateStoreMock := mock.ChangerMock{}
	store := IndexPricer{&nextMock, &headMock, &rateStoreMock}

	actual, err := store.UnitPrice(42, startDate)

	assert.AssertNoError(t, err)
	assert.AssertEquals(t, "The money amount is not as expected", expected, actual)
	assert.AssertFalse(t, "The head store was called", headMock.Called)
	assert.AssertTrue(t, "The next store was not called", nextMock.Called)
}

func TestConvertToTargetCurrencyBeforeComputePerformanceRatio(t *testing.T) {
	headMock := mock.PricerMock{Values: map[mock.AssetAndDate]assets.MoneyAmount{
		mock.AssetAndDate{Asset: assets.Currency("USD"), Date: startDate}: assets.MoneyAmount{Amount: 1.0, Currency: "USD"},
		mock.AssetAndDate{Asset: assets.Currency("USD"), Date: endDate}:   assets.MoneyAmount{Amount: 1.0, Currency: "USD"},
	}}
	nextMock := mock.PricerMock{Values: map[mock.AssetAndDate]assets.MoneyAmount{}}
	rates := map[mock.CurrencyPairWithDate]float64{
		mock.CurrencyPairWithDate{Date: startDate, L: "EUR", R: "USD"}: 1.25,
		mock.CurrencyPairWithDate{Date: endDate, L: "EUR", R: "USD"}:   1.00,
	}
	rateStoreMock := mock.ChangerMock{Rates: rates}
	store := IndexPricer{&nextMock, &headMock, &rateStoreMock}
	actual, err := store.UnitPrice(index1, endDate)
	assert.AssertFalse(t, "The next store in chain was called", nextMock.Called)
	assert.AssertNil(t, "An error occured", err)
	expected := assets.MoneyAmount{Amount: 1250, Currency: "EUR"}
	assert.AssertEquals(t, "The value of the index is not as expected", expected, actual)
}

var index2 = assets.Index{
	Name:         "Index 2",
	InitialValue: assets.MoneyAmount{Amount: 1000, Currency: "EUR"},
	Creation:     time.Date(2016, 1, 1, 0, 0, 0, 0, time.UTC),
	Weights: map[interface{}]float64{
		assets.Currency("USD"): 1.0,
		assets.Currency("EUR"): 1.0,
	},
}

func TestThatIndexWeightingIsNotObviouslyWrong(t *testing.T) {
	headMock := mock.PricerMock{Values: map[mock.AssetAndDate]assets.MoneyAmount{
		mock.AssetAndDate{Asset: assets.Currency("EUR"), Date: startDate}: assets.MoneyAmount{Amount: 1.0, Currency: "EUR"},
		mock.AssetAndDate{Asset: assets.Currency("EUR"), Date: endDate}:   assets.MoneyAmount{Amount: 1.0, Currency: "EUR"},
		mock.AssetAndDate{Asset: assets.Currency("USD"), Date: startDate}: assets.MoneyAmount{Amount: 1.0, Currency: "USD"},
		mock.AssetAndDate{Asset: assets.Currency("USD"), Date: endDate}:   assets.MoneyAmount{Amount: 1.0, Currency: "USD"},
	}}
	nextMock := mock.PricerMock{Values: map[mock.AssetAndDate]assets.MoneyAmount{}}
	rates := map[mock.CurrencyPairWithDate]float64{
		mock.CurrencyPairWithDate{Date: startDate, L: "EUR", R: "USD"}: 1.50,
		mock.CurrencyPairWithDate{Date: endDate, L: "EUR", R: "USD"}:   1.00,
	}
	rateStoreMock := mock.ChangerMock{Rates: rates}
	store := IndexPricer{&nextMock, &headMock, &rateStoreMock}
	actual, err := store.UnitPrice(index2, endDate)
	assert.AssertFalse(t, "The next store in chain was called", nextMock.Called)
	assert.AssertNil(t, "An error occured.", err)
	expected := assets.MoneyAmount{Amount: 1250, Currency: "EUR"}
	assert.AssertEquals(t, "The value of the index is not as expected.", expected, actual)
}
