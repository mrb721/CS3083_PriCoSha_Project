package query

import (
	"../tables"
	"github.com/adamsanghera/mysqlBus"
)

// ContentByGroup returns a content table, given the content id, containing
// all content that has been shared to a given group.
func ContentByGroup(groupName string) ([]tables.Content, error) {
	rows, err := mysqlBus.DB.Query("select id, username, timest, file_path, content_name, public from Share NATURAL JOIN Content WHERE group_name=?", groupName)
	if err != nil {
		return []tables.Content{}, err
	}
	defer rows.Close()

	contents := []tables.Content{}

	// Get all the contents
	for rows.Next() {
		content := tables.Content{}
		if err := rows.Scan(&content.ID, &content.Username, &content.Timestamp, &content.FilePath, &content.ContentName, &content.Public); err != nil {
			return []tables.Content{}, err
		}
		contents = append(contents, content)
	}
	return contents, nil
}

func PublicContent() ([]tables.Content, error) {
	rows, err := mysqlBus.DB.Query("select id, username, timest, file_path, content_name, public FROM Content WHERE public=?", true)
	if err != nil {
		return []tables.Content{}, err
	}
	defer rows.Close()

	contents := []tables.Content{}

	// Get all the contents
	for rows.Next() {
		content := tables.Content{}
		if err := rows.Scan(&content.ID, &content.Username, &content.Timestamp, &content.FilePath, &content.ContentName, &content.Public); err != nil {
			return []tables.Content{}, err
		}
		contents = append(contents, content)
	}
	return contents, nil
}
