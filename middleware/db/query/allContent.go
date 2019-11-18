package query

import (
	"../tables"
	"github.com/adamsanghera/mysqlBus"
)

//AllContent ...
// Returns a slice of all content items
func AllContent() ([]tables.Content, error) {
	rows, err := mysqlBus.DB.Query("SELECT * FROM Content")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	allContent := []tables.Content{}

	for rows.Next() {
		var c1 tables.Content
		if err := rows.Scan(&c1.ID, &c1.Username, &c1.Timestamp, &c1.FilePath, &c1.ContentName, &c1.Public); err != nil {
			return nil, err
		}

		allContent = append(allContent, c1)
	}
	return allContent, nil
}
