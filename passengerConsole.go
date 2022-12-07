package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

type Passenger struct { // map this type to the record in the table
	ID        string
	FirstName string
	LastName  string
	Mobile    string
	Email     string
	Password  string
}

type Passengers struct {
	Passengers map[string]Passenger `json:"Passengers"`
}

func main() {
	//outer:
	for {
		fmt.Println(strings.Repeat("=", 10))
		fmt.Println("Passenger Console\n",
			"1. Create new passenger account\n",
			"2. Update passenger account\n",
			/* "3. Update passenger account\n",
			"4. Delete course\n", */
			"9. Quit")
		fmt.Print("Enter an option: ")

		var choice int
		fmt.Scanf("%d", &choice)

		switch choice {
		case 1:
			create()
			//break
		case 2:
			update()
		/* case 3:
			update()
		case 4:
			delete() */
		case 9:
			//break outer
			break
		default:
			fmt.Println("### Invalid Input ###")
		}
	}
}

func create() {

	var passengerID string
	passengerID = "P0002"
	/* fmt.Println("Enter Your First Name: ")

	// var then variable name then variable type
	var first string

	// Taking input from user
	fmt.Scanln(&first)
	fmt.Println("Enter Second Last Name: ")
	var second string
	fmt.Scanln(&second)

	// Print function is used to
	// display output in the same line
	fmt.Print("Your Full Name is: ")

	// Addition of two string
	fmt.Print(first + " " + second) */
	var passenger Passenger
	fmt.Println("Enter your first name: ")

	/* reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	passenger.FirstName = strings.TrimSpace(input) */
	fmt.Scanln(&(passenger.FirstName))
	fmt.Println("Enter your last name: ")
	fmt.Scanln(&(passenger.LastName))
	//fmt.Scanf("%s", &(passenger.LastName))
	fmt.Println("Enter your mobile number: ")
	fmt.Scanln(&(passenger.Mobile))
	fmt.Println("Enter your email: ")
	fmt.Scanln(&(passenger.Email))
	fmt.Println("Enter your password: ")
	fmt.Scanln(&(passenger.Password))

	postBody, _ := json.Marshal(passenger)
	resBody := bytes.NewBuffer(postBody)

	client := &http.Client{}
	if req, err := http.NewRequest(http.MethodPost, "http://localhost:5000/api/v1/passengers/"+passengerID, resBody); err == nil {
		if res, err := client.Do(req); err == nil {
			if res.StatusCode == 202 {
				fmt.Println("Passenger", passengerID, "created successfully")
			} else if res.StatusCode == 409 {
				fmt.Println("Error - passenger", passengerID, "exists")
			}
		} else {
			fmt.Println(2, err)
		}
	} else {
		fmt.Println(3, err)
	}
}
func update() {

	var passenger Passenger
	fmt.Println("Enter the ID of the account to be updated: ")
	fmt.Scanln(&(passenger.ID))
	if passenger.ID != " " {
		updateDetails(passenger.ID)
	}

}
func updateDetails(passengerID string) {
	var passenger Passenger
	fmt.Println("Enter your first name: ")

	/* reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	passenger.FirstName = strings.TrimSpace(input) */
	fmt.Scanln(&(passenger.FirstName))
	fmt.Println("Enter your last name: ")
	fmt.Scanln(&(passenger.LastName))
	//fmt.Scanf("%s", &(passenger.LastName))
	fmt.Println("Enter your mobile number: ")
	fmt.Scanln(&(passenger.Mobile))
	fmt.Println("Enter your email: ")
	fmt.Scanln(&(passenger.Email))
	fmt.Println("Enter your password: ")
	fmt.Scanln(&(passenger.Password))
	/* fmt.Print("Enter the new first name: ")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	passenger.FirstName = strings.TrimSpace(input)
	fmt.Print("Enter your last name: ")
	fmt.Scanf("%s", &(passenger.LastName))
	fmt.Print("Enter your mobile number: ")
	fmt.Scanf("%d", &(passenger.Mobile))
	fmt.Print("Enter your email: ")
	fmt.Scanf("%s", &(passenger.Email))
	fmt.Print("Enter your password: ")
	fmt.Scanf("%s", &(passenger.Password)) */

	postBody, _ := json.Marshal(passenger)

	client := &http.Client{}
	if req, err := http.NewRequest(http.MethodPut, "http://localhost:5000/api/v1/passengers/"+passengerID, bytes.NewBuffer(postBody)); err == nil {
		if res, err := client.Do(req); err == nil {
			if res.StatusCode == 202 {
				fmt.Println("passenger", passengerID, "updated successfully")
			} else if res.StatusCode == 404 {
				fmt.Println("Error - passenger", passengerID, "does not exist")
			}
		} else {
			fmt.Println(2, err)
		}
	} else {
		fmt.Println(3, err)
	}
}
