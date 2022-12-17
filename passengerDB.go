package main

//import required packages
import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

// Defining Passenger struct
type Passenger struct { // map this type to the record in the table
	FirstName string `json:"First Name"`
	LastName  string `json:"Last Name"`
	Mobile    string `json:"Mobile"`
	Email     string `json:"Email"`
	Password  string `json:"Password"`
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
	router.HandleFunc("/api/v1/passengers/{passengerid}", passenger).Methods("GET", "POST", "PUT") //register HTTP methods GET,POST,PUT
	router.HandleFunc("/api/v1/passengers", allPassengers).Methods("GET")                          //Register HTTP methods GET
	fmt.Println("Listening at port 5000")
	log.Fatal(http.ListenAndServe(":5000", router))
}

// passenger function-handles methods GET,POST,PUT
func passenger(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	if r.Method == "POST" {
		if body, err := ioutil.ReadAll(r.Body); err == nil {
			var data Passenger
			fmt.Println(string(body))
			if err := json.Unmarshal(body, &data); err == nil { //Decoding from JSON
				if _, ok := isExist(params["passengerid"]); !ok {
					fmt.Println(data)
					insertPassenger(params["passengerid"], data)
					w.WriteHeader(http.StatusAccepted)
				} else {
					w.WriteHeader(http.StatusConflict)
					fmt.Fprintf(w, "passenger ID exist")
				}

			} else {
				fmt.Println(err)
			}
		}
	} else if r.Method == "PUT" {
		if body, err := ioutil.ReadAll(r.Body); err == nil {
			var data Passenger

			if err := json.Unmarshal(body, &data); err == nil { //Decoding from JSON
				if _, ok := isExist(params["passengerid"]); ok {
					fmt.Println(data)
					updatePassenger(params["passengerid"], data)
					w.WriteHeader(http.StatusAccepted)
				} else {
					w.WriteHeader(http.StatusNotFound)
					fmt.Fprintf(w, "passenger ID does not exist")
				}
			} else {
				fmt.Println(err)
			}
		}

	} else if val, ok := isExist(params["passengerid"]); ok {

		json.NewEncoder(w).Encode(val) //Encoding into JSON
	} else {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "Invalid passenger ID")
	}
}

// function to check if passenger exists using passenger id
func isExist(id string) (Passenger, bool) {
	var p Passenger
	results := db.QueryRow("select * from passenger where id=?", id) //Retrieving from DB
	err = results.Scan(&id, &p.FirstName, &p.LastName, &p.Mobile, &p.Email, &p.Password)
	if err == sql.ErrNoRows {
		return p, false
	}
	return p, true
}

// function to insert a new passenger record into database
func insertPassenger(id string, p Passenger) {
	//inserting into db
	_, err = db.Exec("insert into Passenger values(?,?,?,?,?,?)", id, p.FirstName, p.LastName, p.Mobile, p.Email, p.Password)
	if err != nil {
		panic(err.Error())
	}
}

// function to update passenger record
func updatePassenger(id string, p Passenger) {
	//updating record in db
	_, err = db.Exec("update Passenger set fname=?, lname=?, mobile=?,email=?,pw=? where id=?", p.FirstName, p.LastName, p.Mobile, p.Email, p.Password, id)
	if err != nil {
		panic(err.Error())
	}
}

// function to display and format all passenger records
func allPassengers(w http.ResponseWriter, r *http.Request) {

	query := r.URL.Query() //Retrieving Value from Query String

	if value := query.Get("q"); len(value) > 0 {
		results, found := queryPassengers(value)

		if !found {
			fmt.Fprintf(w, "No passenger found")
		} else {
			json.NewEncoder(w).Encode(struct {
				Results map[string]Passenger `json:"Search Results"`
			}{results})
		}

	} else {
		passengersWrapper := struct {
			Passengers map[string]Passenger `json:"Passengers"`
		}{getPassengers()}
		json.NewEncoder(w).Encode(passengersWrapper) //Encoding into JSON
		return
	}
}

// function to get all passenger records
func getPassengers() map[string]Passenger {
	results, err := db.Query("select * from passenger") //Retrieving from DB
	if err != nil {
		panic(err.Error())
	}

	var passengers map[string]Passenger = map[string]Passenger{}

	for results.Next() {
		var p Passenger
		var id string
		err = results.Scan(&id, &p.FirstName, &p.LastName, &p.Email, &p.Mobile)
		if err != nil {
			panic(err.Error())
		}

		passengers[id] = p
	}

	return passengers
}

// function to search for passenger record
func queryPassengers(query string) (map[string]Passenger, bool) {
	results, err := db.Query("select * from Passenger where lower(name) like lower(?)", "%"+query+"%") //Retrieving from DB
	if err != nil {
		panic(err.Error())
	}

	var passengers map[string]Passenger = map[string]Passenger{}

	for results.Next() {
		var p Passenger
		var id string
		err = results.Scan(&id, &p.FirstName, &p.LastName, &p.Email, &p.Mobile)
		if err != nil {
			panic(err.Error())
		}

		passengers[id] = p
	}

	if len(passengers) == 0 {
		return passengers, false
	}
	return passengers, true
}
