package main

import (
	"fmt"
	"strings"
	"sort"
)


type Flight struct {
	flightNumber string
	origin       string
	destination  string
	departure    string
	arrival      string
	status       string
}

var flights []Flight

func main() {
	for {
		fmt.Println("\nAirport Management System")
		fmt.Println("1. Add Flight")
		fmt.Println("2. View All Flights")
		fmt.Println("3. Search Flight by Flight Number")
		fmt.Println("4. Delete Flight")
		fmt.Println("5. Update Flight")
		fmt.Println("6. Sort Flights")
		fmt.Println("7. Filter Flights")
		fmt.Println("8. Update Flight Status")
		fmt.Println("9. Generate Reports")
		fmt.Println("10. Exit")
		fmt.Print("Enter your choice: ")

		var choice int
		if _, err := fmt.Scanln(&choice); err != nil {
			fmt.Println("Invalid input. Please enter a number.")
			continue
		}

		switch choice {
		case 1:
			addFlight()
		case 2:
			viewAllFlights()
		case 3:
			searchFlight()
		case 4:
			deleteFlight()
		case 5:
			updateFlight()
		case 6:
			sortFlights()
		case 7:
			filterFlights()
		case 8:
			updateFlightStatus()
		case 9:
			generateReports()
		case 10:
			fmt.Println("Exiting...")
			return
		default:
			fmt.Println("Invalid choice. Please try again.")
		}
	}
}