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
	var flight Flight

	fmt.Print("Enter flight number: ")
	fmt.Scanln(&flight.flightNumber)

	fmt.Print("Enter origin: ")
	fmt.Scanln(&flight.origin)

	fmt.Print("Enter destination: ")
	fmt.Scanln(&flight.destination)

	fmt.Print("Enter departure time: ")
	fmt.Scanln(&flight.departure)

	fmt.Print("Enter arrival time: ")
	fmt.Scanln(&flight.arrival)

	flights = append(flights, flight)
	fmt.Println("Flight added successfully!")
}

func viewAllFlights() {
	if len(flights) == 0 {
		fmt.Println("No flights available.")
		return
	}

	fmt.Println("All Flights:")
	for _, flight := range flights {
		fmt.Printf("Flight Number: %s, Origin: %s, Destination: %s, Departure: %s, Arrival: %s\n",
			flight.flightNumber, flight.origin, flight.destination, flight.departure, flight.arrival)
	}
}

func searchFlight() {
	
}