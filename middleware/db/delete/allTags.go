package delete

import (
	"github.com/adamsanghera/mysqlBus"
)

//AllTags ...
// Deletes all tags from the database; returns nil if it works ok.
func AllTags() error {
	stmt, err := mysqlBus.DB.Prepare("DELETE FROM Tag")
	if err != nil {
		return err
	}
	_, err = stmt.Exec()
	if err != nil {
		return err
	}
	return nil
}
