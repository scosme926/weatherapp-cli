package controllers

import (
	"fmt"
	"time"

	"github.com/scosme926/weatherapp-cli/internal/models"
	"github.com/scosme926/weatherapp-cli/internal/utils"
)

func AddInformation() models.TimeSeriesDatum {
	fmt.Print("\033[H\033[2J")
	fmt.Println("###############################")
	fmt.Println("# Weather App Command Console #")
	fmt.Println("###############################")
	fmt.Println("Please enter the following records:")
	fmt.Println("Temperature (deg):")
	temp := utils.ReadConsoleFloat64()
	fmt.Println("Humidity (%):")
	humidity := utils.ReadConsoleFloat64()
	fmt.Println("Pressure (Pa):")
	pressure := utils.ReadConsoleFloat64()
	fmt.Println("CO2 (ppm):")
	co2 := utils.ReadConsoleFloat64()
	fmt.Println("TVOC (ppb):")
	tvoc := utils.ReadConsoleFloat64()
	fmt.Println("Date/Time:")
	timestamp := time.Now()
	fmt.Printf("\n")
	fmt.Print("\033[H\033[2J")

	time_series_datum := models.TimeSeriesDatum{
		Temperature: temp,
		Humidity:    humidity,
		Pressure:    pressure,
		Co2:         co2,
		Tvoc:        tvoc,
		Timestamp:   timestamp,
	}
	return time_series_datum
}

func AddRecordList(tsd []models.TimeSeriesDatum) {
	fmt.Println("###############################")
	fmt.Println("# Weather App Command Console #")
	fmt.Println("###############################")
	for _, v := range tsd {
		fmt.Println("")
		fmt.Printf("T:%f \n", v.Temperature)
		fmt.Printf("H:%f \n", v.Humidity)
		fmt.Printf("P:%f \n", v.Pressure)
		fmt.Printf("C:%f \n", v.Co2)
		fmt.Printf("T:%f \n", v.Tvoc)
		fmt.Println("Time:", v.Timestamp)
		fmt.Println("")
	}
}

func MakeCalculation(tsd []models.TimeSeriesDatum) {

	temperature := []float64{}
	humidity := []float64{}
	pressure := []float64{}
	co2 := []float64{}
	tvoc := []float64{}

	for _, v := range tsd {
		temperature = append(temperature, v.Temperature)
		humidity = append(humidity, v.Humidity)
		pressure = append(pressure, v.Pressure)
		co2 = append(co2, v.Co2)
		tvoc = append(tvoc, v.Tvoc)
	}
	minTemp, maxTemp, sumTemp, avgTemp := initializeGoRoutines(temperature)
	minHumid, maxHumid, sumHumid, avgHumid := initializeGoRoutines(humidity)
	minPress, maxPress, sumPress, avgPress := initializeGoRoutines(pressure)
	minCo2, maxCo2, sumCo2, avgCo2 := initializeGoRoutines(co2)
	minTvoc, maxTvoc, sumTvoc, avgTvoc := initializeGoRoutines(tvoc)
	fmt.Println("###############################")
	fmt.Println("# Weather App Command Console #")
	fmt.Println("###############################")
	getPrintStatistics("Temperature", minTemp, maxTemp, len(temperature), sumTemp, avgTemp)
	getPrintStatistics("Humidity", minHumid, maxHumid, len(humidity), sumHumid, avgHumid)
	getPrintStatistics("Pressure", minPress, maxPress, len(pressure), sumPress, avgPress)
	getPrintStatistics("CO2", minCo2, maxCo2, len(co2), sumCo2, avgCo2)
	getPrintStatistics("TVOC", minTvoc, maxTvoc, len(tvoc), sumTvoc, avgTvoc)
}

func initializeGoRoutines(arr []float64) (float64, float64, float64, float64) {
	chan1 := make(chan float64)
	chan2 := make(chan float64)
	chan3 := make(chan float64)
	chan4 := make(chan float64)
	go utils.CalculateMinimum(arr, chan1)
	go utils.CalculateMaximum(arr, chan2)
	go utils.CalculateSum(arr, chan3)
	go utils.CalculateAverage(arr, chan4)
	min, max, sum, avg := <-chan1, <-chan2, <-chan3, <-chan4
	return min, max, sum, avg
}

func getPrintStatistics(name string, min float64, max float64, count int, sum float64, avg float64) {
	fmt.Println("")
	fmt.Println("* * * * * *")
	fmt.Printf("%s\n", name)
	fmt.Println("* * * * * *")
	fmt.Println("Min: ", min)
	fmt.Println("Max: ", max)
	fmt.Println("Count: ", count)
	fmt.Println("Summation: ", sum)
	fmt.Println("Average: ", avg)
	fmt.Println("")
}
