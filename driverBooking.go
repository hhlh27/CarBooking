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
	fmt.Println("1. View pending bookings")
	fmt.Println("2. Initiate start trip")
	fmt.Println("3. End Trip")

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
				ViewPendingBookings()
				fmt.Println("")
				bufio.NewReader(os.Stdin).ReadBytes('\n')
				listMenu()
			case 2:
				InitiateStartTrip()
				fmt.Println("")
				bufio.NewReader(os.Stdin).ReadBytes('\n')
				listMenu()
			case 3:
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
func ViewPendingBookings() {

	client := &http.Client{}

	if req, err := http.NewRequest(http.MethodGet, "http://localhost:5003/api/v1/pendingbookings", nil); err == nil {
		if res, err := client.Do(req); err == nil {
			if body, err := ioutil.ReadAll(res.Body); err == nil {
				var res Bookings
				json.Unmarshal(body, &res)

				for k, v := range res.Bookings {
					fmt.Println(v.BookingDateTime, "(", k, ")")
					fmt.Println("PickUp postal code", v.PickUp)
					fmt.Println("DropOff postal code", v.DropOff)
					fmt.Println("Booking Date Time", v.BookingDateTime)
					fmt.Println("Passenger ID", v.PassengerID)

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
func InitiateStartTrip() {
	var booking Booking
	var bookingID string
	//driverID = "D0001"
	fmt.Println()
	fmt.Print("Enter Booking ID to initiate trip: ")
	fmt.Scan(&(bookingID))
	fmt.Print("Enter Driver ID: ")
	fmt.Scan(&(booking.DriverID))
	//booking.DriverID = driverID
	booking.BookingStatus = "Confimed"

	postBody, _ := json.Marshal(booking)
	resBody := bytes.NewBuffer(postBody)

	client := &http.Client{}
	if req, err := http.NewRequest(http.MethodPut, "http://localhost:5003/api/v1/bookings/"+bookingID, resBody); err == nil {
		if res, err := client.Do(req); err == nil {
			if res.StatusCode == 202 {
				fmt.Println("Booking", bookingID, "confirmed successfully")
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
func EndTrip() {
	var booking Booking
	//var driverID string
	var bookingID string
	//driverID = "D0001"
	//bookingID = "B0001"
	fmt.Print("Enter Booking ID to end trip: ")
	fmt.Scan(&(bookingID))
	//booking.DriverID = driverID
	booking.BookingStatus = "Completed"

	postBody, _ := json.Marshal(booking)
	resBody := bytes.NewBuffer(postBody)

	client := &http.Client{}
	if req, err := http.NewRequest(http.MethodPut, "http://localhost:5003/api/v1/bookings/"+bookingID, resBody); err == nil {
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
