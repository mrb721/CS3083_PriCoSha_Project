package query

import (
	"../tables"
	"github.com/adamsanghera/mysqlBus"
)

// User returns a user table, given the username
func User(username string) (tables.Person, error) {
	rows, err := mysqlBus.DB.Query("SELECT * FROM Person WHERE username =?", username)
	if err != nil {
		return tables.Person{}, err
	}
	defer rows.Close()

	p := tables.Person{}

	for rows.Next() {
		if err := rows.Scan(&p.Username, &p.HashedPassword, &p.Salt, &p.Fname, &p.Lname, &p.ColorPalette); err != nil {
			return tables.Person{}, err
		}
	}
	return p, nil
}
