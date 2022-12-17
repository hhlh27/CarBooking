package main

//import required packages
import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	_ "github.com/go-sql-driver/mysql"
)

// Defining driver struct
type Driver struct { // map this type to the record in the table
	FirstName string
	LastName  string
	CarNo     string
	Mobile    string
	Email     string
	Password  string
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
	router.HandleFunc("/api/v1/driver/{driverid}", driver).Methods("GET", "POST", "PUT") //register HTTP methods GET,POST,PUT
	router.HandleFunc("/api/v1/drivers", alldrivers)
	fmt.Println("Listening at port 7000")
	log.Fatal(http.ListenAndServe(":7000", router))
}

// driver function-handles methods GET,POST,PUT
func driver(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	if r.Method == "POST" {
		if body, err := ioutil.ReadAll(r.Body); err == nil {
			var data Driver

			if err := json.Unmarshal(body, &data); err == nil { //Decoding from JSON
				if _, ok := isExist(params["driverid"]); !ok {
					fmt.Println(data)
					insertDriver(params["driverid"], data)

					w.WriteHeader(http.StatusAccepted)
				} else {
					w.WriteHeader(http.StatusConflict)
					fmt.Fprintf(w, "Driver ID exist")
				}
			} else {
				fmt.Println(err)
			}
		}
	} else if r.Method == "PUT" {
		if body, err := ioutil.ReadAll(r.Body); err == nil {
			var data Driver

			if err := json.Unmarshal(body, &data); err == nil { //Decoding from JSON
				if _, ok := isExist(params["driverid"]); ok {
					fmt.Println(data)
					updateDriver(params["driverid"], data)
					w.WriteHeader(http.StatusAccepted)
				} else {
					w.WriteHeader(http.StatusNotFound)
					fmt.Fprintf(w, "Driver ID does not exist")
				}
			} else {
				fmt.Println(err)
			}
		}

	} else if val, ok := isExist(params["driverid"]); ok {

		json.NewEncoder(w).Encode(val) //Encoding into JSON

	} else {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "Invalid driver ID")
	}
}

// function to check if driver exists using driver id
func isExist(id string) (Driver, bool) {
	var d Driver
	results := db.QueryRow("select * from driver where id=?", id)
	err = results.Scan(&id, &d.FirstName, &d.LastName, &d.CarNo, &d.Mobile, &d.Email, &d.Password)
	if err == sql.ErrNoRows {
		return d, false
	}
	return d, true
}

// function to insert a new driver record into database
func insertDriver(id string, d Driver) {
	//inserting into db
	_, err = db.Exec("insert into driver values(?,?,?,?,?)", id, d.FirstName, d.LastName, d.CarNo, d.Mobile, d.Email, d.Password)
	if err != nil {
		panic(err.Error())
	}

}

// function to update driver record
func updateDriver(id string, d Driver) {
	_, err = db.Exec("update driver set fname=?, lname=?,cNo=?, mobile=?,email=?,pw=? where id=?", d.FirstName, d.LastName, d.CarNo, d.Mobile, d.Email, d.Password, id)
	if err != nil {
		panic(err.Error())
	}

}

// function to display and format all driver records
func alldrivers(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	if value := query.Get("q"); len(value) > 0 {
		results, found := queryDrivers(value)

		if !found {
			fmt.Fprintf(w, "No driver found")
		} else {
			json.NewEncoder(w).Encode(struct {
				Results map[string]Driver `json:"driver Search Results"`
			}{results})
		}

	} else {
		driversWrapper := struct {
			Drivers map[string]Driver `json:"Drivers"`
		}{getDrivers()}
		json.NewEncoder(w).Encode(driversWrapper) //Encoding into JSON
		return
	}
}

// function to get all driver records
func getDrivers() map[string]Driver {
	results, err := db.Query("select * from Driver") //Retrieving from DB
	if err != nil {
		panic(err.Error())
	}

	var drivers map[string]Driver = map[string]Driver{}

	for results.Next() {
		var d Driver
		var id string
		err = results.Scan(&id, &d.FirstName, &d.LastName, &d.CarNo, &d.Mobile, &d.Email)
		if err != nil {
			panic(err.Error())
		}

		drivers[id] = d
	}

	return drivers
}

// function to search for a driver record
func queryDrivers(query string) (map[string]Driver, bool) {
	results, err := db.Query("select * from Driver where lower(firstname) like lower(?)", "%"+query+"%")
	if err != nil {
		panic(err.Error())
	}

	var drivers map[string]Driver = map[string]Driver{}

	for results.Next() {
		var d Driver
		var id string
		err = results.Scan(&id, &d.FirstName, &d.LastName, &d.CarNo, &d.Mobile, &d.Email)
		if err != nil {
			panic(err.Error())
		}

		drivers[id] = d
	}

	if len(drivers) == 0 {
		return drivers, false
	}
	return drivers, true
}
