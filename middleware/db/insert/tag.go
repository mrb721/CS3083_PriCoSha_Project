package insert

import (
	"../tables"
	"github.com/adamsanghera/mysqlBus"
)

//InsertTag ...
// Adds a tag to the database; returns nil if it works ok.
func Tag(tag tables.Tag) error {
	stmt, err := mysqlBus.DB.Prepare("INSERT INTO Tag VALUES(?, ?, ?, ?, ?)")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(tag.ID, tag.UsernameTagger, tag.UsernameTaggee, tag.Timestamp, tag.Status)
	if err != nil {
		return err
	}
	return nil
}
