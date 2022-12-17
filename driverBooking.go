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
	menu()//display menu
}

// function to list menu options
func listMenu() {
	fmt.Println("1. View pending bookings")
	fmt.Println("2. Initiate start trip")
	fmt.Println("3. End Trip")
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
				ViewPendingBookings()//View pending bookings
				fmt.Println("")
				bufio.NewReader(os.Stdin).ReadBytes('\n')
				listMenu()
			case 2:
				InitiateStartTrip()//Initiate a start trip for pending booking
				fmt.Println("")
				bufio.NewReader(os.Stdin).ReadBytes('\n')
				listMenu()
			case 3:
				EndTrip()//End trip for confirmed booking
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

//function to view all pending bookings
func ViewPendingBookings() {
	client := &http.Client{}
	//send GET request
	if req, err := http.NewRequest(http.MethodGet, "http://localhost:5003/api/v1/pendingbookings", nil); err == nil {
		if res, err := client.Do(req); err == nil {
			if body, err := ioutil.ReadAll(res.Body); err == nil {
				var res Bookings
				json.Unmarshal(body, &res)//Decoding from JSON
				//print booking details
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

//function to initiate start trip
func InitiateStartTrip() {
	var booking Booking
	var bookingID string
	fmt.Println()
	fmt.Print("Enter Booking ID to initiate trip: ")
	fmt.Scan(&(bookingID))
	fmt.Print("Enter Driver ID: ")
	fmt.Scan(&(booking.DriverID))
	booking.BookingStatus = "Confimed"//booking status becomes confirmed

	postBody, _ := json.Marshal(booking)//Sending PUT Request with data
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
//function to end trip 
func EndTrip() {
	var booking Booking
	var bookingID string
	fmt.Print("Enter Booking ID to end trip: ")
	fmt.Scan(&(bookingID))
	booking.BookingStatus = "Completed"//booking status becomes completed

	postBody, _ := json.Marshal(booking)//Sending PUT Request with data
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
