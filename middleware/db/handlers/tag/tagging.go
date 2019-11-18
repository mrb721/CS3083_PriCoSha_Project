package tag

import (
	"fmt"
	"time"

	"../../db/delete"
	"../../db/insert"
	"../../db/query"
	"../../db/tables"
	"../../db/update"
)

//Add ...
//Allows user to tag other users in content
func add(id int, usernameTagger string, usernameTaggee string, status bool) error {
	newTag := tables.Tag{}
	newTag.UsernameTagger = usernameTagger
	newTag.UsernameTaggee = usernameTaggee
	newTag.Timestamp = time.Now()
	newTag.Status = status
	newTag.ID = id

	err := insert.Tag(newTag)

	return err
}

//Remove ...
//Allows user to remove posted content
func remove(theID int, usernameTagger string, usernameTaggee string) error {
	tag := tables.Tag{}
	tag.ID = theID
	tag.UsernameTagger = usernameTagger
	tag.UsernameTaggee = usernameTaggee
	return delete.Tag(tag)
}

//Permission ...
//Allows user to change status of tags referring to them
func Permission(id int, usernameTagger string, usernameTaggee string, approved bool) error {
	tag := tables.Tag{}
	tag.ID = id
	tag.UsernameTagger = usernameTagger
	tag.UsernameTaggee = usernameTaggee
	tag.Status = approved
	fmt.Println("NEXT IS THE TAG:", tag)
	return update.Permission(tag)

}

//get ...
//retreives tags
func get(id int) ([]tables.Tag, error) {

	return query.TagsOnContent(id)
}
