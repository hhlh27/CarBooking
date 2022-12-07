package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"strings"
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
	//now := time.Now()
outer:
	for {
		fmt.Println(strings.Repeat("=", 10))
		fmt.Println("Driver  Console\n",
			"1. Iniate start trip\n",
			"2. End trip\n",
			/* "3. Update passenger account\n",
			"4. Delete course\n", */
			"9. Quit")
		fmt.Print("Enter an option: ")

		var choice int
		fmt.Scanf("%d", &choice)

		switch choice {
		case 1:
			InitiateStartTrip()
		case 2:
			EndTrip()
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

func InitiateStartTrip() {
	var booking Booking
	var driverID string
	driverID = "D0001"

	booking.DriverID = driverID
	booking.BookingStatus = "Confimed"

	postBody, _ := json.Marshal(booking)
	resBody := bytes.NewBuffer(postBody)

	client := &http.Client{}
	if req, err := http.NewRequest(http.MethodPut, "http://localhost:5000/api/v1/bookings/"+booking.ID, resBody); err == nil {
		if res, err := client.Do(req); err == nil {
			if res.StatusCode == 202 {
				fmt.Println("Booking", driverID, "confirmed successfully")
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
	var driverID string
	var bookingID string
	driverID = "D0001"
	bookingID = "B0001"

	booking.DriverID = driverID
	booking.BookingStatus = "Completed"

	postBody, _ := json.Marshal(booking)
	resBody := bytes.NewBuffer(postBody)

	client := &http.Client{}
	if req, err := http.NewRequest(http.MethodPut, "http://localhost:5000/api/v1/bookings/"+bookingID, resBody); err == nil {
		if res, err := client.Do(req); err == nil {
			if res.StatusCode == 202 {
				fmt.Println("Booking", driverID, "completed")
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
