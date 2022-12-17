package main

//import required packages
import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

// define Booking struct
type Booking struct { // map this type to the record in the table
	PassengerID     string `json:"Passenger Id"`
	DriverID        string `json:"Driver Id"`
	PickUp          string `json:"Pickup"`
	DropOff         string `json:"Dropoff"`
	BookingDateTime string `json:"Booking DateTime"`
	BookingStatus   string `json:"Booking Status"`
}

type Bookings struct {
	Bookings map[string]Booking `json:"Bookings"`
}

func main() {
	menu() //display menu
}

// function to list menu options
func listMenu() {
	fmt.Println("1. Make a booking")
	fmt.Println("2. View all bookings")
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
				create() //make a new booking
				fmt.Println("")
				bufio.NewReader(os.Stdin).ReadBytes('\n')
				listMenu()
			case 2:
				View() //view all bookings
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
//function to make a new booking 
func create() {
	var booking Booking
	var bookingID string
	//prompt user for details
	fmt.Println()
	fmt.Print("Enter Passenger Id: ")
	fmt.Scan(&(booking.DropOff))
	booking.DriverID = ""
	fmt.Print("Enter pickup postal code: ")
	fmt.Scan(&(booking.DropOff))
	fmt.Print("Enter your drop off postal code: ")
	fmt.Scan(&(booking.DropOff))
	fmt.Print("Enter booking date time: ")
	fmt.Scan(&(booking.BookingDateTime))
	booking.BookingStatus = "pending"

	postBody, _ := json.Marshal(booking) //Sending POST Request with data
	resBody := bytes.NewBuffer(postBody)

	client := &http.Client{}
	//send POST request
	if req, err := http.NewRequest(http.MethodPost, "http://localhost:5003/api/v1/bookings/"+bookingID, resBody); err == nil {
		if res, err := client.Do(req); err == nil {
			if res.StatusCode == 202 {
				fmt.Println("Booking", bookingID, "created successfully")
			} else if res.StatusCode == 409 {
				fmt.Println("Error - booking", bookingID, "exists")
			}
		} else {
			fmt.Println(2, err)
		}
	} else {
		fmt.Println(3, err)
	}
}
//function to view bookings
func View() {
	var booking Booking
	//prompt user for id
	fmt.Print("Enter your identification number: ")
	fmt.Scan(&booking.PassengerID)
	client := &http.Client{}
	//send GET request
	if req, err := http.NewRequest(http.MethodGet, "http://localhost:5003/api/v1/bookings?passengerid"+booking.PassengerID, nil); err == nil {
		if res, err := client.Do(req); err == nil {
			if body, err := ioutil.ReadAll(res.Body); err == nil {
				var res Bookings
				json.Unmarshal(body, &res) //Decoding from JSON
				for k, v := range res.Bookings {
					//print booking details
					fmt.Println(v.BookingDateTime, "(", k, ")")
					fmt.Println("PickUp postal code", v.PickUp)
					fmt.Println("DropOff postal code", v.DropOff)
					fmt.Println("Driver ID", v.DriverID)

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
