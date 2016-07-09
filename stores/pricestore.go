package stores

import "time"
import "bitbucket.org/virgilequintin/customindex/assets"

type PriceStore interface {
	UnitPrice(asset interface{}, date time.Time) (assets.MoneyAmount, error)
}
