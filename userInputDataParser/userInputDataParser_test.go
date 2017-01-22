package userInputDataParser

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestUserInputDataParser(t *testing.T) {
	Convey("Invalid file for userInputData parsing", t, func() {
		var userInputDataFileName string = ""
		var userInputData UserInputDataObject

		Convey("Given invalid file for reading", func() {
			So(ParseUserInputDataFile(userInputDataFileName, &userInputData), ShouldNotBeNil)
		})

		Convey("Non existent file for reading", func() {
			So(ParseUserInputDataFile("test123.json", &userInputData), ShouldNotBeNil)

		})
	})

	Convey("json format for parsing", t, func() {
		var userInputDataFileName string = "../testdata/invalidLocationWithTemperature.json"
		var userInputData UserInputDataObject

		Convey("Invalid json for parsing", func() {
			So(ParseUserInputDataFile(userInputDataFileName, &userInputData), ShouldNotBeNil)
		})

		Convey("userInputData with 0 locations for parsing", func() {
			userInputDataFileName = "../testdata/locationTemperatureWithoutLocation.json"
			err := ParseUserInputDataFile(userInputDataFileName, &userInputData)
			So(err.Error(), ShouldEqual, "No valid locations provided in the userInputData file")
		})

		Convey("userInputData with more then 60 locations for parsing", func() {
			userInputDataFileName = "../testdata/locationTemperatureWithMoreThanMaxLocations.json"
			err := ParseUserInputDataFile(userInputDataFileName, &userInputData)
			So(err.Error(), ShouldEqual, "Input data has:62 greater than:60 supported"+
				" locations, please alter the userInputData to have less locations.")
		})

		Convey("Valid json for parsing", func() {
			userInputDataFileName = "../locationWithTemperatureLimits.json"
			So(ParseUserInputDataFile(userInputDataFileName, &userInputData), ShouldBeNil)
		})
	})
}
