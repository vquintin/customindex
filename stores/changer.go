package stores

import "time"
import "bitbucket.org/virgilequintin/customindex/assets"

// Changer is an interface for a type which is able to change a money amount to another currency.
type Changer interface {
	Change(moneyAmount assets.MoneyAmount, targetCurrency assets.Currency, date time.Time) (assets.MoneyAmount, error)
}
