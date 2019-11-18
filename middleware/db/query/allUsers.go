package query

import (
	"github.com/adamsanghera/mysqlBus"
	"../tables"
)

//AllUsers ...
// Returns a slice of all users
func AllUsers() ([]tables.Person, error) {
	rows, err := mysqlBus.DB.Query("SELECT * FROM Person")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	allUsers := []tables.Person{}

	for rows.Next() {
		var p1 tables.Person
		if err := rows.Scan(&p1.Username, &p1.HashedPassword, &p1.Salt, &p1.Fname, &p1.Lname); err != nil {
			return nil, err
		}

		allUsers = append(allUsers, p1)
	}
	return allUsers, nil
}
