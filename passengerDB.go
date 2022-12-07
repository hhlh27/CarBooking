package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Passenger struct { // map this type to the record in the table
	ID        string
	FirstName string
	LastName  string
	Mobile    string
	Email     string
	Password  string
}

func main() {

	router := mux.NewRouter()
	router.HandleFunc("/api/v1/passenger/{passengerid}", passenger).Methods("GET", "POST", "PUT")
	//router.HandleFunc("/api/v1/courses", allpassengers)
	fmt.Println("Listening at port 5000")
	log.Fatal(http.ListenAndServe(":5000", router))
}
func passenger(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	if r.Method == "POST" {
		if body, err := ioutil.ReadAll(r.Body); err == nil {
			var data Passenger

			if err := json.Unmarshal(body, &data); err == nil {
				if _, ok := isExist(params["passengerid"]); !ok {
					fmt.Println(data)
					//courses[params["courseid"]] = data
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

			if err := json.Unmarshal(body, &data); err == nil {
				if _, ok := isExist(params["passengerid"]); ok {
					fmt.Println(data)

					//courses[params["courseid"]] = data
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

	} else {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "Invalid passenger ID")
	}
}
func isExist(id string) (Passenger, bool) {
	db, err := sql.Open("mysql", "user:password@tcp(127.0.0.1:3306)/passenger_db")

	// handle error
	if err != nil {
		panic(err.Error())
	}

	defer db.Close()
	results := db.QueryRow("select * from passenger where id=?", id)

	var p Passenger
	//var id string
	err = results.Scan(&id, &p.FirstName, &p.LastName, &p.Mobile, &p.Email, &p.Password)
	if err == sql.ErrNoRows {
		return p, false
	}
	return p, true
}
func insertPassenger(id string, p Passenger) {
	db, err := sql.Open("mysql", "user:password@tcp(127.0.0.1:3306)/passenger_db")

	// handle error
	if err != nil {
		panic(err.Error())
	}

	defer db.Close()
	_, err = db.Exec("insert into passenger values(?,?,?,?,?)", id, p.FirstName, p.LastName, p.Mobile, p.Email, p.Password)
	if err != nil {
		panic(err.Error())
	}
	//use curl to insert in cmd
}
func updatePassenger(id string, p Passenger) {
	db, err := sql.Open("mysql", "user:password@tcp(127.0.0.1:3306)/passenger_db")

	// handle error
	if err != nil {
		panic(err.Error())
	}

	defer db.Close()
	_, err = db.Exec("update passenger set fname=?, lname=?, mobile=?,email=?,pw=? where id=?", p.FirstName, p.LastName, p.Mobile, p.Email, p.Password, id)
	if err != nil {
		panic(err.Error())
	}
	//use curl to insert in cmd
}
