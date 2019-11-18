package content

import (
	"errors"
	"io/ioutil"
	"os"
	"path"

	"../../db/query"
	"../../db/tables"
)

// get returns content
func get(id int) ([]byte, string, tables.Content, error) {
	// 1. Get content filepath from sql.
	contentRow, err := query.Content(id)
	if err != nil {

		return nil, "", tables.Content{}, err
	}

	lookupPath := contentRow.FilePath

	// 2. Read content into byte string.
	if _, err := os.Stat(lookupPath); os.IsNotExist(err) {
		return nil, "", tables.Content{}, errors.New("File cannot be found")
	}
	c, err := ioutil.ReadFile(lookupPath)
	if err != nil {

		return nil, "", tables.Content{}, err
	}

	// 3. Return it all back

	return c, path.Ext(lookupPath), contentRow, nil
}

func getByGroup(GroupName string) ([]tables.Content, error) {
	return query.ContentByGroup(GroupName)
}

func getPublic() ([]tables.Content, error) {
	return query.PublicContent()
}
