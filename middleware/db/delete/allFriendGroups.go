package delete

import (
	"github.com/adamsanghera/mysqlBus"
)

//AllFriendGroups ...
// Deletes all friend groups from the database; returns nil if it works ok.
func AllFriendGroups() error {
	stmt, err := mysqlBus.DB.Prepare("DELETE FROM FriendGroup")
	if err != nil {
		return err
	}
	_, err = stmt.Exec()
	if err != nil {
		return err
	}
	return nil
}
