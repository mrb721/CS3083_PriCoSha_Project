package query

import (
	"../tables"
	"github.com/adamsanghera/mysqlBus"
)

// GroupsByMember returns all the groups associated with a given user
func GroupsByMember(user string) ([]tables.Member, error) {
	rows, err := mysqlBus.DB.Query("SELECT * FROM Member where username=?", user)
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
