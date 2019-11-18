package query

import (
	"../tables"
	"github.com/adamsanghera/mysqlBus"
)

// CommentsByID returns every comment with the same ID
func CommentsByID(id int) ([]tables.Comment, error) {
	rows, err := mysqlBus.DB.Query("SELECT * FROM Comment WHERE id=?", id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	comments := []tables.Comment{}

	for rows.Next() {
		var c tables.Comment
		if err := rows.Scan(&c.ID, &c.Username, &c.Timestamp, &c.CommentText); err != nil {
			return nil, err
		}
		comments = append(comments, c)
	}
	return comments, nil
}
