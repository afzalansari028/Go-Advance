package main

import (
	"database/sql"
	"fmt"
	"log"
	"reflect"

	_ "github.com/go-sql-driver/mysql"
)

const (
	db_user     = "root"
	db_password = "king"
	db_addr     = "127.0.0.1"
	db_db       = "company"
)

type Employee struct {
	Id   int
	Name string
	City string
	Dept string
}

func main() {
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:3306)/%s", db_user, db_password, db_addr, db_db))
	if err != nil {
		errorHandler(err)
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		errorHandler(err)
	}

	emp, err := getEmployees(db)
	if err != nil {
		errorHandler(err)
	}
	for _, rs := range emp {
		fmt.Printf("%+v\n", rs)
	}
	fmt.Printf("We have %v people in our db!\n", len(emp))
}

//get employees data from db
func getEmployees(db *sql.DB) (emp []Employee, err error) {
	rs, err := db.Query("SELECT * FROM employees")
	fmt.Println(reflect.TypeOf(rs))
	if err != nil {
		return emp, err
	}
	defer rs.Close()

	var E Employee
	for rs.Next() {
		err = rs.Scan(&E.Id, &E.Name, &E.City, &E.Dept)
		if err != nil {
			return emp, nil
		}
		emp = append(emp, E)
	}
	fmt.Println(reflect.TypeOf(emp))
	return emp, nil
}

// error hadler function
func errorHandler(err error) {
	log.Fatal(err)
}
