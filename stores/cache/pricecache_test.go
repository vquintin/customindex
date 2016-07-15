package cache

import (
	"sync"
	"testing"
	"time"

	"bitbucket.org/virgilequintin/customindex/assert"
	"bitbucket.org/virgilequintin/customindex/assets"
)

type priceStoreMock struct {
	mutex sync.RWMutex
	calls uint
}

func (mock *priceStoreMock) UnitPrice(asset interface{}, date time.Time) (assets.MoneyAmount, error) {
	mock.IncrementCalls()
	return assets.MoneyAmount{}, nil
}

func (mock *priceStoreMock) IncrementCalls() {
	mock.mutex.Lock()
	mock.calls++
	mock.mutex.Unlock()
}

func (mock *priceStoreMock) Calls() uint {
	mock.mutex.RLock()
	defer mock.mutex.RUnlock()
	return mock.calls
}

func TestCachedPriceStoreIsOnlyCalledOnce(t *testing.T) {
	mock := priceStoreMock{}
	store := NewPriceCache(&mock)
	asset := 42
	date := time.Now()
	var n sync.WaitGroup
	for i := 0; i < 1000; i++ {
		n.Add(1)
		go func() {
			store.UnitPrice(asset, date)
			n.Done()
		}()
	}
	n.Wait()
	assert.AssertEquals(t, "The cached store was not called once", uint(1), mock.Calls())
}
