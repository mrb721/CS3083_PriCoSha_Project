package query

import (
	"github.com/adamsanghera/mysqlBus"
	"../tables"
)

//AllShares ...
// Returns a slice of all tags
func AllShares() ([]tables.Share, error) {
	rows, err := mysqlBus.DB.Query("SELECT * FROM Share")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	allShares := []tables.Share{}

	for rows.Next() {
		var s1 tables.Share
		if err := rows.Scan(&s1.ID, &s1.GroupName, &s1.Username); err != nil {
			return nil, err
		}

		allShares = append(allShares, s1)
	}
	return allShares, nil
}
