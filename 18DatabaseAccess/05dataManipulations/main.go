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

	// INSERT OPERATION--------------------------------------------------------------
	stmtInsert, err := db.Prepare("INSERT INTO employees(name) VALUES(?)")
	if err != nil {
		errorHandler(err)
	}
	defer stmtInsert.Close()

	resInsert, err := stmtInsert.Exec("Shubham")
	if err != nil {
		errorHandler(err)
	}

	lastId, err := resInsert.LastInsertId()
	if err != nil {
		errorHandler(err)
	}

	rowInsert, err := resInsert.RowsAffected()
	if err != nil {
		errorHandler(err)
	}
	log.Printf("id = %d, affected = %d\n", lastId, rowInsert)

	//UPDATE OPERATION------------------------------------------------------------------
	updateName, updateId := "Kingshuk", 1010
	stmtUpdate, err := db.Prepare("UPDATE employees SET name=? WHERE id = ?")
	if err != nil {
		errorHandler(err)
	}
	defer stmtUpdate.Close()

	resUpdate, err := stmtUpdate.Exec(updateName, updateId)
	if err != nil {
		errorHandler(err)
	}

	rowUpdate, err := resUpdate.RowsAffected()
	if err != nil {
		errorHandler(err)
	}
	log.Printf("Updated successfully", rowUpdate, updateName, updateId)

	// DELETE OPERATION-----------------------------------------------------------------
	deletedId := 1000
	stmtDelete, err := db.Prepare("DELETE FROM employees WHERE id = ?")
	if err != nil {
		errorHandler(err)
	}
	defer stmtUpdate.Close()

	resDelete, err := stmtDelete.Exec(deletedId)
	if err != nil {
		errorHandler(err)
	}

	rowDelete, err := resDelete.RowsAffected()
	if err != nil {
		errorHandler(err)
	}
	log.Printf("Delete %d successfully, where id = %d\n", rowDelete, deletedId)

	// ------------------------------------------------------------------------------------
}

func errorHandler(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
