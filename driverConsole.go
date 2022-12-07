package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"
)

type Driver struct { // map this type to the record in the table
	DriverID  string
	FirstName string
	LastName  string
	IdenNo    string
	CarNo     string
	Mobile    int
	Email     string
	Password  string
}
type Drivers struct {
	Drivers map[string]Driver `json:"Drivers"`
}

func main() {
outer:
	for {
		fmt.Println(strings.Repeat("=", 10))
		fmt.Println("Driver  Console\n",
			"1. Create new driver account\n",
			"2. Update driver account\n",
			/* "3. Update passenger account\n",
			"4. Delete course\n", */
			"9. Quit")
		fmt.Print("Enter an option: ")

		var choice int
		fmt.Scanf("%d", &choice)

		switch choice {
		case 1:
			create()
		case 2:
			update()
		/* case 3:
			update()
		case 4:
			delete() */
		case 9:
			break outer
		default:
			fmt.Println("### Invalid Input ###")
		}
	}
}

/*
	 func listAll() {
		client := &http.Client{}

		if req, err := http.NewRequest(http.MethodGet, "http://localhost:5000/api/v1/courses", nil); err == nil {
			if res, err := client.Do(req); err == nil {
				if body, err := ioutil.ReadAll(res.Body); err == nil {
					var res Passengers
					json.Unmarshal(body, &res)

					for k, v := range res.Passengers {
						fmt.Println(v.Name, "(", k, ")")
						fmt.Println("Planned Intake:", v.Intake)
						fmt.Println("Min GPA:", v.MinGPA)
						fmt.Println("Max GPA:", v.MaxGPA)
						fmt.Println()
					}
				}
			} else {
				fmt.Println(2, err)
			}
		} else {
			fmt.Println(3, err)
		}
	}
*/
func create() {
	var driver Driver
	var driverID string
	driverID = "D001"

	fmt.Print("Enter your first name: ")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	driver.FirstName = strings.TrimSpace(input)

	fmt.Print("Enter your last name: ")
	fmt.Scanf("%s", &(driver.LastName))
	fmt.Print("Enter your Identification number: ")
	fmt.Scanf("%s", &(driver.IdenNo))
	fmt.Print("Enter your car Number: ")
	fmt.Scanf("%s", &(driver.CarNo))
	fmt.Print("Enter your mobile number: ")
	fmt.Scanf("%d", &(driver.Mobile))
	fmt.Print("Enter your email: ")
	fmt.Scanf("%s", &(driver.Email))
	fmt.Print("Enter your password: ")
	fmt.Scanf("%s", &(driver.Password))

	postBody, _ := json.Marshal(driver)
	resBody := bytes.NewBuffer(postBody)

	client := &http.Client{}
	if req, err := http.NewRequest(http.MethodPost, "http://localhost:5000/api/v1/courses/"+driverID, resBody); err == nil {
		if res, err := client.Do(req); err == nil {
			if res.StatusCode == 202 {
				fmt.Println("Course", driverID, "created successfully")
			} else if res.StatusCode == 409 {
				fmt.Println("Error - course", driverID, "exists")
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
	fmt.Print("Enter the ID of the account to be updated: ")
	fmt.Scanf("%v", &driverID)
	fmt.Print("Enter the new first name: ")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	driver.FirstName = strings.TrimSpace(input)
	fmt.Print("Enter your last name: ")
	fmt.Scanf("%s", &(driver.LastName))
	fmt.Print("Enter your car number: ")
	fmt.Scanf("%s", &(driver.CarNo))
	fmt.Print("Enter your mobile number: ")
	fmt.Scanf("%d", &(driver.Mobile))
	fmt.Print("Enter your email: ")
	fmt.Scanf("%s", &(driver.Email))
	fmt.Print("Enter your password: ")
	fmt.Scanf("%s", &(driver.Password))

	postBody, _ := json.Marshal(driver)

	client := &http.Client{}
	if req, err := http.NewRequest(http.MethodPut, "http://localhost:5000/api/v1/courses/"+driverID, bytes.NewBuffer(postBody)); err == nil {
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

/* func delete() {
	var course string
	fmt.Print("Enter the ID of the course to be deleted: ")
	fmt.Scanf("%v", &course)

	client := &http.Client{}
	if req, err := http.NewRequest(http.MethodDelete, "http://localhost:5000/api/v1/courses/"+course, nil); err == nil {
		if res, err := client.Do(req); err == nil {
			if res.StatusCode == 200 {
				fmt.Println("Course", course, "deleted successfully")
			} else if res.StatusCode == 404 {
				fmt.Println("Error - course", course, "does not exist")
			}
		} else {
			fmt.Println(2, err)
		}
	} else {
		fmt.Println(3, err)
	}
}
*/
