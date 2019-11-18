package remove

import (
	"fmt"

	"../query"
	"github.com/adamsanghera/mysqlBus"
)

//Group deletes all members of a group, and all shares to the group, before deleting the group itself.
func Group(creator string, groupname string) error {
	// Get all the members in the group
	members, err := query.MembersByGroup(groupname, creator)
	if err != nil {
		return err
	}

	fmt.Println("Deleting the members of " + groupname + " prior to deleting the group itself")
	// Delete all the memberships!
	for m := range members {
		err = Member(members[m].Username, creator, groupname)
		if err != nil {
			return err
		}
	}

	// Get all the shares made to the group
	shares, err := query.SharesByGroup(groupname, creator)
	if err != nil {
		return err
	}

	fmt.Println("Deleting the shares to " + groupname + " prior to deleting the group itself")
	// Delete all the shares!
	for s := range shares {
		err = Member(members[s].Username, creator, groupname)
		if err != nil {
			return err
		}
	}

	// Finally delete the friend group
	fmt.Println("Deleting the friend group " + groupname)
	stmt, err := mysqlBus.DB.Prepare("DELETE FROM FriendGroup WHERE username = ? AND group_name = ?")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(creator, groupname)
	if err != nil {
		return err
	}
	return nil
}

//AllGroups ...
//Deletes member from a specific friend group
func AllGroups(username string) error {
	stmt, err := mysqlBus.DB.Prepare("DELETE FROM FriendGroup WHERE username = ?")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(username)
	if err != nil {
		return err
	}
	return nil
}
