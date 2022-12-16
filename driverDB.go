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
	//DriverID  string
	FirstName string
	LastName  string
	CarNo     string
	Mobile    string
	Email     string
	Password  string
}

var (
	db  *sql.DB
	err error
)

func main() {
	db, err = sql.Open("mysql", "user:password@tcp(127.0.0.1:3306)/db")

	if err != nil {
		panic(err.Error())
	}
	router := mux.NewRouter()
	router.HandleFunc("/api/v1/driver/{driverid}", driver).Methods("GET", "POST", "PUT")
	//router.HandleFunc("/api/v1/courses", allpassengers)
	router.HandleFunc("/api/v1/drivers", alldrivers)
	fmt.Println("Listening at port 7000")
	log.Fatal(http.ListenAndServe(":7000", router))
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

	} else if val, ok := isExist(params["driverid"]); ok {

		json.NewEncoder(w).Encode(val)

	} else {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "Invalid driver ID")
	}
}
func isExist(id string) (Driver, bool) {
	/* //db, err := sql.Open("mysql", "user:password@tcp(127.0.0.1:3306)/driver_db")

	// handle error
	if err != nil {
		panic(err.Error())
	}

	defer db.Close() */
	var d Driver
	results := db.QueryRow("select * from driver where id=?", id)

	//var id string
	err = results.Scan(&id, &d.FirstName, &d.LastName, &d.CarNo, &d.Mobile, &d.Email, &d.Password)
	if err == sql.ErrNoRows {
		return d, false
	}
	return d, true
}
func insertDriver(id string, d Driver) {
	/* db, err := sql.Open("mysql", "user:password@tcp(127.0.0.1:3306)/driver_db")

	// handle error
	if err != nil {
		panic(err.Error())
	}

	defer db.Close() */
	_, err = db.Exec("insert into driver values(?,?,?,?,?)", id, d.FirstName, d.LastName, d.CarNo, d.Mobile, d.Email, d.Password)
	if err != nil {
		panic(err.Error())
	}
	//use curl to insert in cmd
}
func updateDriver(id string, d Driver) {
	/* db, err := sql.Open("mysql", "user:password@tcp(127.0.0.1:3306)/driver_db")

	// handle error
	if err != nil {
		panic(err.Error())
	} */

	//defer db.Close()
	_, err = db.Exec("update driver set fname=?, lname=?,cNo=?, mobile=?,email=?,pw=? where id=?", d.FirstName, d.LastName, d.CarNo, d.Mobile, d.Email, d.Password, id)
	if err != nil {
		panic(err.Error())
	}
	//use curl to insert in cmd
}
func alldrivers(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()

	/* found := false
	results := map[string]Driver{} */

	if value := query.Get("q"); len(value) > 0 {
		//for k, v := range drivers {
		/* 	if strings.Contains(strings.ToLower(v.FirstName), strings.ToLower(value)) {
				results[k] = v
				found = true
			}
		} */
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
		json.NewEncoder(w).Encode(driversWrapper)
		return
	}
}
func getDrivers() map[string]Driver {
	results, err := db.Query("select * from Driver")
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
