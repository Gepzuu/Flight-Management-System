package main

import (
	"fmt"
	"string"
)

type Flight struct {
	flightNumber string
	origin string
	destination string
	departure string
	arrival string
}

var flights []Flight

func main() {
	for{
		fmt.Println("Airport Management System")
		fmt.Println("1. Add Flight")
		fmt.Println("2. View Flights")
		fmt.Println("3. Exit")
		fmt.Print("Enter your choice: ")
		var choice int
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			addFlight()
		case 2:
			viewAllFlights()
		case 3:
			searchFlight()
		case 4:
			fmt.Println("Exiting...")
			return
		default:
			fmt.Println("Invalid choice. Please try again.")
		}

	}
}

func addFlight() {
	
}