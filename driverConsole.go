package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

type Driver struct { // map this type to the record in the table
	DriverID  string
	FirstName string
	LastName  string
	CarNo     string
	Mobile    int
	Email     string
	Password  string
}
type Drivers struct {
	Drivers map[string]Driver `json:"Drivers"`
}

func main() {
	menu()
}

func listMenu() {
	fmt.Println("1. Create New Driver Account")
	fmt.Println("2. Update Driver Account")
	fmt.Println("0. Exit")
	fmt.Print("Enter an option: ")
}
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
				create()
				fmt.Println("")
				bufio.NewReader(os.Stdin).ReadBytes('\n')
				listMenu()
			case 2:
				update()
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
func create() {
	var driver Driver
	var driverID string

	//consoleReader := bufio.NewReader(os.Stdin)
	fmt.Println()
	fmt.Print("Enter your identification number: ")

	fmt.Scan(&(driver.DriverID))
	fmt.Print("Enter your first name: ")

	fmt.Scan(&(driver.FirstName))
	/* reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	driver.FirstName = strings.TrimSpace(input)   */
	//fmt.Scanf("%s", &(driver.FirstName))
	//fmt.Print("\n")
	fmt.Print("Enter your last name: ")

	fmt.Scan(&(driver.LastName))

	fmt.Print("Enter your car Number: ")
	fmt.Scan(&(driver.CarNo))
	//fmt.Print("\n")
	fmt.Print("Enter your mobile number: ")
	fmt.Scan(&(driver.Mobile))
	//fmt.Print("\n")
	fmt.Print("Enter your email: ")
	fmt.Scan(&(driver.Email))
	fmt.Println()
	//fmt.Print("\n")
	fmt.Print("Enter your password: ")
	fmt.Scan(&(driver.Password))
	//fmt.Print("\n")
	postBody, _ := json.Marshal(driver)
	resBody := bytes.NewBuffer(postBody)

	client := &http.Client{}
	if req, err := http.NewRequest(http.MethodPost, "http://localhost:5000/api/v1/drivers/"+driverID, resBody); err == nil {
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
func update() {
	var driver Driver
	var driverID string

	fmt.Println()
	fmt.Print("Enter your identification number: ")

	fmt.Scan(&(driver.DriverID))

	fmt.Print("Enter your first name: ")

	fmt.Scan(&(driver.FirstName))
	/* reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	driver.FirstName = strings.TrimSpace(input)   */
	//fmt.Scanf("%s", &(driver.FirstName))
	//fmt.Print("\n")
	fmt.Print("Enter your last name: ")

	fmt.Scan(&(driver.LastName))
	//fmt.Print("\n")

	//fmt.Print("\n")
	fmt.Print("Enter your car Number: ")
	fmt.Scan(&(driver.CarNo))
	//fmt.Print("\n")
	fmt.Print("Enter your mobile number: ")
	fmt.Scan(&(driver.Mobile))
	//fmt.Print("\n")
	fmt.Print("Enter your email: ")
	fmt.Scan(&(driver.Email))

	//fmt.Println()
	//fmt.Print("\n")
	fmt.Print("Enter your password: ")
	fmt.Scan(&(driver.Password))

	postBody, _ := json.Marshal(driver)

	client := &http.Client{}
	if req, err := http.NewRequest(http.MethodPut, "http://localhost:5000/api/v1/drivers/"+driverID, bytes.NewBuffer(postBody)); err == nil {
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
