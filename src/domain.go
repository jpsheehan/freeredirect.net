package main

import (
	"database/sql"
	"fmt"
)

// Domain represents a single domain redirection
type Domain struct {
	id      int
	account int
	domain  string
	url     string
	active  bool
}

func (d *Domain) load(id int) error {

	var (
		db   *sql.DB
		stmt *sql.Stmt
		row  *sql.Row
		err  error
	)

	db = dbGet()

	stmt, err = db.Prepare("SELECT id, account, domain, url, active FROM domains WHERE id = ? LIMIT 1")

	if err != nil {
		return err
	}

	row = stmt.QueryRow(id)
	row.Scan(&d.id, &d.account, &d.domain, &d.url, &d.active)

	return nil
}

func (d *Domain) toString() string {
	return fmt.Sprintf("(%d, %d, %s, %s, %t)", d.id, d.account, d.domain, d.url, d.active)
}
