package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

const (
	db_user     = "root"
	db_password = "king"
	db_addr     = "127.0.0.1"
	db_db       = "learn"
)

func main() {
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:3306)/%s", db_user, db_password, db_addr, db_db))
	if err != nil {
		errorHandler(err)
	}
	defer db.Close()

	//validate the DSN
	if err = db.Ping(); err != nil {
		errorHandler(err)
	} else {
		fmt.Println("Connection sucessfull, you are ready to use!\n\n")
	}

	var name, city, dept string
	id := 2000
	err = db.QueryRow("SELECT emp_name,emp_city,emp_dep FROM employee WHERE emp_id = ?", id).Scan(&name, &city, &dept)
	// log.Print(err)
	if err != nil {
		errorHandler(err)
	}
	fmt.Printf("Name: %s, City: %s, Dept: %s\n", name, city, dept)
}

func errorHandler(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
