package main

import (
	"errors"
	"flag"
	"fmt"
	"log"
	"os"
	"time"

	. "github.com/weatherMonitoring/userInputDataParser"
	. "github.com/weatherMonitoring/weatherRequestHandler"
)

func startWeatherMonitoring(userInputData *UserInputDataObject, ch chan<- string) {
	for _, location := range userInputData.Locations {
		go FetchAndCheckWeatherLimits(location, ch)
	}
}

func startPeriodicWeatherChecks(userInputData *UserInputDataObject, ch chan<- string, frequency int) {
	t := time.NewTicker(time.Duration(frequency) * time.Minute)
	for {
		startWeatherMonitoring(userInputData, ch)
		<-t.C
	}
}

func handleCommandLineArgs() (string, int) {
	dataInputFileName := flag.String("locationAndTemperatureLimit", "", "Name of the input file"+
		"with locations and temperature limit")
	frequency := flag.Int("frequency", 0, "frequency of weather checks in minutes")

	flag.Parse()

	return *dataInputFileName, *frequency
}

func validateCommandLineArguments(dataInputFileName string, frequency int) error {
	if len(dataInputFileName) == 0 {
		log.Printf("Data input filename is missing in arguments")
		return errors.New("Data input file name is missing")
	}

	if frequency < 0 {
		log.Printf("Invalid weather check frequency:%v", frequency)
		return errors.New("Invalid weather check frequency configuration")
	}

	return nil
}

func main() {
	var logFileName = "logs"
	logFileHandle, err := os.OpenFile(logFileName, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalln("Failed to open log file:%v error:%v", logFileName, err.Error())
		os.Exit(2)
	}

	defer logFileHandle.Close()

	log.SetOutput(logFileHandle)

	if err != nil {
		fmt.Printf("Couldnt open log file error:%v", err.Error())
		os.Exit(2)
	}

	fmt.Printf("Successfully opened logfile:[%v] \n", logFileName)

	dataInputFileName, frequency := handleCommandLineArgs()

	err = validateCommandLineArguments(dataInputFileName, frequency)
	if err != nil {
		fmt.Printf("Validation of command line args failed, error:%v \n", err.Error())
		log.Printf("Validation of command line args failed, error:%v \n", err.Error())
		os.Exit(2)
	}

	var userInputData UserInputDataObject
	err = ParseUserInputDataFile(dataInputFileName, &userInputData)

	if err != nil {
		fmt.Printf("Data input file parsing failed error:%v\n", err.Error())
		log.Printf("Data input parsing failed error:%v\n", err.Error())
		os.Exit(2)
	}

	ch := make(chan string)

	if frequency > 0 {
		log.Printf("Starting periodic weather checks for the scheduled frequency:%v\n", frequency)
		startPeriodicWeatherChecks(&userInputData, ch, frequency)
	}

	startWeatherMonitoring(&userInputData, ch)
	log.Printf("Frequency is 0, No periodic weather checks done")

	for _ = range userInputData.Locations {
		log.Println(<-ch)
	}
}
