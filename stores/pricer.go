package stores

import "time"
import "bitbucket.org/virgilequintin/customindex/assets"

// Pricer is an interface for a type able to give the price of an asset at a given date.
type Pricer interface {
	UnitPrice(asset interface{}, date time.Time) (assets.MoneyAmount, error)
}
