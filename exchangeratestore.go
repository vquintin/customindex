package customindex

import "time"

type ExchangeRateStore interface {
	Convert(moneyAmount MoneyAmount, targetCurrency Currency, date time.Time) (MoneyAmount, error)
}
