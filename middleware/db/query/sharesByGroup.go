package query

import (
	"../tables"
	"github.com/adamsanghera/mysqlBus"
)

func SharesByGroup(groupname string, creator string) ([]tables.Share, error) {
	rows, err := mysqlBus.DB.Query("SELECT * FROM Share WHERE group_name=? AND username=?", groupname, creator)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	shares := []tables.Share{}

	for rows.Next() {
		var s tables.Share
		if err := rows.Scan(&s.ID, &s.GroupName, &s.Username); err != nil {
			return nil, err
		}
		shares = append(shares, s)
	}
	return shares, nil
}
