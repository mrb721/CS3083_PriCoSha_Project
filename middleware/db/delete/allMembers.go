package delete

import (
	"github.com/adamsanghera/mysqlBus"
)

//AllMembers ...
// Deletes all members from the database; returns nil if it works ok.
func AllMembers() error {
	stmt, err := mysqlBus.DB.Prepare("DELETE FROM Member")
	if err != nil {
		return err
	}
	_, err = stmt.Exec()
	if err != nil {
		return err
	}
	return nil
}
