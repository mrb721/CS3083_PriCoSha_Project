package feed

import (
	"errors"

	//"../../db/methods"

	"../../db/tables"
	"fmt"
	"github.com/adamsanghera/mysqlBus"
)

//public ...
//retreives all public posts
func public() ([]tables.Content, error) {
	rows, err := mysqlBus.DB.Query("SELECT * FROM Content WHERE Public = 'True' ")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	allContent := []tables.Content{}

	for rows.Next() {
		var c1 tables.Content
		if err := rows.Scan(&c1.ID, &c1.Username, &c1.Timestamp, &c1.FilePath, &c1.ContentName, &c1.Public); err != nil {
			return nil, err
		}

		allContent = append(allContent, c1)
	}
	return allContent, nil
}

//friendgroup ...
//retreives all friend group specific posts
//the query selects all the content from groups the user is a part of via matching usernames to group names
func friendgroup(username string) ([]tables.Content, error) {
	rows, err := mysqlBus.DB.Query("SELECT * AS c FROM Content WHERE group_name IN (SELECT group_name FROM Member WHERE username = ?)", username)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	allContent := []tables.Content{}

	for rows.Next() {
		var c1 tables.Content
		if err := rows.Scan(&c1.ID, &c1.Username, &c1.Timestamp, &c1.FilePath, &c1.ContentName, &c1.Public); err != nil {
			return nil, err
		}

		allContent = append(allContent, c1)
	}
	return allContent, nil
}

//allPosts ...
//retreives all posts viewable by this specific user
func allPosts(username string) ([]tables.Content, error) {
	allContent := []tables.Content{}
	var errMess error
	pubPosts, pubErr := public()
	if pubErr != nil {
		allContent = append(allContent, pubPosts...)
	} else {
		errMess = pubErr
	}

	fgPosts, fgErr := friendgroup(username)
	if fgErr != nil {
		allContent = append(allContent, fgPosts...)
	}

	//allContent = append(allContent, fgPosts...)

	//feedContent := methods.RemoveDupPosts(allContent)

	err := ""
	if fgErr != nil {
		err += fgErr.Error()
	}
	if errMess != nil {
		err += errMess.Error()
	}
	fmt.Println("HERE IS WHAT WE HAVE: ")
	fmt.Println(allContent)
	return allContent, errors.New(err)

	//return feedContent, nil

}
