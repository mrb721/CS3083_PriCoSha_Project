package delete

import (
	"../tables"
	"github.com/adamsanghera/mysqlBus"
)

//Member ...
// Deletes a single member from the database; returns nil if it works ok.
func Member(member tables.Member) error {
	stmt, err := mysqlBus.DB.Prepare("DELETE FROM Person WHERE username = ? AND group_name = ? AND username_creator = ?")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(member.Username, member.GroupName, member.UsernameCreator)
	if err != nil {
		return err
	}
	return nil
}
