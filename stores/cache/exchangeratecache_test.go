package cache

import (
	"sync"
	"testing"
	"time"

	"bitbucket.org/virgilequintin/customindex/assert"
	"bitbucket.org/virgilequintin/customindex/assets"
)

type exchangeRateStoreMock struct {
	mutex sync.RWMutex
	calls uint
}

func (mock *exchangeRateStoreMock) Convert(moneyAmount assets.MoneyAmount, targetCurrency assets.Currency, date time.Time) (assets.MoneyAmount, error) {
	mock.IncrementCalls()
	return assets.MoneyAmount{moneyAmount.Amount * 2.0, targetCurrency}, nil
}

func (mock *exchangeRateStoreMock) IncrementCalls() {
	mock.mutex.Lock()
	mock.calls++
	mock.mutex.Unlock()
}

func (mock *exchangeRateStoreMock) Calls() uint {
	mock.mutex.RLock()
	defer mock.mutex.RUnlock()
	return mock.calls
}

func TestCachedExchangeRateStoreIsOnlyCalledOnce(t *testing.T) {
	mock := exchangeRateStoreMock{}
	store := NewExchangeRateCache(&mock)
	moneyAmount := assets.MoneyAmount{1.0, "EUR"}
	currency := assets.Currency("USD")
	date := time.Now()
	var n sync.WaitGroup
	for i := 0; i < 1000; i++ {
		n.Add(1)
		go func() {
			store.Convert(moneyAmount, currency, date)
			n.Done()
		}()
	}
	n.Wait()
	assert.AssertEquals(t, "The cached store was not called once", uint(1), mock.Calls())
}

func TestCacheReturnsValueFromCachedStore(t *testing.T) {
	mock := exchangeRateStoreMock{}
	store := NewExchangeRateCache(&mock)
	moneyAmount := assets.MoneyAmount{42.0, "EUR"}
	currency := assets.Currency("USD")
	date := time.Now()

	actual, err := store.Convert(moneyAmount, currency, date)

	assert.AssertNoError(t, err)
	expected := assets.MoneyAmount{84.0, "USD"}
	assert.AssertEquals(t, "The result is not as expected", expected, actual)
}

func TestCachedExchangeRateStoreIsOnlyCalledOnceWhenAmountsAreDifferent(t *testing.T) {
	mock := exchangeRateStoreMock{}
	store := NewExchangeRateCache(&mock)
	a := assets.MoneyAmount{42.0, "EUR"}
	b := assets.MoneyAmount{60.0, "EUR"}
	currency := assets.Currency("USD")
	date := time.Now()

	_, err1 := store.Convert(a, currency, date)
	actual, err2 := store.Convert(b, currency, date)

	assert.AssertNoError(t, err1)
	assert.AssertNoError(t, err2)
	expected := assets.MoneyAmount{120.0, "USD"}
	assert.AssertEquals(t, "The result is not as expected", expected, actual)
	assert.AssertEquals(t, "The cached store was not called once", uint(1), mock.Calls())
}
