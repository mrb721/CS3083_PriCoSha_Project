package share

import (
	"../../db/delete"
	"../../db/insert"
	"../../db/tables"
)

//post ...
//allows for the sharing of a post to friend groups
func post(contentID int, groupName string, username string) error {
	newShare := tables.Share{}
	newShare.ID = contentID
	newShare.GroupName = groupName
	newShare.Username = username

	return insert.Share(newShare)

}

//remove ...
//allows for a share to be removed
func remove(contentID int, groupName string, username string) error {
	return delete.Share(contentID, groupName, username)
}
