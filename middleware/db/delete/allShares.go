package delete

import (
	"github.com/adamsanghera/mysqlBus"
)

//AllShares ...
// Deletes all shares from the database; returns nil if it works ok.
func AllShares() error {
	stmt, err := mysqlBus.DB.Prepare("DELETE FROM Share")
	if err != nil {
		return err
	}
	_, err = stmt.Exec()
	if err != nil {
		return err
	}
	return nil
}
