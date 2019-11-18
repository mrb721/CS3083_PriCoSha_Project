package delete

import (
	"../tables"
	"github.com/adamsanghera/mysqlBus"
)

//Tag ...
// Deletes a single tag from the database; returns nil if it works ok.
func Tag(tag tables.Tag) error {
	stmt, err := mysqlBus.DB.Prepare("DELETE FROM Tag WHERE id = ? AND username_tagger = ? AND username_taggee = ?")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(tag.ID, tag.UsernameTagger, tag.UsernameTaggee)
	if err != nil {
		return err
	}
	return nil
}
