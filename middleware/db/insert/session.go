package insert

import (
	"../tables"
	"github.com/adamsanghera/mysqlBus"
)

func Session(sesh tables.Session) error {
	stmt, err := mysqlBus.DB.Prepare("INSERT INTO Session VALUES(?, ?, ?)")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(sesh.Username, sesh.Token, sesh.ExpirationTime)
	if err != nil {
		return err
	}
	return nil
}
