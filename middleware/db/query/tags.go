package query

import (
	"../tables"
	"github.com/adamsanghera/mysqlBus"
)

//TagsOnContent ...
// Returns a slice of all tags
func TagsOnContent(id int) ([]tables.Tag, error) {
	rows, err := mysqlBus.DB.Query("SELECT * FROM Tag WHERE id = ?", id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	allTags := []tables.Tag{}

	for rows.Next() {
		var t1 tables.Tag
		if err := rows.Scan(&t1.ID, &t1.UsernameTagger, &t1.UsernameTaggee, &t1.Timestamp, &t1.Status); err != nil {
			return nil, err
		}

		allTags = append(allTags, t1)
	}
	return allTags, nil
}

//TagsPerPerson ...
// Returns a slice of all tags
func TagsPerPerson(theTag tables.Tag) ([]tables.Tag, error) {
	rows, err := mysqlBus.DB.Query("SELECT * FROM Tag WHERE username_taggee = ? AND status = 'FALSE' ", theTag.UsernameTaggee, theTag.Status)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	allTags := []tables.Tag{}

	for rows.Next() {
		var t1 tables.Tag
		if err := rows.Scan(&t1.ID, &t1.UsernameTagger, &t1.UsernameTaggee, &t1.Timestamp, &t1.Status); err != nil {
			return nil, err
		}

		allTags = append(allTags, t1)
	}
	return allTags, nil
}

func TagsByID(id int) ([]tables.Tag, error) {
	rows, err := mysqlBus.DB.Query("SELECT * FROM Tag WHERE id=?", id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	allTags := []tables.Tag{}

	for rows.Next() {
		var t tables.Tag
		if err := rows.Scan(&t.ID, &t.UsernameTagger, &t.UsernameTaggee, &t.Timestamp, &t.Status); err != nil {
			return nil, err
		}

		allTags = append(allTags, t)
	}
	return allTags, nil
}
