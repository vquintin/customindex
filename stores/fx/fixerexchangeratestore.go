package fx

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"bitbucket.org/virgilequintin/customindex/assets"
)

type rates struct {
	AUD float64
	BGN float64
	BRL float64
	CAD float64
	CHF float64
	CNY float64
	CZK float64
	DKK float64
	GBP float64
	HKD float64
	HRK float64
	HUF float64
	IDR float64
	ILS float64
	INR float64
	JPY float64
	KRW float64
	MXN float64
	MYR float64
	NOK float64
	NZD float64
	PHP float64
	PLN float64
	RON float64
	RUB float64
	SEK float64
	SGD float64
	THB float64
	TRY float64
	USD float64
	ZAR float64
}

type fixerResponse struct {
	Base  string `json:"base"`
	Date  string `json:"date"`
	Rates rates  `json:"rates"`
}

func (fr fixerResponse) getRateFromEURTo(targetCurrency assets.Currency) (float64, error) {
	if targetCurrency == "EUR" {
		return 1.0, nil
	}
	// Maybe use reflection there
	switch targetCurrency {
	case "AUD":
		return fr.Rates.AUD, nil
	case "BGN":
		return fr.Rates.BGN, nil
	case "BRL":
		return fr.Rates.BRL, nil
	case "CAD":
		return fr.Rates.CAD, nil
	case "CHF":
		return fr.Rates.CHF, nil
	case "CNY":
		return fr.Rates.CNY, nil
	case "CZK":
		return fr.Rates.CZK, nil
	case "DKK":
		return fr.Rates.DKK, nil
	case "GBP":
		return fr.Rates.GBP, nil
	case "HKD":
		return fr.Rates.HKD, nil
	case "HRK":
		return fr.Rates.HRK, nil
	case "HUF":
		return fr.Rates.HUF, nil
	case "IDR":
		return fr.Rates.IDR, nil
	case "ILS":
		return fr.Rates.ILS, nil
	case "INR":
		return fr.Rates.INR, nil
	case "JPY":
		return fr.Rates.JPY, nil
	case "KRW":
		return fr.Rates.KRW, nil
	case "MXN":
		return fr.Rates.MXN, nil
	case "MYR":
		return fr.Rates.MYR, nil
	case "NOK":
		return fr.Rates.NOK, nil
	case "NZD":
		return fr.Rates.NZD, nil
	case "PHP":
		return fr.Rates.PHP, nil
	case "PLN":
		return fr.Rates.PLN, nil
	case "RON":
		return fr.Rates.RON, nil
	case "RUB":
		return fr.Rates.RUB, nil
	case "SEK":
		return fr.Rates.SEK, nil
	case "SGD":
		return fr.Rates.SGD, nil
	case "THB":
		return fr.Rates.THB, nil
	case "TRY":
		return fr.Rates.TRY, nil
	case "USD":
		return fr.Rates.USD, nil
	case "ZAR":
		return fr.Rates.ZAR, nil
	default:
		return 0.0, fmt.Errorf("Unknown currency %v", targetCurrency)
	}
}

type FixerExchangeRateStore struct {
}

func (store FixerExchangeRateStore) Convert(moneyAmount assets.MoneyAmount, targetCurrency assets.Currency, date time.Time) (assets.MoneyAmount, error) {
	fxResp, err := store.getFixerResponse(date)
	if err != nil {
		return assets.MoneyAmount{}, err
	}
	sourceRate, err := fxResp.getRateFromEURTo(moneyAmount.Currency)
	if err != nil {
		return assets.MoneyAmount{}, err
	}
	targetRate, err := fxResp.getRateFromEURTo(targetCurrency)
	if err != nil {
		return assets.MoneyAmount{}, err
	}
	return assets.MoneyAmount{moneyAmount.Amount * targetRate / sourceRate, targetCurrency}, nil
}

const fixerURL = "http://api.fixer.io/"

func (store FixerExchangeRateStore) getFixerResponse(date time.Time) (fixerResponse, error) {
	requestURL := fixerURL + date.Format("2006-01-02")
	resp, err := http.Get(requestURL)
	if err != nil {
		return fixerResponse{}, err
	}
	defer resp.Body.Close()
	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fixerResponse{}, err
	}
	var fxResp fixerResponse
	err = json.Unmarshal(bytes, &fxResp)
	if err != nil {
		return fixerResponse{}, err
	}
	return fxResp, nil
}
