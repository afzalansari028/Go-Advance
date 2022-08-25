package main

import (
	"database/sql"
	"fmt"
	"log"
	"strconv"

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
	// fmt.Printf("%s:%s@tcp(%s:3306)/%s", db_user, db_password, db_addr, db_db)
	if err != nil {
		errorHandler(err)
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		errorHandler(err)
	}

	//**********tx starts
	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}

	defer tx.Rollback()

	for i := 0; i < 10; i++ {
		name := "user" + strconv.Itoa(i)
		_, err = tx.Exec("INSERT INTO employees(name) VALUES (?)", name)

		//simulate a failer in traction
		// if i == 5 {
		// 	_, err = tx.Exec("INSERT INTO employess(name) VALUES (?)", name, "fail")
		// }
		if err != nil {
			tx.Rollback()
			fmt.Printf("error occured writing: %v\n", name)
			return
		}
	}
	if err := tx.Commit(); err != nil {
		log.Fatal(err)
	}

	emp, err := getEmployees(db)
	if err != nil {
		errorHandler(err)
	}
	for _, rs := range emp {
		fmt.Printf("%+v", rs)
	}
	fmt.Printf("We have %v people in out db!\n", len(emp))
}

func getEmployees(db *sql.DB) (emp []Employee, err error) {
	rs, err := db.Query("SELECT * FROM employees")
	if err != nil {
		return emp, err
	}
	defer rs.Close()

	var E Employee
	for rs.Next() {
		err = rs.Scan(&E.Id, &E.Name, &E.City, &E.Dept)
		if err != nil {
			return emp, err
		}
		emp = append(emp, E)
	}
	return emp, nil
}

//error hadler
func errorHandler(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
