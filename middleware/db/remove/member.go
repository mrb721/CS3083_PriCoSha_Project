package remove

import (
	"fmt"

	"github.com/adamsanghera/mysqlBus"
)

//Member ...
//Deletes member from a specific friend group
func Member(username string, creator string, groupname string) error {
	fmt.Println("Trying to delete " + username + " from " + groupname + ", a group created by " + creator)
	stmt, err := mysqlBus.DB.Prepare("DELETE FROM Member WHERE username = ? AND username_creator = ? AND group_name = ?")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(username, creator, groupname)
	if err != nil {
		return err
	}
	return nil
}

//MemberFromAllGroups ...
//Deletes member from a specific friend group
func MemberFromAllGroups(username string) error {
	stmt, err := mysqlBus.DB.Prepare("DELETE FROM Member WHERE username = ?")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(username)
	if err != nil {
		return err
	}
	return nil
}
