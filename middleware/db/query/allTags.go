package query

import (
	"github.com/adamsanghera/mysqlBus"
	"../tables"
)

//AllTags ...
// Returns a slice of all tags
func AllTags() ([]tables.Tag, error) {
	rows, err := mysqlBus.DB.Query("SELECT * FROM Tag")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	allTags := []tables.Tag{}

	for rows.Next() {
		var t1 tables.Tag
		if err := rows.Scan(&t1.ID, &t1.UsernameTagger, &t1.UsernameTaggee, &t1.Timestamp, &t1.Status); err != nil {
			return nil, err
		}

		allTags = append(allTags, t1)
	}
	return allTags, nil
}
