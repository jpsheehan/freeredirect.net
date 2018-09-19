package main

import "database/sql"
import _ "github.com/mattn/go-sqlite3"

func database() error {
	db, err := sql.Open("sqlite3", "./test.db")

	if err != nil {
		return err
	}

	_, err = db.Exec("CREATE TABLE IF NOT EXISTS domains (id INT PRIMARY KEY NOT NULL, account INT NOT NULL, domain CHAR(256) NOT NULL, url CHAR(1024) NOT NULL, active BOOL NOT NULL)")

	if err != nil {
		return err
	}

	_, err = db.Exec("INSERT INTO domains (account, domain, url, active) VALUES (20, 'sheehan.nz', 'https://facebook.com/jpsheehan', true)")

	return nil
}

func main() {
	err := database()

	if err != nil {
		panic(err)
	}
}
