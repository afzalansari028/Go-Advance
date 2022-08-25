package main

import (
	"fmt"

	_ "example.local/connection"
)

func main() {
	fmt.Println("Hi database...")
	connection.dbConnection()
}
