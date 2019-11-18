package query

import (
	"github.com/adamsanghera/mysqlBus"
	"../tables"
)

//AllFriendGroups ...
// Returns a slice of all friend groups
func AllFriendGroups() ([]tables.FriendGroup, error) {
	rows, err := mysqlBus.DB.Query("SELECT * FROM FriendGroup")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	allFriendGroups := []tables.FriendGroup{}

	for rows.Next() {
		var fg1 tables.FriendGroup
		if err := rows.Scan(&fg1.GroupName, &fg1.Username, &fg1.Description); err != nil {
			return nil, err
		}

		allFriendGroups = append(allFriendGroups, fg1)
	}
	return allFriendGroups, nil
}
