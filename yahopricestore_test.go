package customindex

import (
	"testing"
	"time"
)

func TestYahooStoreCallsNextInChainWhenAssetIsNotEquity(t *testing.T) {
	mock := priceStoreMock{}
	store := YahooPriceStore{&mock}

	store.UnitPrice(42, time.Now())

	assertTrue(t, "The mock was not called", mock.Called)
}

func TestYahooStoreReturnsValueFromNextInChainWhenAssetIsNotEquity(t *testing.T) {
	date := time.Date(2016, 01, 01, 0, 0, 0, 0, time.UTC)
	expected := MoneyAmount{19.0, "USD"}
	mock := priceStoreMock{false, map[assetAndDate]MoneyAmount{
		assetAndDate{42, date}: expected,
	}}
	store := YahooPriceStore{&mock}

	actual, err := store.UnitPrice(42, date)

	assertNoError(t, err)
	assertEquals(t, "yahoo store does not return value from next in chain", expected, actual)
}
