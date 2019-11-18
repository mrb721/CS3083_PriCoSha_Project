package query

import (
	"../tables"
	"github.com/adamsanghera/mysqlBus"
)

// MembersByGroup returns all the members with a given the primary key of the table.
func MembersByGroup(groupName string, userCreator string) ([]tables.Member, error) {
	rows, err := mysqlBus.DB.Query("SELECT * FROM Member where group_name=? and username_creator=?", groupName, userCreator)
	if err != nil {
		return []tables.Member{}, err
	}
	defer rows.Close()

	members := []tables.Member{}

	for rows.Next() {
		m := tables.Member{}
		if err := rows.Scan(&m.Username, &m.GroupName, &m.UsernameCreator); err != nil {
			return []tables.Member{}, err
		}
		members = append(members, m)
	}
	return members, nil
}
