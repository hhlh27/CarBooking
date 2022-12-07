package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"
)

type Booking struct { // map this type to the record in the table
	ID              string
	PassengerID     string
	DriverID        string
	PickUp          string
	DropOff         string
	BookingDateTime time.Time
	BookingStatus   string
}

type Bookings struct {
	Bookings map[string]Booking `json:"Bookings"`
}

func main() {
	menu()

}
func listMenu() {
	fmt.Println("1. Initiate start strip")
	fmt.Println("2. End Trip")

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
				InitiateStartTrip()
				fmt.Println("")
				bufio.NewReader(os.Stdin).ReadBytes('\n')
				listMenu()
			case 2:
				EndTrip()
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
func InitiateStartTrip() {
	var booking Booking
	var bookingID string
	//driverID = "D0001"

	//booking.DriverID = driverID
	booking.BookingStatus = "Confimed"

	postBody, _ := json.Marshal(booking)
	resBody := bytes.NewBuffer(postBody)

	client := &http.Client{}
	if req, err := http.NewRequest(http.MethodPut, "http://localhost:5000/api/v1/bookings/"+bookingID, resBody); err == nil {
		if res, err := client.Do(req); err == nil {
			if res.StatusCode == 202 {
				fmt.Println("Booking", bookingID, "confirmed successfully")
			} else if res.StatusCode == 404 {
				fmt.Println("Error - Booking", booking.ID, "does not exist")
			}
		} else {
			fmt.Println(2, err)
		}
	} else {
		fmt.Println(3, err)
	}
}
func EndTrip() {
	var booking Booking
	//var driverID string
	var bookingID string
	//driverID = "D0001"
	bookingID = "B0001"

	//booking.DriverID = driverID
	booking.BookingStatus = "Completed"

	postBody, _ := json.Marshal(booking)
	resBody := bytes.NewBuffer(postBody)

	client := &http.Client{}
	if req, err := http.NewRequest(http.MethodPut, "http://localhost:5000/api/v1/bookings/"+bookingID, resBody); err == nil {
		if res, err := client.Do(req); err == nil {
			if res.StatusCode == 202 {
				fmt.Println("Booking", bookingID, "completed")
			} else if res.StatusCode == 404 {
				fmt.Println("Error - Booking", bookingID, "does not exist")
			}
		} else {
			fmt.Println(2, err)
		}
	} else {
		fmt.Println(3, err)
	}
}
