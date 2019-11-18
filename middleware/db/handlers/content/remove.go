package content

import (
	"errors"
	"os"

	"../../db/delete"
	"../../db/query"
)

// remove deletes the content file, and removes the row from sql
func remove(theID int, deletor string) error {
	// Find the content
	markedContent, err := query.Content(theID)
	if err != nil {
		return err
	}

	if deletor == markedContent.Username {
		// Delete the file
		if err = os.Remove(markedContent.FilePath); err != nil {
			return err
		}

		// Delete the sql row (which also deletes all comments AND tags)
		return delete.Content(markedContent.ID)
	}
	return errors.New("User does not have the authority to delete")
}
