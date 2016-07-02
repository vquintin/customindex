package customindex

import "time"

type Capital interface {
	Value(date time.Time) (MoneyAmount, error)
}
