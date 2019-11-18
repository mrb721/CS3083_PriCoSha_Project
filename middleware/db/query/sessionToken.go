package query

import (
	"errors"

	"../tables"
	"github.com/adamsanghera/mysqlBus"
)

func SessionToken(username string) (tables.Session, error) {
	rows, err := mysqlBus.DB.Query("SELECT * FROM Session WHERE Username=?", username)
	if err != nil {
		return tables.Session{}, err
	}
	defer rows.Close()

	seshData := tables.Session{}

	if rows.Next() {
		if err := rows.Scan(&seshData.Username, &seshData.Token, &seshData.ExpirationTime); err != nil {
			return tables.Session{}, err
		}
	} else {
		return tables.Session{}, errors.New("no entry for username")
	}
	if rows.Next() {
		return tables.Session{}, errors.New("duplicate entires for username" + username + "in database")
	}

	return seshData, nil
}
