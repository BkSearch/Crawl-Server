package main

import (
  "os"
  "log"
  "strconv"
)

var (
	port                                                           int
	host, userRead, userWrite, passwordRead, passwordWrite, dbName string
)

func loadConfig() {
	//load host
	host = os.Getenv("HOST")
	if host == "" {
		host = "localhost"
	}
	// load port
	port, _ = strconv.Atoi(os.Getenv("PORT"))
	if port == 0 {
		port = 5432
	}
	// read user
	userRead = os.Getenv("UserRead")
	if userRead == "" {
		log.Fatal("Invalid read user")
	}

	passwordRead = os.Getenv("PasswordRead")

	// write user
	userWrite = os.Getenv("UserWrite")
	if userWrite == "" {
		log.Fatal("Invalid write user")
	}
	passwordWrite = os.Getenv("PasswordWrite")
	// load db
	dbName = os.Getenv("DBName")
}

func main() {

}
