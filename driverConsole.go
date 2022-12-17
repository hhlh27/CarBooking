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

// define Driver struct
type Driver struct { // map this type to the record in the table
	FirstName string `json:"First Name"`
	LastName  string `json:"Last Name"`
	CarNo     string `json:"Car Number"`
	Mobile    string `json:"Mobile"`
	Email     string `json:"Email"`
	Password  string `json:"Password"`
}

type Drivers struct {
	Drivers map[string]Driver `json:"Drivers"`
}

func main() {
	menu() //display menu
}

// function to list menu options
func listMenu() {
	fmt.Println("1. Create New Driver Account")
	fmt.Println("2. Update Driver Account")
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

// function to create new driver account
func create() {
	var driver Driver
	var driverID string
	fmt.Println()
	//prompt user for details
	fmt.Print("Enter your identification number: ")
	fmt.Scan(&(driverID))
	fmt.Print("Enter your first name: ")
	fmt.Scan(&(driver.FirstName))
	fmt.Print("Enter your last name: ")
	fmt.Scan(&(driver.LastName))
	fmt.Print("Enter your car Number: ")
	fmt.Scan(&(driver.CarNo))
	fmt.Print("Enter your mobile number: ")
	fmt.Scan(&(driver.Mobile))
	fmt.Print("Enter your email: ")
	fmt.Scan(&(driver.Email))
	fmt.Println()
	fmt.Print("Enter your password: ")
	fmt.Scan(&(driver.Password))

	postBody, _ := json.Marshal(driver) //Sending POST Request with data
	resBody := bytes.NewBuffer(postBody)

	client := &http.Client{}
	if req, err := http.NewRequest(http.MethodPost, "http://localhost:7000/api/v1/drivers/"+driverID, resBody); err == nil {
		if res, err := client.Do(req); err == nil {
			if res.StatusCode == 202 {
				fmt.Println("driver", driverID, "created successfully")
			} else if res.StatusCode == 409 {
				fmt.Println("Error - driver", driverID, "exists")
			}
		} else {
			fmt.Println(2, err)
		}
	} else {
		fmt.Println(3, err)
	}
}

// function to update driver account
func update() {
	var driver Driver
	var driverID string
	fmt.Println()
	fmt.Print("Enter your identification number: ")
	fmt.Scan(&(driverID))
	//prompt user for updated details
	fmt.Print("Enter your first name: ")
	fmt.Scan(&(driver.FirstName))
	fmt.Print("Enter your last name: ")
	fmt.Scan(&(driver.LastName))
	fmt.Print("Enter your car Number: ")
	fmt.Scan(&(driver.CarNo))
	fmt.Print("Enter your mobile number: ")
	fmt.Scan(&(driver.Mobile))
	fmt.Print("Enter your email: ")
	fmt.Scan(&(driver.Email))
	fmt.Print("Enter your password: ")
	fmt.Scan(&(driver.Password))
	//Sending PUT Request with data
	postBody, _ := json.Marshal(driver)

	client := &http.Client{}
	if req, err := http.NewRequest(http.MethodPut, "http://localhost:7000/api/v1/drivers/"+driverID, bytes.NewBuffer(postBody)); err == nil {
		if res, err := client.Do(req); err == nil {
			if res.StatusCode == 202 {
				fmt.Println("Driver", driverID, "updated successfully")
			} else if res.StatusCode == 404 {
				fmt.Println("Error - course", driverID, "does not exist")
			}
		} else {
			fmt.Println(2, err)
		}
	} else {
		fmt.Println(3, err)
	}
}
