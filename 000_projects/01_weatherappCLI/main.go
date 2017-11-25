package main

import (
	"net/http"
	"log"
	"io/ioutil"
	"net/url"
	"encoding/json"
	"fmt"
	"strconv"
	"flag"
)

var (
	address *string
)

type geocode struct {
	Results []result `json:"results"`
	Status string `json:"status"`
}

type result struct {
	Geometry struct{
		Location struct{
			Lat float64 `json:lat`
			Lng float64 `json:lng`
		} `json:location`
	} `json:geometry`
}

type weather struct {
	Currently struct{
		Summary string `json:"summary"`
		Temperature float32 `json:"temperature"`
	} `json:currently`
}

func init() {
	address = flag.String("a", "Charlotte", "Enter the address")
}

func main() {


	flag.Parse()

	encodedAddress := url.QueryEscape(*address)
	url_geocode := "https://maps.googleapis.com/maps/api/geocode/json?address="+encodedAddress+"&key=AIzaSyDOBkrEixE2UY10ww2Zsrod3ZHUP5qxOfs"

	res, err := http.Get(url_geocode)
	if err != nil {
		log.Fatalln(err)
	}

	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}

	gc := geocode{}
	jsonErr := json.Unmarshal(body, &gc)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	location := gc.Results[0].Geometry.Location

	url_weather := "https://api.darksky.net/forecast/6e204a2cf67e1ae67497355898a6cb0b/"+strconv.FormatFloat(location.Lat,'f', -1, 64)+","+strconv.FormatFloat(location.Lng, 'f', -1, 64)

	res, err = http.Get(url_weather)
	if err != nil {
		log.Fatalln(err)
	}


	body, readErr = ioutil.ReadAll(res.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}

	w := weather{}
	jsonErr = json.Unmarshal(body, &w)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	fmt.Printf("The temperature is %4.2f Fahrenheit \n",w.Currently.Temperature)

}
