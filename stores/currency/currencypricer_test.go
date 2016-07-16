package currency

import (
	"testing"
	"time"

	"github.com/vquintin/customindex/assert"
	"github.com/vquintin/customindex/assets"
	"github.com/vquintin/customindex/stores/mock"
)

func TestReturnsUnitOfCurrencyWhenAssetIsCurrency(t *testing.T) {
	mock := mock.PricerMock{Values: map[mock.AssetAndDate]assets.MoneyAmount{}}
	store := CurrencyPricer{&mock}

	actual, err := store.UnitPrice(assets.Currency("SGD"), time.Now())

	assert.AssertNil(t, "Returned an error", err)
	assert.AssertFalse(t, "The mock was called", mock.Called)
	expected := assets.MoneyAmount{Amount: 1.0, Currency: "SGD"}
	assert.AssertEquals(t, "The MoneyAmount is not as expected", expected, actual)
}

func TestCallsNextInChainWhenAssetIsNotCurrency(t *testing.T) {
	mock := mock.PricerMock{Values: map[mock.AssetAndDate]assets.MoneyAmount{}}
	store := CurrencyPricer{&mock}

	store.UnitPrice(42, time.Now())

	assert.AssertTrue(t, "The mock was not called", mock.Called)
}
