package content

import (
	"encoding/hex"
	"io/ioutil"
	"os"
	"time"

	"../../db/insert"
	"../../db/tables"
)

// make writes content to file, and then stores the metadata in sql
func make(username string, timestamp time.Time, contentName string, contentType string, content []byte) error {
	newPost := tables.Content{}
	newPost.Username = username
	newPost.Timestamp = timestamp
	newPost.ContentName = contentName
	ti, _ := timestamp.MarshalBinary()
	newPost.FilePath = "/files/static/" + username + "_" + hex.EncodeToString(ti) + "." + contentType

	if _, err := os.Stat("files/static/"); os.IsNotExist(err) {
		if err = os.MkdirAll("/files/static/", 777); err != nil {
			panic(err)
		}

	}

	err := ioutil.WriteFile(newPost.FilePath, content, 777)

	if err != nil {
		return err
	}

	err = insert.Content(newPost)
	return err
}
