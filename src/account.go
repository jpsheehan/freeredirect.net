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

func (a *Account) getDomains() []Domain {
	var (
		domains []Domain
		db      *sql.DB
		stmt    *sql.Stmt
		rows    *sql.Rows
		err     error
	)

	db = dbGet()
	stmt, err = db.Prepare("SELECT id, account, domain, url, active FROM domains WHERE account = ?")

	if err != nil {
		return domains
	}

	rows, err = stmt.Query(a.id)

	if err != nil {
		return domains
	}

	for rows.Next() {
		var domain Domain
		rows.Scan(&domain.id, &domain.account, &domain.domain, &domain.url, &domain.active)
		domains = append(domains, domain)
	}

	return domains
}
