package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"time"

	"bitbucket.org/virgilequintin/customindex"
)

func main() {
	var configFile = flag.String("config", "config.json", "Specifies the config file to use")
	content, err := ioutil.ReadFile(*configFile)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	var index customindex.Index
	err = json.Unmarshal(content, &index)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	ma, err := index.Value(time.Now())
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	fmt.Println(ma)
}
