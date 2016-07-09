package equity

import (
	"testing"
	"time"

	"bitbucket.org/virgilequintin/customindex/assert"
	"bitbucket.org/virgilequintin/customindex/assets"
	"bitbucket.org/virgilequintin/customindex/stores/mock"
)

func TestYahooStoreCallsNextInChainWhenAssetIsNotEquity(t *testing.T) {
	mock := mock.PriceStoreMock{}
	store := YahooPriceStore{&mock}

	store.UnitPrice(42, time.Now())

	assert.AssertTrue(t, "The mock was not called", mock.Called)
}

func TestYahooStoreReturnsValueFromNextInChainWhenAssetIsNotEquity(t *testing.T) {
	date := time.Date(2016, 01, 01, 0, 0, 0, 0, time.UTC)
	expected := assets.MoneyAmount{19.0, "USD"}
	mock := mock.PriceStoreMock{false, map[mock.AssetAndDate]assets.MoneyAmount{
		mock.AssetAndDate{42, date}: expected,
	}}
	store := YahooPriceStore{&mock}

	actual, err := store.UnitPrice(42, date)

	assert.AssertNoError(t, err)
	assert.AssertEquals(t, "yahoo store does not return value from next in chain", expected, actual)
}
