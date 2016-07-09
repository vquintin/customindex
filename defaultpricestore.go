package customindex

func NewDefaultPriceStore() PriceStore {
	store1 := FailPriceStore{}
	store2 := CurrencyPriceStore{store1}
	store3 := YahooPriceStore{&store2}
	rateStore := FixerExchangeRateStore{}
	store4 := IndexPriceStore{&store3, &store3, &rateStore}
	store4.head = store4
	return store4
}
