package search

import (
	"../../db/tables"
	"github.com/adamsanghera/mysqlBus"
)

func Username(request string) ([]tables.Person, error) {
	rows, err := mysqlBus.DB.Query("SELECT * FROM Person WHERE username COLLATE UTF8_GENERAL_CI LIKE '%?'", request)
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

//FirstName ...
//Search by user's first name
func FirstName(request string) ([]tables.Person, error) {
	rows, err := mysqlBus.DB.Query("SELECT * FROM Person WHERE first_name COLLATE UTF8_GENERAL_CI LIKE '%?'", request)
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

//LastName ...
//Search by user's last  name
func LastName(request string) ([]tables.Person, error) {
	rows, err := mysqlBus.DB.Query("SELECT * FROM Person WHERE last_name COLLATE UTF8_GENERAL_CI  LIKE '%?'", request)
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
