package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"time"

	"bitbucket.org/virgilequintin/customindex/assets"
	"bitbucket.org/virgilequintin/customindex/stores/factory"
)

func main() {
	var configFile = flag.String("config", "config.json", "Specifies the config file to use")
	var noCurrency = flag.Bool("nocur", false, "Don't print the currency")
	var digits = flag.Uint("digits", 3, "Number of digits after the decimal point")
	flag.Parse()
	content, err := ioutil.ReadFile(*configFile)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	var index assets.Index
	err = json.Unmarshal(content, &index)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	store := factory.NewPriceStore()
	ma, err := store.UnitPrice(index, time.Now())
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	if *noCurrency {
		format := fmt.Sprintf("%%.%vf\n", *digits)
		fmt.Printf(format, ma.Amount)
	} else {
		format := fmt.Sprintf("%%.%vf %%v\n", *digits)
		fmt.Printf(format, ma.Amount, ma.Currency)
	}
}
