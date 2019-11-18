package insert

import (
	"github.com/adamsanghera/mysqlBus"
	"../tables"
)

//InsertMember ...
// Adds a member of a FriendGroup to the database; returns nil if it works ok.
func Member(member tables.Member) error {
	stmt, err := mysqlBus.DB.Prepare("INSERT INTO Member VALUES(?, ?, ?)")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(member.Username, member.GroupName, member.UsernameCreator)
	if err != nil {
		return err
	}
	return nil
}
