package index

import (
	"testing"
	"time"

	"bitbucket.org/virgilequintin/customindex/assert"
	"bitbucket.org/virgilequintin/customindex/assets"
	"bitbucket.org/virgilequintin/customindex/stores/mock"
)

var startDate = time.Date(2016, 1, 1, 0, 0, 0, 0, time.UTC)
var endDate = time.Date(2016, 2, 1, 0, 0, 0, 0, time.UTC)

var index1 = assets.Index{
	"Index 1",
	assets.MoneyAmount{1000, "EUR"},
	startDate,
	map[interface{}]float64{
		assets.Currency("USD"): 1.0,
	},
}

func TestIndexStoreCallsNextInChainWhenAssetIsNotIndex(t *testing.T) {
	headMock := mock.PriceStoreMock{}
	expected := assets.MoneyAmount{19.0, "USD"}
	nextMock := mock.PriceStoreMock{false, map[mock.AssetAndDate]assets.MoneyAmount{
		mock.AssetAndDate{42, startDate}: expected,
	}}
	rateStoreMock := mock.ExchangeRateStoreMock{}
	store := IndexPriceStore{&nextMock, &headMock, &rateStoreMock}

	actual, err := store.UnitPrice(42, startDate)

	assert.AssertNoError(t, err)
	assert.AssertEquals(t, "The money amount is not as expected", expected, actual)
	assert.AssertFalse(t, "The head store was called", headMock.Called)
	assert.AssertTrue(t, "The next store was not called", nextMock.Called)
}

func TestConvertToTargetCurrencyBeforeComputePerformanceRatio(t *testing.T) {
	headMock := mock.PriceStoreMock{false, map[mock.AssetAndDate]assets.MoneyAmount{
		mock.AssetAndDate{assets.Currency("USD"), startDate}: assets.MoneyAmount{1.0, "USD"},
		mock.AssetAndDate{assets.Currency("USD"), endDate}:   assets.MoneyAmount{1.0, "USD"},
	}}
	nextMock := mock.PriceStoreMock{false, map[mock.AssetAndDate]assets.MoneyAmount{}}
	rates := map[mock.CurrencyPairWithDate]float64{
		mock.CurrencyPairWithDate{startDate, "EUR", "USD"}: 1.25,
		mock.CurrencyPairWithDate{endDate, "EUR", "USD"}:   1.00,
	}
	rateStoreMock := mock.ExchangeRateStoreMock{rates}
	store := IndexPriceStore{&nextMock, &headMock, &rateStoreMock}
	actual, err := store.UnitPrice(index1, endDate)
	assert.AssertFalse(t, "The next store in chain was called", nextMock.Called)
	assert.AssertNil(t, "An error occured", err)
	expected := assets.MoneyAmount{1250, "EUR"}
	assert.AssertEquals(t, "The value of the index is not as expected", expected, actual)
}

var index2 = assets.Index{
	"Index 2",
	assets.MoneyAmount{1000, "EUR"},
	time.Date(2016, 1, 1, 0, 0, 0, 0, time.UTC),
	map[interface{}]float64{
		assets.Currency("USD"): 1.0,
		assets.Currency("EUR"): 1.0,
	},
}

func TestThatIndexWeightingIsNotObviouslyWrong(t *testing.T) {
	headMock := mock.PriceStoreMock{false, map[mock.AssetAndDate]assets.MoneyAmount{
		mock.AssetAndDate{assets.Currency("EUR"), startDate}: assets.MoneyAmount{1.0, "EUR"},
		mock.AssetAndDate{assets.Currency("EUR"), endDate}:   assets.MoneyAmount{1.0, "EUR"},
		mock.AssetAndDate{assets.Currency("USD"), startDate}: assets.MoneyAmount{1.0, "USD"},
		mock.AssetAndDate{assets.Currency("USD"), endDate}:   assets.MoneyAmount{1.0, "USD"},
	}}
	nextMock := mock.PriceStoreMock{false, map[mock.AssetAndDate]assets.MoneyAmount{}}
	rates := map[mock.CurrencyPairWithDate]float64{
		mock.CurrencyPairWithDate{startDate, "EUR", "USD"}: 1.50,
		mock.CurrencyPairWithDate{endDate, "EUR", "USD"}:   1.00,
	}
	rateStoreMock := mock.ExchangeRateStoreMock{rates}
	store := IndexPriceStore{&nextMock, &headMock, &rateStoreMock}
	actual, err := store.UnitPrice(index2, endDate)
	assert.AssertFalse(t, "The next store in chain was called", nextMock.Called)
	assert.AssertNil(t, "An error occured.", err)
	expected := assets.MoneyAmount{1250, "EUR"}
	assert.AssertEquals(t, "The value of the index is not as expected.", expected, actual)
}
