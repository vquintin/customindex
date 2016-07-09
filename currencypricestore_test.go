package customindex

import (
	"testing"
	"time"
)

func TestReturnsUnitOfCurrencyWhenAssetIsCurrency(t *testing.T) {
	mock := priceStoreMock{false, map[assetAndDate]MoneyAmount{}}
	store := CurrencyPriceStore{&mock}

	actual, err := store.UnitPrice(Currency("SGD"), time.Now())

	assertNil(t, "Returned an error", err)
	assertFalse(t, "The mock was called", mock.Called)
	expected := MoneyAmount{1.0, "SGD"}
	assertEquals(t, "The MoneyAmount is not as expected", expected, actual)
}

func TestCallsNextInChainWhenAssetIsNotCurrency(t *testing.T) {
	mock := priceStoreMock{false, map[assetAndDate]MoneyAmount{}}
	store := CurrencyPriceStore{&mock}

	store.UnitPrice(42, time.Now())

	assertTrue(t, "The mock was not called", mock.Called)
}
