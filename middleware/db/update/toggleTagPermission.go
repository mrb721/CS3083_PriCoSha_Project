package update

import (
	"../tables"
	"github.com/adamsanghera/mysqlBus"
)

// Permission grants or denies permission of tag
func Permission(tag tables.Tag) error {
	stmt, err := mysqlBus.DB.Prepare("UPDATE Tag SET status=? WHERE id=? AND username_tagger=? AND username_taggee=?")
	if err != nil {
		return err
	}

	_, err = stmt.Exec(tag.Status, tag.ID, tag.UsernameTagger, tag.UsernameTaggee)
	if err != nil {
		return err
	}
	return nil

}
