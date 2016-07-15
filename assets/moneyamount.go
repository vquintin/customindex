package assets

import "fmt"

// MoneyAmount represents an amount of money in a currency.
type MoneyAmount struct {
	Amount   float64  `json:"amount"`
	Currency Currency `json:"currency"`
}

func (p MoneyAmount) String() string {
	return fmt.Sprintf("%v %v", p.Amount, p.Currency)
}

// Mul multiplies a MoneyAmount by a ratio.
func (p MoneyAmount) Mul(ratio float64) MoneyAmount {
	return MoneyAmount{p.Amount * ratio, p.Currency}
}

// Div divides a MoneyAmount by another one. It fails if the currencies are different.
func (p MoneyAmount) Div(other MoneyAmount) (float64, error) {
	if p.Currency == other.Currency {
		return p.Amount / other.Amount, nil
	}
	return 0.0, fmt.Errorf("Can't divide an amount in %v by an amount in %v", p.Currency, other.Currency)
}
