package weather_api

import (
	"fmt"
	"time"
	"strings"
	"net/http"
	"io/ioutil"
	"encoding/json"
)

type Location struct {
	Query Query `json:"query"`
}

type Query struct {
	Count int `json:"count"`
	Created time.Time `json:"created"`
	Lang string `json:"lang"`
	Results Results `json:"results"`
}

type Results struct {
	Place Place `json:"place"`
}

type Place struct {
	Woeid string `json:"woeid"`
}




func ShowWeather(loc string) string{
	locCode := GetLocation(loc)
	return GetWeather(locCode)
}

func GetLocation(loc string)   string {
	query := `q=select%20woeid%20from%20geo.places(1)%20where%20text%3D'`+ loc +`'&format=json`
	body := FetchData(query)
	var data =  Location{}

	fmt.Println(string(body))
	json.Unmarshal(body,&data)

	return data.Query.Results.Place.Woeid
}

func GetWeather(locCode string) string{
	query := `q=select%20*%20from%20weather.forecast%20where%20woeid%3D` + locCode +`&format=json`
	body := FetchData(query)

	return string(body)

}

func FetchData(query string) []byte {
	url := "https://query.yahooapis.com/v1/public/yql"

	payload := strings.NewReader(query)

	req, _ := http.NewRequest("POST", url, payload)

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
	return body

}