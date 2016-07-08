package customindex

import "time"

type Currency string

func (c Currency) UnitPrice(date time.Time) (MoneyAmount, error) {
	return MoneyAmount{1.0, c}, nil
}
