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

type Driver struct { // map this type to the record in the table
	DriverID  string
	FirstName string
	LastName  string
	IdenNo    string
	CarNo     string
	Mobile    string
	Email     string
	Password  string
}

func main() {

	router := mux.NewRouter()
	router.HandleFunc("/api/v1/driver/{driverid}", driver).Methods("GET", "POST", "PUT")
	//router.HandleFunc("/api/v1/courses", allpassengers)
	fmt.Println("Listening at port 5000")
	log.Fatal(http.ListenAndServe(":5000", router))
}
func driver(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	if r.Method == "POST" {
		if body, err := ioutil.ReadAll(r.Body); err == nil {
			var data Driver

			if err := json.Unmarshal(body, &data); err == nil {
				if _, ok := isExist(params["driverid"]); !ok {
					fmt.Println(data)
					//courses[params["courseid"]] = data
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

			if err := json.Unmarshal(body, &data); err == nil {
				if _, ok := isExist(params["driverid"]); ok {
					fmt.Println(data)

					//courses[params["courseid"]] = data
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

	} else {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "Invalid driver ID")
	}
}
func isExist(id string) (Driver, bool) {
	db, err := sql.Open("mysql", "user:password@tcp(127.0.0.1:3306)/driver_db")

	// handle error
	if err != nil {
		panic(err.Error())
	}

	defer db.Close()
	results := db.QueryRow("select * from driver where id=?", id)

	var d Driver
	//var id string
	err = results.Scan(&id, &d.FirstName, &d.LastName, &d.IdenNo, &d.CarNo, &d.Mobile, &d.Email, &d.Password)
	if err == sql.ErrNoRows {
		return d, false
	}
	return d, true
}
func insertDriver(id string, d Driver) {
	db, err := sql.Open("mysql", "user:password@tcp(127.0.0.1:3306)/driver_db")

	// handle error
	if err != nil {
		panic(err.Error())
	}

	defer db.Close()
	_, err = db.Exec("insert into driver values(?,?,?,?,?)", id, d.FirstName, d.LastName, d.IdenNo, d.CarNo, d.Mobile, d.Email, d.Password)
	if err != nil {
		panic(err.Error())
	}
	//use curl to insert in cmd
}
func updateDriver(id string, d Driver) {
	db, err := sql.Open("mysql", "user:password@tcp(127.0.0.1:3306)/driver_db")

	// handle error
	if err != nil {
		panic(err.Error())
	}

	defer db.Close()
	_, err = db.Exec("update driver set fname=?, lname=?,cNo=?, mobile=?,email=?,pw=? where id=?", d.FirstName, d.LastName, d.CarNo, d.Mobile, d.Email, d.Password, id)
	if err != nil {
		panic(err.Error())
	}
	//use curl to insert in cmd
}
