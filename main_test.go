package main

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestWeatherMonitoring(t *testing.T) {
	Convey("Validate command line arguments", t, func() {
		var configFileName = ""
		var frequency = 0

		Convey("Missing config file name", func() {
			So(validateCommandLineArguments(configFileName, frequency), ShouldNotBeNil)
		})

		configFileName = "test.json"
		frequency = -100
		Convey("Invalid frequency should be greater than 0", func() {
			So(validateCommandLineArguments(configFileName, frequency), ShouldNotBeNil)

		})
	})
}
