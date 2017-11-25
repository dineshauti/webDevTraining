package main

import (
	"net/url"
	"net/http"
	"log"
	"io/ioutil"
	"fmt"
	"encoding/json"
	"flag"
	"os"
)

var (
	address *string
)

type tempObject struct {
	Cod string `json:"cod"`
	List []result `json:"list"`
}

type result struct {
	Main struct{
		Temp float64 `json:"temp"`
	} `json:"main"`
}

func init() {
	address = flag.String("city", "", "cwl <-city|--city> <cityName>")
}

func main() {

	flag.Parse()

	// Make a Flag Required
	if *address == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}

	encodedAddress := url.QueryEscape(*address)
	url_geocode := "http://api.openweathermap.org/data/2.5/forecast?q="+encodedAddress+"&APPID=314f5902c8925e7ad8c3cd75325650e7"

	res, err := http.Get(url_geocode)
	if err != nil {
		log.Fatalln(err)
	}

	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}

	temperature := tempObject{}
	jsonErr := json.Unmarshal(body, &temperature)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	temp := temperature.List[0].Main.Temp -  273.15

	fmt.Printf("The temperature at %s is %4.2f celcius. \n",*address,temp)

}
