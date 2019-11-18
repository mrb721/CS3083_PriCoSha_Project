package delete

import (
	"github.com/adamsanghera/mysqlBus"
)

//Share ...
// Deletes a single share from the database; returns nil if it works ok.
func Share(id int, gName string, uName string) error {
	stmt, err := mysqlBus.DB.Prepare("DELETE FROM Share WHERE id=? AND group_name=? AND username=?")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(id, gName, uName)
	if err != nil {
		return err
	}
	return nil
}
