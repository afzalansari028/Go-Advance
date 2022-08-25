package connection

import (
	_ "github.com/go-sql-driver/mysql"
)

func Db("hi..........!")

// const (
// 	db_user     = "root"
// 	db_password = "king"
// 	db_addr     = "127.0.0.1"
// 	db_db       = "company"
// )

// func dbConnection() {
// 	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:3306)/%s", db_user, db_password, db_addr, db_db))
// 	fmt.Printf("%s:%s@tcp(%s:3306)/%s\n", db_user, db_password, db_addr, db_db)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer db.Close()

// 	//db settings to consider
// 	db.SetConnMaxIdleTime(time.Minute * 3)
// 	db.SetMaxOpenConns(10)
// 	db.SetMaxIdleConns(10)

// 	//validate the DSN
// 	if err := db.Ping(); err != nil {
// 		log.Fatal(err)
// 	} else {
// 		fmt.Println("connection successfull, you are ready to Go!")
// 	}
// }
