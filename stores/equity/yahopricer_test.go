package equity

import (
	"testing"
	"time"

	"github.com/vquintin/customindex/assert"
	"github.com/vquintin/customindex/assets"
	"github.com/vquintin/customindex/stores/mock"
)

func TestYahooStoreCallsNextInChainWhenAssetIsNotEquity(t *testing.T) {
	mock := mock.PricerMock{}
	store := YahooPricer{&mock}

	store.UnitPrice(42, time.Now())

	assert.AssertTrue(t, "The mock was not called", mock.Called)
}

func TestYahooStoreReturnsValueFromNextInChainWhenAssetIsNotEquity(t *testing.T) {
	date := time.Date(2016, 01, 01, 0, 0, 0, 0, time.UTC)
	expected := assets.MoneyAmount{19.0, "USD"}
	mock := mock.PricerMock{false, map[mock.AssetAndDate]assets.MoneyAmount{
		mock.AssetAndDate{42, date}: expected,
	}}
	store := YahooPricer{&mock}

	actual, err := store.UnitPrice(42, date)

	assert.AssertNoError(t, err)
	assert.AssertEquals(t, "yahoo store does not return value from next in chain", expected, actual)
}
