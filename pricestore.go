package customindex

import "time"

type PriceStore interface {
	UnitPrice(asset interface{}, date time.Time) (MoneyAmount, error)
}
