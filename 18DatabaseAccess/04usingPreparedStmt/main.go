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

	if err := db.Ping(); err != nil {
		errorHandler(err)
	}

	stmt, err := db.Prepare("INSERT INTO employees(name) VALUES(?)")
	if err != nil {
		errorHandler(err)
	}
	defer stmt.Close()

	_, err = stmt.Exec("Vikram") //andreas
	if err != nil {
		errorHandler(err)
	}

	res, err := stmt.Exec("Akash")
	if err != nil {
		errorHandler(err)
	}

	lastId, err := res.LastInsertId()
	if err != nil {
		errorHandler(err)
	}

	rowCount, err := res.RowsAffected()
	if err != nil {
		errorHandler(err)
	}
	log.Printf("id = %d, affected = %d\n", lastId, rowCount)
}

func errorHandler(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
