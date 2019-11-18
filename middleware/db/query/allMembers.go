package query

import (
	"github.com/adamsanghera/mysqlBus"
	"../tables"
)

//AllMembers ...
// Returns a slice of all members
func AllMembers() ([]tables.Member, error) {
	rows, err := mysqlBus.DB.Query("SELECT * FROM Member")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	allMembers := []tables.Member{}

	for rows.Next() {
		var m1 tables.Member
		if err := rows.Scan(&m1.Username, &m1.GroupName, &m1.UsernameCreator); err != nil {
			return nil, err
		}

		allMembers = append(allMembers, m1)
	}
	return allMembers, nil
}
