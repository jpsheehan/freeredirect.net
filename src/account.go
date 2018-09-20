package main

import (
	"database/sql"
	"fmt"
)

// Account represents a user account
type Account struct {
	id       int
	email    string
	name     string
	verified bool
}

func (a *Account) load(id int) error {

	var (
		db   *sql.DB
		stmt *sql.Stmt
		row  *sql.Row
		err  error
	)

	db = dbGet()

	stmt, err = db.Prepare("SELECT id, email, name, verified FROM accounts WHERE id = ? LIMIT 1")

	if err != nil {
		return err
	}

	row = stmt.QueryRow(id)

	row.Scan(&a.id, &a.email, &a.name, &a.verified)

	return nil
}

func (a *Account) toString() string {
	return fmt.Sprintf("(%d, %s, %s, %t)", a.id, a.email, a.name, a.verified)
}
