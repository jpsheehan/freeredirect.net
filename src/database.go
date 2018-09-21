package main

import (
	"database/sql"
	"fmt"
	"strings"

	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

func dbGet() *sql.DB {
	return db
}

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

	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS accounts (
			id INTEGER PRIMARY KEY,
			email CHAR(254) NOT NULL UNIQUE,
			name CHAR(32),
			verified BOOL NOT NULL
		)`)

	if err != nil {
		return err
	}

	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS domains (
			id INTEGER PRIMARY KEY,
			account INTEGER NOT NULL,
			domain CHAR(256) NOT NULL UNIQUE,
			url CHAR(1024) NOT NULL,
			active BOOL NOT NULL,
			FOREIGN KEY(account) REFERENCES accounts(id)
		)`)

	if err != nil {
		return err
	}

	return nil
}

func dbCreateTestData() error {

	var (
		err       error
		statement *sql.Stmt
	)

	statement, err = db.Prepare("INSERT INTO accounts (email, name, verified) VALUES (?, ?, true)")

	if err == nil {
		statement.Exec("jesse@sheehan.nz", "Jesse Sheehan")
		statement.Exec("yvette@bodykarma.nz", "Yvette Merrin")
		statement.Exec("mcg50@uclive.ac.nz", "Mickey Gallagher")
	}

	statement, err = db.Prepare(`
		INSERT INTO domains (
			account, domain, url, active
		) VALUES (
			?, ?, ?, true
		)`)

	if err == nil {
		statement.Exec(1, "google.local", "https://www.google.com/")
		statement.Exec(1, "maps.local", "https://www.google.com/")
		statement.Exec(1, "images.local", "https://www.google.com/")
		statement.Exec(1, "test.local", "https://www.google.com/")
		statement.Exec(2, "auracreative.nz", "https://www.facebook.com/AuraCreativeArt/")
		statement.Exec(3, "mclovin.co.nz", "https://www.linkedin.com/in/michael-gallagher-2a3538122/")
	}

	return nil
}

func getFallThroughURL(domain string) string {
	return fmt.Sprintf("https://freeredirect.net/#%s", domain)
}

func dbGetRedirectURL(domain string) string {

	var (
		statement  *sql.Stmt
		err        error
		rows       *sql.Rows
		defaultURL = getFallThroughURL(domain)
	)

	statement, err = db.Prepare(`SELECT url FROM domains WHERE domain = ? AND active = true LIMIT 1`)

	if err != nil {
		return defaultURL
	}

	rows, err = statement.Query(normalizeDomain(domain))

	if err != nil {
		return defaultURL
	}

	for rows.Next() {
		var url string
		rows.Scan(&url)
		return url
	}

	return defaultURL
}

func normalizeDomain(domain string) string {
	parts := strings.Split(domain, ".")
	if len(parts) > 1 {
		if strings.Compare(parts[0], "www") == 0 {
			return strings.Join(parts[1:], ".")
		}
	}
	return domain
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
