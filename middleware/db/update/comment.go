package update

import (
	"../tables"
	"github.com/adamsanghera/mysqlBus"
)

//ModifyComment ...
//allows user to modify existing comments
//definitely needs work
//this is mainly to introduce updates to our current code
func ModifyComment(comment tables.Comment, newComment string) error {
	stmt, err := mysqlBus.DB.Prepare("UPDATE Comment SET comment_text = ? WHERE id = ? AND username = ? AND timest = ?")
	if err != nil {
		return err
	}

	_, err = stmt.Exec(newComment, comment.ID, comment.Username, comment.Timestamp)
	if err != nil {
		return err
	}
	return nil

}
