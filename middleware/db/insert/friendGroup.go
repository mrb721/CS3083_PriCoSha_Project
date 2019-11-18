package insert

import (
	"github.com/adamsanghera/mysqlBus"
	"../tables"
)

//InsertFriendGroup ...
// Adds a Friend Group to the database; returns nil if it works ok.
func FriendGroup(group tables.FriendGroup) error {
	stmt, err := mysqlBus.DB.Prepare("INSERT INTO FriendGroup VALUES(?, ?, ?)")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(group.GroupName, group.Username, group.Description)
	if err != nil {
		return err
	}
	return nil
}
