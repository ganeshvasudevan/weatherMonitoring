package weatherRequestHandler

import (
	"errors"
	"fmt"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	. "github.com/weatherMonitoring/userInputDataParser"
)

func TestWeatherRequestHandler(t *testing.T) {
	Convey("Check weather limits", t, func() {
		var location = LocationsType{"Test City", CoordType{20, 30}, 20, 10}
		ch := make(chan string)
		getWeatherForecast = func(location LocationsType, webResponse *WeatherQueryResponse) error {
			return errors.New("getWeatherForecast fails with exception")
		}

		Convey("Web request fails", func() {
			go FetchAndCheckWeatherLimits(location, ch)
			fmt.Println(<-ch)
		})
	})
}
