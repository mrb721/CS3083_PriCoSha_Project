package insert

import (
	"fmt"

	"../tables"
	"github.com/adamsanghera/mysqlBus"
)

//InsertComment ...
// Adds a user to the database; returns nil if it works ok.
func Comment(comment tables.Comment) error {
	stmt, err := mysqlBus.DB.Prepare("INSERT INTO Comment VALUES(?, ?, ?, ?)")
	if err != nil {
		return err
	}
	fmt.Println("Inserting a comment with ID " + string(comment.ID))
	_, err = stmt.Exec(comment.ID, comment.Username, comment.Timestamp, comment.CommentText)
	if err != nil {
		return err
	}
	return nil
}
