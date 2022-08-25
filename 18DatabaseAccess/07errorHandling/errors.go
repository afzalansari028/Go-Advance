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
	db_db       = "company"
)

func main() {
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:3306)/%s", db_user, db_password, db_addr, db_db))
	if err != nil {
		errorHandler(err)
	}
	defer db.Close()

	//generally test for errors anytime a function call return error
	if err = db.Ping(); err != nil {
		errorHandler(err)
	}

	var name, city, dept string
	id := 9999
	err = db.QueryRow("select name, city, dept from employees where id = ?", id).Scan(&name, &city, &dept)
	// test for errors related to performing QueryRow operation
	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Printf("Sorry, no rows found for id = %d\n", id)
		} else {
			log.Fatal(err)
		}
	} else {
		fmt.Printf("name: %s, city: %s,dept: %s\n\n", name, city, dept)
	}

}

func errorHandler(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
