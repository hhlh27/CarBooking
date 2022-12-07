package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

type Booking struct { // map this type to the record in the table
	ID        string
	PassengerID string
	DriverID  string
	PickUp    string
	DropOff     string
	BookingDateTime  time.Time
	BookingStatus string
}

type Bookings struct {
	Bookings map[string]Booking `json:"Bookings"`
}

func main() {
outer:
	for {
		fmt.Println(strings.Repeat("=", 10))
		fmt.Println("Passenger Booking Management Console\n",
			"1. Make a booking\n",
			"2. View all bookings\n",
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
			View()
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

func create() {
	var booking Booking
	
	var bookingID string
	

	fmt.Println("Enter pickup postal code: ")
	fmt.Scanln( &(booking.DropOff))

	fmt.Println("Enter your drop off postal code: ")
	fmt.Scanln( &(booking.DropOff))
	
	booking.BookingDateTime=time.Now()
	booking.BookingStatus="pending"

	postBody, _ := json.Marshal(booking)
	resBody := bytes.NewBuffer(postBody)

	client := &http.Client{}
	if req, err := http.NewRequest(http.MethodPost, "http://localhost:5000/api/v1/bookings/"+bookingID, resBody); err == nil {
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
	client := &http.Client{}

	if req, err := http.NewRequest(http.MethodGet, "http://localhost:5000/api/v1/bookings", nil); err == nil {
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

