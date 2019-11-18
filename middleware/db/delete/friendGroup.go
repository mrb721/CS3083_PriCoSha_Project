package delete

import (
	"../tables"
	"github.com/adamsanghera/mysqlBus"
)

//FriendGroup ...
// Deletes a single friend group from the database; returns nil if it works ok.
func FriendGroup(group tables.FriendGroup) error {
	stmt, err := mysqlBus.DB.Prepare("DELETE FROM FriendGroup WHERE group_name = ? AND username = ? ")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(group.GroupName, group.Username)
	if err != nil {
		return err
	}
	return nil
}
