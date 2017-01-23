package weatherRequestHandler

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"

	. "github.com/weatherMonitoring/userInputDataParser"
)

type WeatherQueryResponse struct {
	City    CityType
	Country string
	List    []ListType
}

type CityType struct {
	Id   int
	Name string
}

type ListType struct {
	Main   MainType
	Dt_txt string
}

type MainType struct {
	Temp float64
}

func createUrlFromLocation(location *LocationsType) string {
	var buffer bytes.Buffer
	buffer.WriteString("http://api.openweathermap.org/data/2.5/forecast?")
	buffer.WriteString("lat=")
	buffer.WriteString(strconv.FormatFloat(location.Coord.Lat, 'f', -1, 64))
	buffer.WriteString("&lon=")
	buffer.WriteString(strconv.FormatFloat(location.Coord.Lon, 'f', -1, 64))
	buffer.WriteString("&units=metric&appid=")
	buffer.WriteString("8dcab6031206960790cb4781ae23e92f")
	return buffer.String()
}

var getWeatherForecast = func(location LocationsType, webResponse *WeatherQueryResponse) error {
	var httpResponse []byte

	url := createUrlFromLocation(&location)

	timeout := time.Duration(5 * time.Second)
	client := http.Client{
		Timeout: timeout,
	}
	resp, err := client.Get(url)

	if err != nil {
		log.Printf("Failed to get weather from webservice, city:%v lat:%v lon:%v err:%v",
			location.PrettyLocalName, location.Coord.Lat, location.Coord.Lon, err.Error())
		return errors.New("Failed to get data from web request")
	}

	httpResponse, err = ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Printf("Failed to read http response city:%v lat:%v lon:%v err:%v",
			location.PrettyLocalName, location.Coord.Lat, location.Coord.Lon, err.Error())
		return errors.New("Failed to read web response")
	}

	err = json.Unmarshal(httpResponse, webResponse)

	if err != nil {
		log.Printf("Failed to parse http json response city:%v lat:%v lon:%v err:%v",
			location.PrettyLocalName, location.Coord.Lat, location.Coord.Lon, err.Error())
		resp.Body.Close()
		return errors.New("Failed to map response to WeatherQueryResponse struct")
	}

	/* Need to close the resp.Body or else there will be resource leak*/
	resp.Body.Close()

	return nil
}

// FetchAndCheckWeatherLimits For a given location the weather forecast is
// obtained from the url and then maps json response to struct. Then a check
// is done to see if the forecasted weather is under the given limits.
func FetchAndCheckWeatherLimits(location LocationsType, ch chan<- string) {
	var webResponse WeatherQueryResponse

	err := getWeatherForecast(location, &webResponse)

	if err != nil {
		log.Printf("Failed to get weather forecast, err:%v", err.Error())
		goto sendMessageAndReturn
	}

	log.Printf("Successfully received weather forecast data for city:%v lat:%v lon:%v",
		location.PrettyLocalName, location.Coord.Lat, location.Coord.Lon)

	for _, weatherForecast := range webResponse.List {
		if weatherForecast.Main.Temp < location.LowerLimit ||
			weatherForecast.Main.Temp > location.UpperLimit {
			log.Printf("Pretty city name:%v city name:%v lat:%v lon:%v limit upper:%v"+
				" lower:%v Forecasted Temperature:%v date:%v \n", location.PrettyLocalName,
				webResponse.City.Name, location.Coord.Lat, location.Coord.Lon,
				location.UpperLimit, location.LowerLimit, weatherForecast.Main.Temp,
				weatherForecast.Dt_txt)
		}
	}

sendMessageAndReturn:
	ch <- fmt.Sprintf("Done processing of %v", location.PrettyLocalName)
}
