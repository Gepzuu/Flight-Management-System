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

func addFlight() {
	var flight Flight

	fmt.Print("Enter flight number: ")
	fmt.Scanln(&flight.flightNumber)
	if flight.flightNumber == "" {
		fmt.Println("Flight number cannot be empty.")
		return
	}

	fmt.Print("Enter origin: ")
	fmt.Scanln(&flight.origin)
	if flight.origin == "" {
		fmt.Println("Origin cannot be empty.")
		return
	}

	fmt.Print("Enter destination: ")
	fmt.Scanln(&flight.destination)
	if flight.destination == "" {
		fmt.Println("Destination cannot be empty.")
		return
	}

	fmt.Print("Enter departure time: ")
	fmt.Scanln(&flight.departure)
	if flight.departure == "" {
		fmt.Println("Departure time cannot be empty.")
		return
	}

	fmt.Print("Enter arrival time: ")
	fmt.Scanln(&flight.arrival)
	if flight.arrival == "" {
		fmt.Println("Arrival time cannot be empty.")
		return
	}

	flight.status = "on-time"
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
		fmt.Printf("Flight Number: %s, Origin: %s, Destination: %s, Departure: %s, Arrival: %s, Status: %s\n",
			flight.flightNumber, flight.origin, flight.destination, flight.departure, flight.arrival, flight.status)
	}
}

func searchFlight() {
	fmt.Print("Enter flight number to search: ")
	var flightNumber string
	fmt.Scanln(&flightNumber)

	for _, flight := range flights {
		if strings.EqualFold(flight.flightNumber, flightNumber) {
			fmt.Printf("Flight Found: Flight Number: %s, Origin: %s, Destination: %s, Departure: %s, Arrival: %s, Status: %s\n",
				flight.flightNumber, flight.origin, flight.destination, flight.departure, flight.arrival, flight.status)
			return
		}
	}

	fmt.Println("Flight not found.")
}

func deleteFlight() {
	fmt.Print("Enter flight number to delete: ")
	var flightNumber string
	fmt.Scanln(&flightNumber)

	for i, flight := range flights {
		if strings.EqualFold(flight.flightNumber, flightNumber) {
			flights = append(flights[:i], flights[i+1:]...)
			fmt.Println("Flight deleted successfully!")
			return
		}
	}

	fmt.Println("Flight not found.")
}

func updateFlight() {
	fmt.Print("Enter flight number to update: ")
	var flightNumber string
	fmt.Scanln(&flightNumber)

	for i, flight := range flights {
		if strings.EqualFold(flight.flightNumber, flightNumber) {
			fmt.Print("Enter new origin: ")
			fmt.Scanln(&flights[i].origin)

			fmt.Print("Enter new destination: ")
			fmt.Scanln(&flights[i].destination)

			fmt.Print("Enter new departure time: ")
			fmt.Scanln(&flights[i].departure)

			fmt.Print("Enter new arrival time: ")
			fmt.Scanln(&flights[i].arrival)

			fmt.Println("Flight updated successfully!")
			return
		}
	}

	fmt.Println("Flight not found.")
}

func sortFlights() {
	sort.Slice(flights, func(i, j int) bool {
		return flights[i].departure < flights[j].departure
	})

	fmt.Println("Flights sorted by departure time:")
	viewAllFlights()
}

func filterFlights() {
	fmt.Print("Enter origin to filter: ")
	var origin string
	fmt.Scanln(&origin)

	fmt.Println("Flights from", origin, ":")
	for _, flight := range flights {
		if strings.EqualFold(flight.origin, origin) {
			fmt.Printf("Flight Number: %s, Destination: %s, Departure: %s, Arrival: %s, Status: %s\n",
				flight.flightNumber, flight.destination, flight.departure, flight.arrival, flight.status)
		}
	}
}

func updateFlightStatus() {
	fmt.Print("Enter flight number to update status: ")
	var flightNumber string
	fmt.Scanln(&flightNumber)

	for i, flight := range flights {
		if strings.EqualFold(flight.flightNumber, flightNumber) {
			fmt.Print("Enter new status: ")
			fmt.Scanln(&flights[i].status)

			fmt.Println("Flight status updated successfully!")
			return
		}
	}

	fmt.Println("Flight not found.")
}

func generateReports() {
	fmt.Println("Generating reports...")
	fmt.Println("Delayed Flights:")
	for _, flight := range flights {
		if strings.EqualFold(flight.status, "delayed") {
			fmt.Printf("Flight Number: %s, Origin: %s, Destination: %s, Departure: %s, Arrival: %s, Status: %s\n",
				flight.flightNumber, flight.origin, flight.destination, flight.departure, flight.arrival, flight.status)
		}
	}

	fmt.Println("Cancelled Flights:")
	for _, flight := range flights {
		if strings.EqualFold(flight.status, "cancelled") {
			fmt.Printf("Flight Number: %s, Origin: %s, Destination: %s, Departure: %s, Arrival: %s, Status: %s\n",
				flight.flightNumber, flight.origin, flight.destination, flight.departure, flight.arrival, flight.status)
		}
	}
}