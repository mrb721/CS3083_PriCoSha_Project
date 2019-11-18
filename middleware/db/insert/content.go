package insert

import (
	"../tables"
	"github.com/adamsanghera/mysqlBus"
)

//InsertContent ...
// Adds a content item to the database; returns nil if it works ok.
func Content(content tables.Content) error {
	stmt, err := mysqlBus.DB.Prepare("INSERT INTO Content VALUES(?, ?, ?, ?, ?, ?)")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(content.ID, content.Username, content.Timestamp, content.FilePath, content.ContentName, 1)
	if err != nil {
		return err
	}
	return nil
}
