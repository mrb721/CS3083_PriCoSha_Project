package query

import (
	"errors"

	"github.com/adamsanghera/mysqlBus"
)

//LoginInfo ...
// Returns the hashed password and salt from the database.
func LoginInfo(username string) (string, string, error) {
	rows, err := mysqlBus.DB.Query("SELECT hashedPassword, salt  FROM Person WHERE username=?", username)
	if err != nil {
		return "", "", nil
	}
	if rows.Next() {
		var hashedPass string
		var salt string
		err = rows.Scan(&hashedPass, &salt)
		if rows.Next() {
			return "", "", errors.New("There's two users with the same username.  Stranger danger")
		}
		return hashedPass, salt, nil
	}
	return "", "", errors.New("No person with that username exists in the database")
}
