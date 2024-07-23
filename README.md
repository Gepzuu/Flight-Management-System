# Flight Management System
# GoLang Flight Management System

This project is a Flight Management System implemented in GoLang. It provides various functionalities to manage flights efficiently.

## Functionalities

### Add Flight
Users can add new flights with details such as flight number, origin, destination, departure time, and arrival time. The initial status of the flight is set to "on-time".

### View All Flights
Displays a list of all added flights, including their details such as flight number, origin, destination, departure time, arrival time, and status.

### Search Flight by Flight Number
Enables users to search for a specific flight by its flight number, displaying the flight's details if found.

### Delete Flight
Allows users to delete a flight by entering its flight number. If the flight is found, it will be removed from the list.

### Update Flight
Enables users to update the details of a specific flight by entering its flight number. Users can update the origin, destination, departure time, and arrival time.

### Sort Flights
Provides the functionality to sort flights based on their departure times. The sorted list of flights is then displayed.

### Filter Flights
Allows users to filter flights based on their origin. It displays all flights that originate from the specified location.

### Update Flight Status
Enables users to update the status of a flight by entering its flight number. Possible statuses include "on-time", "delayed", and "cancelled".

### Generate Reports
Generates reports for flights based on their status. It displays lists of delayed and cancelled flights separately.

### Exit
Exits the application.

This Flight Management System is designed to be user-friendly, with a simple menu-driven interface that guides users through the available options.

## Getting Started

To run this project, follow these steps:

1. Clone the repository.
2. Install GoLang on your system.
3. Open a terminal and navigate to the project directory.
4. Run the command `go run flight.go` to start the application.

## Contributing

Contributions are welcome! If you find any issues or have suggestions for improvements, please open an issue or submit a pull request.

## License

This project is licensed under the [MIT License](https://opensource.org/licenses/MIT).

