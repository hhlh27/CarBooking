package main

//import required packages
import (
	"database/sql"
	"encoding/json"
	"fmt"

	"github.com/gorilla/mux"

	"io/ioutil"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

// Defining Passenger struct
type Booking struct { // map this type to the record in the table

	PassengerID     string `json:"Passenger Id"`
	DriverID        string `json:"Driver Id"`
	PickUp          string `json:"Pickup"`
	DropOff         string `json:"Dropoff"`
	BookingDateTime string `json:"Booking DateTime"`
	BookingStatus   string `json:"Booking Status"`
}

// create database object
var (
	db  *sql.DB
	err error
)

func main() {
	db, err = sql.Open("mysql", "user:password@tcp(127.0.0.1:3306)/db") //Connecting to MySQL DB
	if err != nil {
		panic(err.Error())
	}
	//initialise router
	router := mux.NewRouter()
	//creating the endpoints, register REST URL
	router.HandleFunc("/api/v1/booking/{bookingid}", booking).Methods("GET", "POST", "PUT") //register HTTP methods GET,POST,PUT
	router.HandleFunc("/api/v1/bookings", allbookings).Methods("GET")                       //Register HTTP methods GET
	router.HandleFunc("/api/v1/pendingbookings", pendingbookings).Methods("GET")            //Register HTTP methods GET
	fmt.Println("Listening at port 5003")
	log.Fatal(http.ListenAndServe(":5003", router))
}

// booking function-handles methods GET,POST,PUT
func booking(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	if r.Method == "POST" {
		if body, err := ioutil.ReadAll(r.Body); err == nil {
			var data Booking

			if err := json.Unmarshal(body, &data); err == nil { //Decoding from JSON
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

			if err := json.Unmarshal(body, &data); err == nil { //Decoding from JSON
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

	} else if val, ok := isExist(params["bookingid"]); ok {

		json.NewEncoder(w).Encode(val) //Encoding into JSON

	} else {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "Invalid Booking ID")
	}
}

// function to check if booking exists using booking id
func isExist(id string) (Booking, bool) {
	var b Booking
	results := db.QueryRow("select * from booking where id=?", id) //Retrieving from DB

	//var id string
	err = results.Scan(&id, &b.PassengerID, &b.DriverID, &b.PickUp, &b.DropOff, &b.BookingDateTime, &b.BookingStatus)
	if err == sql.ErrNoRows {
		return b, false
	}
	return b, true
}

// function to insert a new booking record into database
func insertBooking(id string, b Booking) {
	//inserting into db
	_, err = db.Exec("insert into booking values(?,?,?,?,?)", id, b.PassengerID, b.DriverID, b.PickUp, b.DropOff, b.BookingDateTime, b.BookingStatus)
	if err != nil {
		panic(err.Error())
	}
}

// function to update booking record
func updateBooking(id string, b Booking) {
	//updating record in db
	_, err = db.Exec("update booking set status=?", b.BookingStatus)
	if err != nil {
		panic(err.Error())
	}
}

// function to display and format all pending booking records
func pendingbookings(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query() //Retrieving Value from Query String
	if value := query.Get("q"); len(value) > 0 {

		results, found := queryBookings(value)
		if !found {
			fmt.Fprintf(w, "No booking found")
		} else {
			json.NewEncoder(w).Encode(struct { //Encoding into JSON
				Results map[string]Booking `json:"Search Results"`
			}{results})
		}

	} else {
		bookingsWrapper := struct {
			Bookings map[string]Booking `json:"Bookings"`
		}{getPendingBookings()}
		json.NewEncoder(w).Encode(bookingsWrapper) //Encoding into JSON
		return
	}

}

// function to display and format all booking records
func allbookings(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query() //Retrieving Value from Query String

	if value := query.Get("q"); len(value) > 0 {
		results, found := queryBookings(value)
		if !found {
			fmt.Fprintf(w, "No booking found")
		} else {
			json.NewEncoder(w).Encode(struct { //Encoding into JSON
				Results map[string]Booking `json:"Search Results"`
			}{results})
		}

	} else if value = query.Get("passengerid"); len(value) > 0 {

		passengerid := value
		results, found := findPassengerBookings(passengerid)

		if !found {
			fmt.Fprintf(w, "No bookings found")
		} else {
			json.NewEncoder(w).Encode(struct { //Encoding into JSON
				Results map[string]Booking `json:"Bookings"`
			}{results})
		}
	} else {
		bookingsWrapper := struct {
			Bookings map[string]Booking `json:"Bookings"`
		}{getBookings()}
		json.NewEncoder(w).Encode(bookingsWrapper) //Encoding into JSON
		return
	}
}

// function to get all booking records
func getBookings() map[string]Booking {
	results, err := db.Query("select * from Booking") //Retrieving from DB
	if err != nil {
		panic(err.Error())
	}

	var bookings map[string]Booking = map[string]Booking{}

	for results.Next() {
		var b Booking
		var id string
		err = results.Scan(&id, &b.BookingDateTime, &b.DriverID, &b.PickUp, &b.DropOff)

		if err != nil {
			panic(err.Error())
		}
		bookings[id] = b
	}

	return bookings
}

// function to search for booking record
func queryBookings(query string) (map[string]Booking, bool) {
	results, err := db.Query("select * from Booking where lower(passengerid) like lower(?)", "%"+query+"%") //Retrieving from DB
	if err != nil {
		panic(err.Error())
	}

	var bookings map[string]Booking = map[string]Booking{}

	for results.Next() {
		var b Booking
		var id string
		err = results.Scan(&id, &b.BookingDateTime, &b.DriverID, &b.PickUp, &b.DropOff)
		if err != nil {
			panic(err.Error())
		}

		bookings[id] = b
	}

	if len(bookings) == 0 {
		return bookings, false
	}
	return bookings, true
}

// function to search for booking records based on passenger id
func findPassengerBookings(query string) (map[string]Booking, bool) {
	results, err := db.Query("select * from Booking where lower(passengerid) like lower(?) order by bookingdatetime desc", "%"+query+"%") //Retrieving from DB
	if err != nil {
		panic(err.Error())
	}

	var bookings map[string]Booking = map[string]Booking{}

	for results.Next() {
		var b Booking
		var id string
		err = results.Scan(&id, &b.BookingDateTime, &b.DriverID, &b.PickUp, &b.DropOff)

		if err != nil {
			panic(err.Error())
		}
		bookings[id] = b
	}

	if len(bookings) == 0 {
		return bookings, false
	}
	return bookings, true
}

// function to get all pending booking records
func getPendingBookings() map[string]Booking {
	var query string
	query = "pending"
	results, err := db.Query("select * from Booking where lower(bookingstatus) like lower(?)", "%"+query+"%") //Retrieving from DB
	if err != nil {
		panic(err.Error())
	}

	var bookings map[string]Booking = map[string]Booking{}

	for results.Next() {
		var b Booking
		var id string
		err = results.Scan(&id, &b.BookingDateTime, &b.DriverID, &b.PickUp, &b.DropOff)

		if err != nil {
			panic(err.Error())
		}
		bookings[id] = b
	}

	return bookings
}
