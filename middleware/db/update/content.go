package update

import (
	"github.com/adamsanghera/mysqlBus"
)

// Content updates the contentID, the file path, and the content name
func Content(id int, newFilePath string, newCName string) error {
	//if we choose to edit and include when it was edited, we can change the timestamp to the current one, to base future edits off
	stmt, err := mysqlBus.DB.Prepare("UPDATE Content SET file_path = ?, content_name = ? WHERE id = ?")
	if err != nil {
		return err
	}

	_, err = stmt.Exec(newFilePath, newCName, id)
	if err != nil {
		return err
	}
	return nil

}
