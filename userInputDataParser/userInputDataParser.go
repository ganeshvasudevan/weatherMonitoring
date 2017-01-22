package userInputDataParser

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
)

type UserInputDataObject struct {
	Locations []LocationsType
}

type LocationsType struct {
	PrettyLocalName string
	Coord           CoordType
	LowerLimit      float64
	UpperLimit      float64
}

type CoordType struct {
	Lon float64
	Lat float64
}

// ParseUserInputDataFile opens and reads the file contents and then maps the
// contents on to the given struct. Validates to make sure that there
// are some locations for which weather monitoring needs to be done.
func ParseUserInputDataFile(userInputDataFileName string, userInputData *UserInputDataObject) error {
	file, err := ioutil.ReadFile(userInputDataFileName)
	maxSupportedLocations := 60
	if err != nil {
		log.Printf("Error reading file:[%v] error: %v\n", userInputDataFileName, err.Error())
		return err
	}

	log.Printf("Successfully opened userInputData file:[%v] for reading", userInputDataFileName)

	err = json.Unmarshal(file, userInputData)
	if err != nil {
		log.Printf("Error while parsing json userInputData :%v", err.Error())
		return err
	}

	if len(userInputData.Locations) == 0 {
		log.Printf("No valid locations to monitor weather for")
		return errors.New("No valid locations provided in the userInputData file")
	}

	if len(userInputData.Locations) > maxSupportedLocations {
		log.Printf("User input data contains:%v greater than:%v locations for monitoring,"+
			"Due to weather API free usage limit of max:%v queries per minute",
			len(userInputData.Locations), maxSupportedLocations, maxSupportedLocations)

		return fmt.Errorf("Input data has:%v greater than:%v supported locations,"+
			" please alter the userInputData to have less locations.",
			len(userInputData.Locations), maxSupportedLocations)
	}

	log.Printf("Successfully parsed userInputData file:[%v] with locations to monitor:%v",
		userInputDataFileName, len(userInputData.Locations))
	return nil
}
