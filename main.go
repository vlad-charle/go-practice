package main

import (
	"booking-app/helper"
	"fmt"
	"time"
)

const conferenceTickets = 50         // const never change
var conferenceName = "Go Conference" // alternative var declaration
var remainingTickets uint = 50
var bookings = make([]UserData, 0) // declare a slice as a list of UserData struct with size of 0, that will auto-expand going forward

type UserData struct { // struct allows to use different data types
	firstName       string
	lastName        string
	email           string
	numberOfTickets uint
}

func main() {

	greetUsers()

	for { // infinite loop

		firstName, lastName, email, userTickets := getUserInput()
		isValidName, isValidEmail, isValidTicketNumber := helper.ValidateUserInput(firstName, lastName, email, userTickets, remainingTickets)

		if isValidName && isValidEmail && isValidTicketNumber {

			bookTicket(userTickets, firstName, lastName, email)
			go sendTickets(userTickets, firstName, lastName, email) // concurrency to escape 10 sec sleep time for main thread

			firstNames := getFirstNames()
			fmt.Printf("First name of our bookings are: %v\n", firstNames)

			if remainingTickets == 0 {
				// end program
				fmt.Println("Our conference is fully booked. Come back next year")
				break // break infinite loop
			}
		} else {
			if !isValidName {
				fmt.Println("First or last name is too short")
			}
			if !isValidEmail {
				fmt.Println("E-mail you entered doesn't containt @ sign")
			}
			if !isValidTicketNumber {
				fmt.Println("Number of tickets you entered is invalid")
			}
			continue // skip to next for loop iteration
		}
	}
}

func greetUsers() {
	fmt.Printf("Welcome to %v booking application!\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available\n", conferenceTickets, remainingTickets)
	fmt.Println("Get you tickets here to attend")
}

func getFirstNames() []string { // []string is output (return)
	firstNames := []string{}
	for _, booking := range bookings { // _ is a variable (index) that we don't want to use
		firstNames = append(firstNames, booking.firstName)
	}
	return firstNames
}

func getUserInput() (string, string, string, uint) {
	var firstName string // declare var w/o value
	var lastName string
	var email string
	var userTickets uint // uint is always > 0
	fmt.Println("Enter your first name:")
	fmt.Scan(&firstName) // user input

	fmt.Println("Enter your last name:")
	fmt.Scan(&lastName)

	fmt.Println("Enter your e-mail:")
	fmt.Scan(&email)

	fmt.Println("Enter how many tickets do you want to buy:")
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets
}

func bookTicket(userTickets uint, firstName string, lastName string, email string) {
	remainingTickets = remainingTickets - userTickets

	var userData = UserData{
		firstName:       firstName,
		lastName:        lastName,
		email:           email,
		numberOfTickets: userTickets,
	}

	bookings = append(bookings, userData)
	fmt.Printf("List of bookings is %v\n", bookings)

	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation e-mail at %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)
}

func sendTickets(userTickets uint, firstName string, lastName string, email string) {
	time.Sleep(10 * time.Second)
	var tickets = fmt.Sprintf("%v tickets for %v %v", userTickets, firstName, lastName) // Sprintf to assign printed value to var
	fmt.Println("#####################")
	fmt.Printf("Sending ticket:\n%v\nto email address %v\n", tickets, email)
	fmt.Println("#####################")
}
