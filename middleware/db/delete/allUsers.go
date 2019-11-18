package delete

import (
	"github.com/adamsanghera/mysqlBus"
)

//AllUsers ...
// Deletes all users from the database; returns nil if it works ok.
func AllUsers() error {
	stmt, err := mysqlBus.DB.Prepare("DELETE FROM Person")
	if err != nil {
		return err
	}
	_, err = stmt.Exec()
	if err != nil {
		return err
	}
	return nil
}
