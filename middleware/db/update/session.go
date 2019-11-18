package update

import (
	"../tables"
	"github.com/adamsanghera/mysqlBus"
)

func Session(sesh tables.Session) error {
	stmt, err := mysqlBus.DB.Prepare("UPDATE Session SET token = ?, expirationTime = ? WHERE username = ?")
	if err != nil {
		return err
	}

	_, err = stmt.Exec(sesh.Token, sesh.ExpirationTime, sesh.Username)
	if err != nil {
		return err
	}
	return nil

}
