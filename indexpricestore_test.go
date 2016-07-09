package customindex

import (
	"testing"
	"time"
)

var startDate = time.Date(2016, 1, 1, 0, 0, 0, 0, time.UTC)
var endDate = time.Date(2016, 2, 1, 0, 0, 0, 0, time.UTC)

var index1 = Index{
	"Index 1",
	MoneyAmount{1000, "EUR"},
	startDate,
	map[interface{}]float64{
		Currency("USD"): 1.0,
	},
}

func TestIndexStoreCallsNextInChainWhenAssetIsNotIndex(t *testing.T) {
	headMock := priceStoreMock{}
	expected := MoneyAmount{19.0, "USD"}
	nextMock := priceStoreMock{false, map[assetAndDate]MoneyAmount{
		assetAndDate{42, startDate}: expected,
	}}
	rateStoreMock := exchangeRateStoreMock{}
	store := IndexPriceStore{&headMock, &nextMock, &rateStoreMock}

	actual, err := store.UnitPrice(42, startDate)

	assertNoError(t, err)
	assertEquals(t, "The money amount is not as expected", expected, actual)
	assertFalse(t, "The head store was called", headMock.Called)
	assertTrue(t, "The next store was not called", nextMock.Called)
}

func TestConvertToTargetCurrencyBeforeComputePerformanceRatio(t *testing.T) {
	headMock := priceStoreMock{false, map[assetAndDate]MoneyAmount{
		assetAndDate{Currency("USD"), startDate}: MoneyAmount{1.0, "USD"},
		assetAndDate{Currency("USD"), endDate}:   MoneyAmount{1.0, "USD"},
	}}
	nextMock := priceStoreMock{false, map[assetAndDate]MoneyAmount{}}
	rates := map[currencyPairWithDate]float64{
		currencyPairWithDate{startDate, "EUR", "USD"}: 1.25,
		currencyPairWithDate{endDate, "EUR", "USD"}:   1.00,
	}
	rateStoreMock := exchangeRateStoreMock{rates}
	store := IndexPriceStore{&headMock, &nextMock, &rateStoreMock}
	actual, err := store.UnitPrice(index1, endDate)
	assertFalse(t, "The next store in chain was called", nextMock.Called)
	assertNil(t, "An error occured", err)
	expected := MoneyAmount{1250, "EUR"}
	assertEquals(t, "The value of the index is not as expected", expected, actual)
}

var index2 = Index{
	"Index 2",
	MoneyAmount{1000, "EUR"},
	time.Date(2016, 1, 1, 0, 0, 0, 0, time.UTC),
	map[interface{}]float64{
		Currency("USD"): 1.0,
		Currency("EUR"): 1.0,
	},
}

func TestThatIndexWeightingIsNotObviouslyWrong(t *testing.T) {
	headMock := priceStoreMock{false, map[assetAndDate]MoneyAmount{
		assetAndDate{Currency("EUR"), startDate}: MoneyAmount{1.0, "EUR"},
		assetAndDate{Currency("EUR"), endDate}:   MoneyAmount{1.0, "EUR"},
		assetAndDate{Currency("USD"), startDate}: MoneyAmount{1.0, "USD"},
		assetAndDate{Currency("USD"), endDate}:   MoneyAmount{1.0, "USD"},
	}}
	nextMock := priceStoreMock{false, map[assetAndDate]MoneyAmount{}}
	rates := map[currencyPairWithDate]float64{
		currencyPairWithDate{startDate, "EUR", "USD"}: 1.50,
		currencyPairWithDate{endDate, "EUR", "USD"}:   1.00,
	}
	rateStoreMock := exchangeRateStoreMock{rates}
	store := IndexPriceStore{&headMock, &nextMock, &rateStoreMock}
	actual, err := store.UnitPrice(index2, endDate)
	assertFalse(t, "The next store in chain was called", nextMock.Called)
	assertNil(t, "An error occured.", err)
	expected := MoneyAmount{1250, "EUR"}
	assertEquals(t, "The value of the index is not as expected.", expected, actual)
}
