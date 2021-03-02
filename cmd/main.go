package main

import (
	"fmt"
	"time"

	"github.com/scosme926/weatherapp-cli/internal/controllers"
	"github.com/scosme926/weatherapp-cli/internal/models"
	"github.com/scosme926/weatherapp-cli/internal/utils"
)

func main() {
	var tsd []models.TimeSeriesDatum
	fmt.Println(tsd)
	for {
    fmt.Print("\033[H\033[2J")
    fmt.Println("###############################")
  	fmt.Println("# Weather App Command Console #")
  	fmt.Println("###############################")
    fmt.Println("Menu")
  	fmt.Println("(A) Add a record")
  	fmt.Println("(B) List all records")
  	fmt.Println("(C) Calculate Summary")
  	fmt.Println("(D) Exit")
  	fmt.Printf("\n")
  	fmt.Println("Please enter choice:")
		choice := utils.ReadConsoleString()
		switch choice {
		case "a", "A":
      fmt.Print("\033[H\033[2J")
			datum := controllers.AddInformation()
			tsd = append(tsd, datum)
      fmt.Println("###############################")
    	fmt.Println("# Weather App Command Console #")
    	fmt.Println("###############################")
      fmt.Println("Value saved.\n\nPress any key to continue.\n")
      utils.ReadConsoleString()

		case "b", "B":
			fmt.Print("\033[H\033[2J")
			controllers.AddRecordList(tsd)
      fmt.Println("\n\nPress any key to continue.\n")
      utils.ReadConsoleString()

		case "c", "C":
			fmt.Print("\033[H\033[2J")
			controllers.MakeCalculation(tsd)
      fmt.Println("\n\nPress any key to continue.\n")
      utils.ReadConsoleString()
		case "d", "D":
			fmt.Print("\033[H\033[2J")
      fmt.Println("###############################")
      fmt.Println("# Weather App Command Console #")
      fmt.Println("###############################")
			fmt.Println("Have nice day!")
      fmt.Println("\n\nPress any key to continue.\n")
      utils.ReadConsoleString()
			fmt.Print("\033[H\033[2J")
		default:
			fmt.Print("\033[H\033[2J")
			fmt.Println("Please enter a valid option. Please try again!")
			time.Sleep(3 * time.Second)
		}
		fmt.Print("\033[H\033[2J")
	}
}
