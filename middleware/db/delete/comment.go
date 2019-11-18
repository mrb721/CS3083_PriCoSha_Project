package delete

import (
	"../tables"
	"github.com/adamsanghera/mysqlBus"
)

//Comment ...
// Deletes a single comment from the database; returns nil if it works ok.
func Comment(comment tables.Comment) error {
	stmt, err := mysqlBus.DB.Prepare("DELETE FROM Comment WHERE id = ? AND username = ? AND timest = ?")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(comment.ID, comment.Username, comment.Timestamp)
	if err != nil {
		return err
	}
	return nil
}
