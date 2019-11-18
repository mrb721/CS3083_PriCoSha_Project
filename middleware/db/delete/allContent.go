package delete

import (
	"github.com/adamsanghera/mysqlBus"
)

//AllContent ...
// Deletes all content from the database; returns nil if it works ok.
func AllContent() error {
	stmt, err := mysqlBus.DB.Prepare("DELETE FROM Content")
	if err != nil {
		return err
	}
	_, err = stmt.Exec()
	if err != nil {
		return err
	}
	return nil
}
