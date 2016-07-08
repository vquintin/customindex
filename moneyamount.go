package customindex

import (
	"fmt"
	"time"

	"github.com/openprovider/ecbrates"
)

type MoneyAmount struct {
	Amount   float64  `json:"amount"`
	Currency Currency `json:"currency"`
}

func (p MoneyAmount) String() string {
	return fmt.Sprintf("%v %v", p.Amount, p.Currency)
}

func (p MoneyAmount) Value(date time.Time) (MoneyAmount, error) {
	return p, nil
}

func (p *MoneyAmount) Convert(other Currency, date time.Time) (MoneyAmount, error) {
	if p.Currency == other {
		return *p, nil
	}
	rates, err := ecbrates.Load()
	var before ecbrates.Rates
	for _, v := range rates {
		rateDate, err := time.Parse("2006-01-02", v.Date)
		if err != nil {
			return MoneyAmount{}, err
		}
		if rateDate.Before(date) {
			before = v
			break
		}
	}
	fmt.Println(before)
	amount, err := before.Convert(p.Amount, ecbrates.Currency(p.Currency), ecbrates.Currency(other))
	return MoneyAmount{amount, other}, err
}

func (p *MoneyAmount) Mul(ratio float64) MoneyAmount {
	return MoneyAmount{p.Amount * ratio, p.Currency}
}

func (p *MoneyAmount) Div(other MoneyAmount) (float64, error) {
	if p.Currency == other.Currency {
		return p.Amount / other.Amount, nil
	} else {
		return 0.0, fmt.Errorf("Can't divide an amount in %v by an amount in %v", p.Currency, other.Currency)
	}
}
