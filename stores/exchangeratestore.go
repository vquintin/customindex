package stores

import "time"
import "bitbucket.org/virgilequintin/customindex/assets"

type ExchangeRateStore interface {
	Convert(moneyAmount assets.MoneyAmount, targetCurrency assets.Currency, date time.Time) (assets.MoneyAmount, error)
}
