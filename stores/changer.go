package stores

import (
	"time"

	"github.com/vquintin/customindex/assets"
)

// Changer is an interface for a type which is able to change a money amount to another currency.
type Changer interface {
	Change(moneyAmount assets.MoneyAmount, targetCurrency assets.Currency, date time.Time) (assets.MoneyAmount, error)
}
