package query

import (
	"../tables"
	"github.com/adamsanghera/mysqlBus"
)

// Content returns a content table, given the content id
func Content(id int) (tables.Content, error) {
	rows, err := mysqlBus.DB.Query("SELECT * FROM Content WHERE id=?", id)
	if err != nil {
		return tables.Content{}, err
	}
	defer rows.Close()

	c := tables.Content{}

	for rows.Next() {
		if err := rows.Scan(&c.ID, &c.Username, &c.Timestamp, &c.FilePath, &c.ContentName, &c.Public); err != nil {
			return tables.Content{}, err
		}
	}
	return c, nil
}
