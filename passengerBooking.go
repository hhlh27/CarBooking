package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

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
	menu()

}
func listMenu() {
	fmt.Println("1. Make a booking")
	fmt.Println("2. View all bookings")

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
				View()
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
	var booking Booking

	var bookingID string

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
	//booking.BookingDateTime = time.Now()
	booking.BookingStatus = "pending"

	postBody, _ := json.Marshal(booking)
	resBody := bytes.NewBuffer(postBody)

	client := &http.Client{}
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
func View() {
	//var passengerID string
	var booking Booking
	fmt.Print("Enter your identification number: ")
	fmt.Scan(&booking.PassengerID)
	client := &http.Client{}

	if req, err := http.NewRequest(http.MethodGet, "http://localhost:5003/api/v1/passengerbookings", nil); err == nil {
		if res, err := client.Do(req); err == nil {
			if body, err := ioutil.ReadAll(res.Body); err == nil {
				var res Bookings
				json.Unmarshal(body, &res)

				for k, v := range res.Bookings {
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
