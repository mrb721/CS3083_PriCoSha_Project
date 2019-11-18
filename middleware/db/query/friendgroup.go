package query

import (
	"../tables"
	"github.com/adamsanghera/mysqlBus"
)

// Group returns a group table, given the group name
func Group(username string, groupName string) (tables.FriendGroup, error) {
	rows, err := mysqlBus.DB.Query("SELECT * FROM FriendGroup WHERE username =? AND group_name = ?", username, groupName)
	if err != nil {
		return tables.FriendGroup{}, err
	}
	defer rows.Close()

	fg := tables.FriendGroup{}

	for rows.Next() {
		if err := rows.Scan(&fg.Username, &fg.GroupName, &fg.Description); err != nil {
			return tables.FriendGroup{}, err
		}
	}
	return fg, nil
}

// AllJoinedGroups  returns all the groups a member is part of
func AllJoinedGroups(username string) ([]tables.FriendGroup, error) {
	rows, err := mysqlBus.DB.Query("SELECT * FROM FriendGroup WHERE group_name IN (SELECT group_name FROM Member WHERE username=?)", username)
	if err != nil {
		return []tables.FriendGroup{}, err
	}
	defer rows.Close()

	groups := []tables.FriendGroup{}

	for rows.Next() {
		fg := tables.FriendGroup{}
		if err := rows.Scan(&fg.Username, &fg.GroupName, &fg.Description); err != nil {
			return []tables.FriendGroup{}, err
		}
		groups = append(groups, fg)
	}
	return groups, nil
}
