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
	Mobile    int
	Email     string
	Password  string
}

/* type Driver struct { // map this type to the record in the table
	DriverID  string
	FirstName string
	LastName  string
	IdenNo    string
	CarNo     string
	Mobile    int
	Email     string
	Password  string
}
*/
/*
	 func DeleteRecord(db *sql.DB, ID int) {
	    query := fmt.Sprintf(
	        "DELETE FROM Persons WHERE ID='%d'", ID)
	    _, err := db.Query(query)
	    if err != nil {
	        panic(err.Error())
	    }
	}
*/
/* func EditRecord(db *sql.DB, ID string, FN string, LN string, Mobile int, Email string, Pw string) {
	query := fmt.Sprintf(
		"UPDATE Passenger SET FirstName='%s', LastName='%s', Mobile=%d,Email='%s',AccPassword='%s' WHERE ID=%s",
		FN, LN, Mobile, Email, Pw, ID)
	_, err := db.Query(query)
	if err != nil {
		panic(err.Error())
	}
}

func InsertRecord(db *sql.DB, ID string, FN string, LN string, Mobile int, Email string, Pw string) {
	query := fmt.Sprintf("INSERT INTO Passenger VALUES (%s, '%s', '%s', %d, '%s', '%s')",
		ID, FN, LN, Mobile, Email, Pw)

	_, err := db.Query(query)

	if err != nil {
		panic(err.Error())
	}
}

func GetRecords(db *sql.DB) {
	results, err := db.Query("Select * FROM account_db.Passenger")

	if err != nil {
		panic(err.Error())
	}

	for results.Next() {
		// map this type to the record in the table
		var passenger Passenger
		err = results.Scan(&passenger.ID, &passenger.FirstName,
			&passenger.LastName, &passenger.Mobile, &passenger.Email, &passenger.Password)
		if err != nil {
			panic(err.Error())
		}

		fmt.Println(passenger.ID, passenger.FirstName,
			passenger.LastName, passenger.Mobile, passenger.Email)
	}
}
*/
/*
// Driver

	func EditDriverRecord(db *sql.DB, ID string, FN string, LN string, IdenNo string, CarNo string, Mobile int, Email string, Pw string) {
		query := fmt.Sprintf(
			"UPDATE Driver SET FirstName='%s', LastName='%s',IdenNo='%s',CarNo='%s', Mobile=%d,Email='%s',AccPassword='%s' WHERE ID=%s",
			FN, LN, IdenNo, CarNo, Mobile, Email, Pw, ID)
		_, err := db.Query(query)
		if err != nil {
			panic(err.Error())
		}
	}

	func InsertDriverRecord(db *sql.DB, ID string, FN string, LN string, IdenNo string, CarNo string, Mobile int, Email string, Pw string) {
		query := fmt.Sprintf("INSERT INTO Driver VALUES (%s, '%s', '%s','%s', '%s', %d, '%s', '%s')",
			ID, FN, LN, IdenNo, CarNo, Mobile, Email, Pw)

		_, err := db.Query(query)

		if err != nil {
			panic(err.Error())
		}
	}

	func GetDriverRecords(db *sql.DB) {
		results, err := db.Query("Select * FROM account_db.Driver")

		if err != nil {
			panic(err.Error())
		}

		for results.Next() {
			// map this type to the record in the table
			var driver Driver
			err = results.Scan(&driver.DriverID, &driver.FirstName,
				&driver.LastName, &driver.IdenNo, &driver.CarNo, &driver.Mobile, &driver.Email, &driver.Password)
			if err != nil {
				panic(err.Error())
			}

			fmt.Println(driver.DriverID, driver.FirstName,
				driver.LastName, driver.Mobile, driver.Email, driver.IdenNo, driver.CarNo)
		}
	}
*/
func main() {
	/* 	fmt.Println("Go MySQL Tutorial")
	   	db, err := sql.Open("mysql", "user:password@tcp(127.0.0.1:3306)/account_db")

	   	// handle error
	   	if err != nil {
	   		panic(err.Error())
	   	} else {
	   		fmt.Println("Database opened")
	   	}

	   	// InsertRecord(db, 2, "Wallace","Tan", 55)
	   	//InsertRecord("P0001", "Jake", "Lee", 99991111,'jakelee@gmail.com','password123');
	   	//EditRecord(db, 2, "Taylor", "Swift", 23)
	   	//DeleteRecord(db, 2)

	   	GetRecords(db)
	   	//GetDriverRecords(db)
	   	// defer the close till after the main function has finished executing
	   	defer db.Close() */
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
					updatePassenger(params["passengerid"])
					w.WriteHeader(http.StatusAccepted)
				} else {
					w.WriteHeader(http.StatusNotFound)
					fmt.Fprintf(w, "passenger ID does not exist")
				}
			} else {
				fmt.Println(err)
			}
		}
		/* else if r.Method == "PATCH" {
			if body, err := ioutil.ReadAll(r.Body); err == nil {
				var data map[string]interface{}

				if err := json.Unmarshal(body, &data); err == nil {
					if orig, ok := courses[params["courseid"]]; ok {
						fmt.Println(data)

						for k, v := range data {
							switch k {
							case "Name":
								orig.Name = v.(string)
							case "Planned Intake":
								orig.Intake = int(v.(float64))
							case "Min GPA":
								orig.MinGPA = int(v.(float64))
							case "Max GPA":
								orig.MaxGPA = int(v.(float64))
							}
						}
						courses[params["courseid"]] = orig
						w.WriteHeader(http.StatusAccepted)
					} else {
						w.WriteHeader(http.StatusNotFound)
						fmt.Fprintf(w, "Course ID does not exist")
					}
				} else {
					fmt.Println(err)
				}
			}
		} */

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
