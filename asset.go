package customindex

import "time"

type Asset interface {
	UnitPrice(date time.Time) (MoneyAmount, error)
}
