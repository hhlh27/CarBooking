package main

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

type Booking struct { // map this type to the record in the table

	PassengerID     string `json:"Passenger Id"`
	DriverID        string `json:"Driver Id"`
	PickUp          string `json:"Pickup"`
	DropOff         string `json:"Dropoff"`
	BookingDateTime string `json:"Booking DateTime"`
	BookingStatus   string `json:"Booking Status"`
}

var (
	db  *sql.DB
	err error
)

func main() {

	router := mux.NewRouter()
	router.HandleFunc("/api/v1/booking/{bookingid}", booking).Methods("GET", "POST", "PUT")
	router.HandleFunc("/api/v1/bookings", allbookings).Methods("GET")
	router.HandleFunc("/api/v1/pendingbookings", pendingbookings).Methods("GET")
	router.HandleFunc("/api/v1/passengerbookings", passengerbooking).Methods("GET", "POST", "PUT")
	//router.HandleFunc("/api/v1/booking/{passengerid}", passengerBooking).Methods("GET")
	//router.HandleFunc("/api/v1/courses", allpassengers)
	fmt.Println("Listening at port 5003")
	log.Fatal(http.ListenAndServe(":5003", router))
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

	} else if val, ok := isExist(params["bookingid"]); ok {

		json.NewEncoder(w).Encode(val)

	} else {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "Invalid Booking ID")
	}
}
func isExist(id string) (Booking, bool) {
	/* db, err := sql.Open("mysql", "user:password@tcp(127.0.0.1:3306)/booking_db")

	// handle error
	if err != nil {
		panic(err.Error())
	}

	defer db.Close() */
	var b Booking
	results := db.QueryRow("select * from booking where id=?", id)

	//var id string
	err = results.Scan(&id, &b.PassengerID, &b.DriverID, &b.PickUp, &b.DropOff, &b.BookingDateTime, &b.BookingStatus)
	if err == sql.ErrNoRows {
		return b, false
	}
	return b, true
}
func insertBooking(id string, b Booking) {
	/* db, err := sql.Open("mysql", "user:password@tcp(127.0.0.1:3306)/booking_db")

	// handle error
	if err != nil {
		panic(err.Error())
	} */

	//defer db.Close()
	_, err = db.Exec("insert into booking values(?,?,?,?,?)", id, b.PassengerID, b.DriverID, b.PickUp, b.DropOff, b.BookingDateTime, b.BookingStatus)
	if err != nil {
		panic(err.Error())
	}
	//use curl to insert in cmd
}
func updateBooking(id string, b Booking) {
	/* db, err := sql.Open("mysql", "user:password@tcp(127.0.0.1:3306)/booking_db")

	// handle error
	if err != nil {
		panic(err.Error())
	}

	defer db.Close() */
	_, err = db.Exec("update booking set status=?", b.BookingStatus)
	if err != nil {
		panic(err.Error())
	}
	//use curl to insert in cmd
}

func searchBooking(q string) (map[string]Booking, bool) {
	db, err := sql.Open("mysql", "user:password@tcp(127.0.0.1:3306)/booking_db")

	// handle error
	if err != nil {
		panic(err.Error())
	}

	defer db.Close()
	results, err := db.Query("select * from booking where lower(passengerid) like lower(?) ", "%"+q+"%") //db query-> returns 2 things, dg query row returns 1 thing only
	if err != nil {
		panic(err.Error())
	}
	var bookingsDB map[string]Booking = map[string]Booking{}
	for results.Next() {
		var b Booking
		var id string
		err = results.Scan(&id, &b.BookingDateTime, &b.DriverID, &b.PickUp, &b.DropOff)
		if err != nil {
			panic(err.Error())
		}
		bookingsDB[id] = b
	}

	if len(bookingsDB) == 0 {
		return bookingsDB, false
	}
	return bookingsDB, true
}
func pendingbookings(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()

	/* found := false
	results := map[string]Booking{} */

	if value := query.Get("q"); len(value) > 0 {
		/* for k, v := range bookings {
			if strings.Contains(strings.ToLower(v.PassengerID), strings.ToLower(value)) {
				results[k] = v
				found = true
			}
		} */
		results, found := queryBookings(value)
		if !found {
			fmt.Fprintf(w, "No booking found")
		} else {
			json.NewEncoder(w).Encode(struct {
				Results map[string]Booking `json:"Search Results"`
			}{results})
		}

	} else {
		bookingsWrapper := struct {
			Bookings map[string]Booking `json:"Bookings"`
		}{getPendingBookings()}
		json.NewEncoder(w).Encode(bookingsWrapper)
		return
	}

}
func allbookings(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()

	/* found := false
	results := map[string]Booking{} */

	if value := query.Get("q"); len(value) > 0 {
		/* for k, v := range bookings {
			if strings.Contains(strings.ToLower(v.PassengerID), strings.ToLower(value)) {
				results[k] = v
				found = true
			}
		} */
		results, found := queryBookings(value)
		if !found {
			fmt.Fprintf(w, "No booking found")
		} else {
			json.NewEncoder(w).Encode(struct {
				Results map[string]Booking `json:"Search Results"`
			}{results})
		}

	} else {
		bookingsWrapper := struct {
			Bookings map[string]Booking `json:"Bookings"`
		}{getBookings()}
		json.NewEncoder(w).Encode(bookingsWrapper)
		return
	}
}
func passengerbooking(w http.ResponseWriter, r *http.Request) {
	//var  passengerid string
	query := r.URL.Query()

	/* found := false
	results := map[string]Booking{} */

	if value := query.Get("q"); len(value) > 0 {
		/* for k, v := range bookings {
			if strings.Contains(strings.ToLower(v.PassengerID), strings.ToLower(value)) {
				results[k] = v
				found = true
			}
		} */
		results, found := queryBookings(value)
		if !found {
			fmt.Fprintf(w, "No booking found")
		} else {
			json.NewEncoder(w).Encode(struct {
				Results map[string]Booking `json:"Search Results"`
			}{results})
		}

	} else {
		bookingsWrapper := struct {
			Bookings map[string]Booking `json:"Bookings"`
		}{getPassengerBookings()}
		json.NewEncoder(w).Encode(bookingsWrapper)
		return
	}

}
func getBookings() map[string]Booking {
	results, err := db.Query("select * from Booking")
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

func queryBookings(query string) (map[string]Booking, bool) {
	results, err := db.Query("select * from Booking where lower(passengerid) like lower(?)", "%"+query+"%")
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
func getPassengerBookings() map[string]Booking {
	var query string
	results, err := db.Query("select * from Booking where lower(passengerid) like lower(?)", "%"+query+"%")
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
func getPendingBookings() map[string]Booking {
	var query string
	results, err := db.Query("select * from Booking where lower(bookingstatus) like lower(?)", "%"+query+"%")
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
