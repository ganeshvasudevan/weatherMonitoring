# Project Weather Monitoring

Periodically checks weather forecast against given limits of temperature and reports if there are forecasted temperature outside the limit.
The weather forecast is retrieved using API(s) provided by https://openweathermap.org/api

## Getting Started

These instructions will get you a copy of the project up and running on your local machine for development and testing purposes.

### Prerequisites

After you clone the repository you can build a docker image with the given docker file.
Run docker with the newly created image.

```
docker build -t weatherMonitoring .
docker run -i -t weatherMonitoring
```

## Program execution
Once the console is ready the program can be executed with the following

```
./weatherMonitoring -locationAndTemperatureLimit <location and temperature input file> -frequency <time in minutes> 
ex: ./weatherMonitoring -locationAndTemperatureLimit locationWithTemperatureLimits.json -frequency 30
```

The user input data file format is json with the below template
```

{
    "locations": [
        {
            "prettyLocalName": "London",
            "coord": {
                "lat": 51.5074,
                "lon": -0.1278
            },
            "lowerLimit": 2,
            "upperLimit": 10
        },
        {
            "prettyLocalName": "New York",
            "coord": {
                "lat": 40.7128,
                "lon": -74.0059
            },
            "lowerLimit": 2,
            "upperLimit": 10
        }
    ]
}

```
A sample input data file is placed with the this repo for execution.

Due to the weather API used in this program currently max supported location is 60.

## Running the tests

Once the docker is launched the Unit tests can be executed from source folder with

```
goconvey . or go test -v
```
## Authors

* **Ganesh Vasudevan**


## License

This project is licensed under the MIT License - see the [LICENSE.md](LICENSE.md) file for details
