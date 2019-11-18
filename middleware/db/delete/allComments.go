package delete

import (
	"github.com/adamsanghera/mysqlBus"
)

//AllComments ...
// Deletes all comments from the database; returns nil if it works ok.
func AllComments() error {
	stmt, err := mysqlBus.DB.Prepare("DELETE FROM Comment")
	if err != nil {
		return err
	}
	_, err = stmt.Exec()
	if err != nil {
		return err
	}
	return nil
}
