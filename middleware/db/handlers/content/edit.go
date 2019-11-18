package content

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"strings"

	"../../db/query"
	"../../db/update"
)

// edit deletes the old content, replacing it with new content
func edit(theID int, user string, c []byte, cName string, cType string) error {
	// Find the content
	fmt.Println(">> EDITING A POST")
	markedContent, err := query.Content(theID)
	if err != nil {
		return err
	}

	if markedContent.Username != user {
		return errors.New("User is not authorized to edit this content.")
	}

	// Update the file extension, if necessary
	oldType := path.Ext(markedContent.FilePath)
	if oldType != cType {
		fmt.Println(">> Updating a file path")
		markedContent.FilePath = markedContent.FilePath[:1+strings.LastIndex(markedContent.FilePath, ".")] + cType
	}

	// Delete the old file
	fmt.Println(">> Deleting the old file")
	if err = os.Remove(markedContent.FilePath); err != nil {
		return err
	}

	// Write the new file
	fmt.Println(">> Writing the new file")
	if err = ioutil.WriteFile(markedContent.FilePath, c, 777); err != nil {
		return err
	}

	// Update the sql row
	fmt.Println(">> Updating the SQL row")
	return update.Content(markedContent.ID, markedContent.FilePath, cName)
}
