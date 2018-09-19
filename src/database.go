package main

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

func dbConnect() error {

	var err error

	db, err = sql.Open("sqlite3", "./test.db")

	if err != nil {
		return err
	}

	return nil

}

func dbInitialize() error {

	var err error

	_, err = db.Exec("CREATE TABLE IF NOT EXISTS domains (id INTEGER PRIMARY KEY, account INT NOT NULL, domain CHAR(256) NOT NULL UNIQUE, url CHAR(1024) NOT NULL, active BOOL NOT NULL)")

	if err != nil {
		return err
	}

	_, err = db.Exec("CREATE TABLE IF NOT EXISTS accounts (id INTEGER PRIMARY KEY, email CHAR(254) NOT NULL UNIQUE, name CHAR(32), verified BOOL NO NULL)")

	if err != nil {
		return err
	}

	return nil
}

func dbCreateTestData() error {

	var (
		err error
	)

	_, err = db.Exec("INSERT INTO domains (account, domain, url, active) VALUES (1, 'sheehan.nz', 'https://facebook.com/jpsheehan', true)")

	if err != nil {
		fmt.Println("Could not insert test data into domains")
	}

	_, err = db.Exec("INSERT INTO accounts (email, name, verified) VALUES ('jesse@sheehan.nz', 'Jesse', false)")

	if err != nil {
		fmt.Println("Could not insert test data into accounts")
	}

	fmt.Println("Finished inserting test data")

	return nil
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	checkError(dbConnect())
	checkError(dbInitialize())
	checkError(dbCreateTestData())

}
