package main

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

type Passenger struct { // map this type to the record in the table
	ID        string `json:"Id"`
	FirstName string `json:"First Name"`
	LastName  string `json:"Last Name"`
	Mobile    string `json:"Mobile"`
	Email     string `json:"Email"`
	Password  string `json:"Password"`
}

var (
	db  *sql.DB
	err error
)

// var passengers map[string]Passenger = map[string]Passenger{}

func main() {
	db, err = sql.Open("mysql", "user:password@tcp(127.0.0.1:3306)/db")

	if err != nil {
		panic(err.Error())
	}
	router := mux.NewRouter()
	router.HandleFunc("/api/v1/passengers/{passengerid}", passenger).Methods("GET", "POST", "PUT")
	//router.HandleFunc("/api/v1/courses", allpassengers)
	router.HandleFunc("/api/v1/passengers", allPassengers).Methods("GET")
	fmt.Println("Listening at port 5001")
	log.Fatal(http.ListenAndServe(":5001", router))
}
func passenger(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	if r.Method == "POST" {
		if body, err := ioutil.ReadAll(r.Body); err == nil {
			var data Passenger
			fmt.Println(string(body))
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
	/* db, err := sql.Open("mysql", "user:password@tcp(127.0.0.1:3306)/passenger_db")

	// handle error
	if err != nil {
		panic(err.Error())
	}

	defer db.Close() */
	var p Passenger
	results := db.QueryRow("select * from passenger where id=?", id)
	//var id string
	err = results.Scan(&id, &p.FirstName, &p.LastName, &p.Mobile, &p.Email, &p.Password)
	if err == sql.ErrNoRows {
		return p, false
	}
	return p, true
}
func insertPassenger(id string, p Passenger) {
	// db, err := sql.Open("mysql", "user:password@tcp(127.0.0.1:3306)/passenger_db")

	// handle error
	/* if err != nil {
		panic(err.Error())
	} */

	//defer db.Close()
	_, err = db.Exec("insert into Passenger values(?,?,?,?,?,?)", id, p.FirstName, p.LastName, p.Mobile, p.Email, p.Password)
	if err != nil {
		panic(err.Error())
	}
	//use curl to insert in cmd
}
func updatePassenger(id string, p Passenger) {
	/* db, err := sql.Open("mysql", "user:password@tcp(127.0.0.1:3306)/passenger_db")

	// handle error
	if err != nil {
		panic(err.Error())
	}

	defer db.Close() */
	_, err = db.Exec("update Passenger set fname=?, lname=?, mobile=?,email=?,pw=? where id=?", p.FirstName, p.LastName, p.Mobile, p.Email, p.Password, id)
	if err != nil {
		panic(err.Error())
	}
	//use curl to insert in cmd
}
func allPassengers(w http.ResponseWriter, r *http.Request) {

	query := r.URL.Query()

	/* found := false
	results := map[string]Passenger{} */

	if value := query.Get("q"); len(value) > 0 {
		//results, found := searchPassengerName(query.Get("q"))
		results, found := queryPassengers(value)

		if !found {
			fmt.Fprintf(w, "No passenger found")
		} else {
			json.NewEncoder(w).Encode(struct {
				Results map[string]Passenger `json:"Search Results"`
			}{results})
		}

		/* else {
		db, err := sql.Open("mysql", "user:password@tcp(127.0.0.1:3306)/passenger_db")

		// handle error
		if err != nil {
			panic(err.Error())
		}

		defer db.Close()
		results, err := db.Query("select * from Passenger")
		if err != nil {
			panic(err.Error()) //kill app when have error, not reccommeded
		}
		var passengersDB map[string]Passenger = map[string]Passenger{}

		for results.Next() {
			var p Passenger
			var id string
			err = results.Scan(&id, &p.FirstName, &p.LastName, &p.Email, &p.Mobile)
			if err != nil {
				panic(err.Error())
			}
			passengersDB[id] = p
		}
		*/
	} else {
		passengersWrapper := struct {
			Passengers map[string]Passenger `json:"Passengers"`
		}{getPasseners()}
		json.NewEncoder(w).Encode(passengersWrapper)
		return
	}
}
func getPasseners() map[string]Passenger {
	results, err := db.Query("select * from Passenger")
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

/* func searchPassengerName(q string) (map[string]Passenger, bool) {
	db, err := sql.Open("mysql", "user:password@tcp(127.0.0.1:3306)/passenger_db")

	// handle error
	if err != nil {
		panic(err.Error())
	}

	defer db.Close()
	results, err := db.Query("select * from Passenger where lower(name) like lower(?)", "%"+q+"%") //db query-> returns 2 things, dg query row returns 1 thing only
	if err != nil {
		panic(err.Error())
	}
	var passengersDB map[string]Passenger = map[string]Passenger{}
	for results.Next() {
		var p Passenger
		var id string
		err = results.Scan(&id, &p.FirstName, &p.LastName, &p.Email, &p.Mobile)
		if err != nil {
			panic(err.Error())
		}
		passengersDB[id] = p
	}

	if len(passengers) == 0 {
		return passengers, false
	}
	return passengers, true
}
*/

func queryPassengers(query string) (map[string]Passenger, bool) {
	results, err := db.Query("select * from Passenger where lower(name) like lower(?)", "%"+query+"%")
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
