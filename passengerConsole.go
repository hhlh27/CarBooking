package main

//import required packages
import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

// define Passenger struct
type Passenger struct { // map this type to the record in the table
	FirstName string `json:"First Name"`
	LastName  string `json:"Last Name"`
	Mobile    string `json:"Mobile"`
	Email     string `json:"Email"`
	Password  string `json:"Password"`
}

type Passengers struct {
	Passengers map[string]Passenger `json:"Passengers"`
}

func main() {
	menu() //display menu
}

// function to list menu options
func listMenu() {
	fmt.Println("1. Create New Passenger Account")
	fmt.Println("2. Update Passenger Account")
	fmt.Println("0. Exit")
	fmt.Print("Enter an option: ")
}

// function to call menu functions based on user input
func menu() {
	var choose int
	listMenu()
	for {
		fmt.Scan(&choose)
		if choose == 0 {
			fmt.Println("Thank You. Exiting...")
			break
		} else {
			switch choose {
			case 1:
				create() //create new passenger account
				fmt.Println("")
				bufio.NewReader(os.Stdin).ReadBytes('\n')
				listMenu()
			case 2:
				update() //update passenger account
				fmt.Println("")
				bufio.NewReader(os.Stdin).ReadBytes('\n')
				listMenu()

			default:
				fmt.Println("Re-enter your choice!")
				listMenu()
			}
		}
	}
}

// function to create new passenger account
func create() {
	var passenger Passenger
	var passengerID string
	fmt.Println()
	//prompt user for details
	fmt.Print("Enter your identification number: ")
	fmt.Scan(&(passengerID))
	fmt.Print("Enter your first name: ")
	fmt.Scan(&(passenger.FirstName))
	fmt.Print("Enter your last name: ")
	fmt.Scan(&(passenger.LastName))
	fmt.Print("Enter your mobile number: ")
	fmt.Scan(&(passenger.Mobile))
	fmt.Print("Enter your email: ")
	fmt.Scan(&(passenger.Email))
	fmt.Print("Enter your password: ")
	fmt.Scan(&(passenger.Password))

	postBody, _ := json.Marshal(passenger) //Sending POST Request with data
	resBody := bytes.NewBuffer(postBody)

	client := &http.Client{}
	//send POST request
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

// function to update passenger account
func update() {
	var passengerID string
	var passenger Passenger
	fmt.Println()
	fmt.Print("Enter identification number: ")
	fmt.Scan(&(passengerID))
	//prompt user for updated details
	fmt.Print("Enter your first name: ")
	fmt.Scan(&(passenger.FirstName))
	fmt.Print("Enter your last name: ")
	fmt.Scan(&(passenger.LastName))
	fmt.Print("Enter your mobile number: ")
	fmt.Scan(&(passenger.Mobile))
	fmt.Print("Enter your email: ")
	fmt.Scan(&(passenger.Email))
	fmt.Print("Enter your password: ")
	fmt.Scan(&(passenger.Password))
	postBody, _ := json.Marshal(passenger)
	//Sending PUT Request with data
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
