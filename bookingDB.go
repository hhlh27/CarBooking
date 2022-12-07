package main

import (
	"database/sql"
	"encoding/json"
	"fmt"

	"github.com/gorilla/mux"

	"io/ioutil"
	"log"
	"net/http"
	"time"

	_ "github.com/go-sql-driver/mysql"
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

func main() {

	router := mux.NewRouter()
	router.HandleFunc("/api/v1/booking/{bookingid}", booking).Methods("GET", "POST", "PUT")
	//router.HandleFunc("/api/v1/courses", allpassengers)
	fmt.Println("Listening at port 5000")
	log.Fatal(http.ListenAndServe(":5000", router))
}
func booking(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	if r.Method == "POST" {
		if body, err := ioutil.ReadAll(r.Body); err == nil {
			var data Booking

			if err := json.Unmarshal(body, &data); err == nil {
				if _, ok := isExist(params["bookingid"]); !ok {
					fmt.Println(data)
					//courses[params["courseid"]] = data
					insertBooking(params["bookingid"], data)

					w.WriteHeader(http.StatusAccepted)
				} else {
					w.WriteHeader(http.StatusConflict)
					fmt.Fprintf(w, "Booking ID exist")
				}
			} else {
				fmt.Println(err)
			}
		}
	} else if r.Method == "PUT" {
		if body, err := ioutil.ReadAll(r.Body); err == nil {
			var data Booking

			if err := json.Unmarshal(body, &data); err == nil {
				if _, ok := isExist(params["bookingid"]); ok {
					fmt.Println(data)

					//courses[params["courseid"]] = data
					updateBooking(params["bookingid"], data)
					w.WriteHeader(http.StatusAccepted)
				} else {
					w.WriteHeader(http.StatusNotFound)
					fmt.Fprintf(w, "Booking ID does not exist")
				}
			} else {
				fmt.Println(err)
			}
		}

	} else {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "Invalid Booking ID")
	}
}
func isExist(id string) (Booking, bool) {
	db, err := sql.Open("mysql", "user:password@tcp(127.0.0.1:3306)/booking_db")

	// handle error
	if err != nil {
		panic(err.Error())
	}

	defer db.Close()
	results := db.QueryRow("select * from booking where id=?", id)

	var b Booking
	//var id string
	err = results.Scan(&id, &b.PassengerID, &b.DriverID, &b.PickUp, &b.DropOff, &b.BookingDateTime, &b.BookingStatus)
	if err == sql.ErrNoRows {
		return b, false
	}
	return b, true
}
func insertBooking(id string, b Booking) {
	db, err := sql.Open("mysql", "user:password@tcp(127.0.0.1:3306)/booking_db")

	// handle error
	if err != nil {
		panic(err.Error())
	}

	defer db.Close()
	_, err = db.Exec("insert into booking values(?,?,?,?,?)", id, b.PassengerID, b.DriverID, b.PickUp, b.DropOff, b.BookingDateTime, b.BookingStatus)
	if err != nil {
		panic(err.Error())
	}
	//use curl to insert in cmd
}
func updateBooking(id string, b Booking) {
	db, err := sql.Open("mysql", "user:password@tcp(127.0.0.1:3306)/booking_db")

	// handle error
	if err != nil {
		panic(err.Error())
	}

	defer db.Close()
	_, err = db.Exec("update booking set status=?", b.BookingStatus)
	if err != nil {
		panic(err.Error())
	}
	//use curl to insert in cmd
}
