package comment

import (
	"time"

	"../../db/delete"
	"../../db/insert"
	"../../db/query"
	"../../db/tables"
	"../../db/update"
	//"time"
)

// post inserts a new comment into SQL
func post(id int, username string, timestamp time.Time, commentText string) error {
	newComm := tables.Comment{}
	newComm.ID = id
	newComm.Username = username
	newComm.Timestamp = timestamp
	//newComm.Timestamp = time.Now()      //Use this to set the timestamp at this point
	newComm.CommentText = commentText
	err := insert.Comment(newComm)

	return err
}

//remove ...
//Allowsfor the removal of posted comment
func remove(theID int, username string) error {
	post := tables.Comment{
		Username: username,
		ID:       theID,
	}
	return delete.Comment(post)
}

//edit ...
//Allows user to edit comments they made
func edit(id int, username string, timestamp time.Time, newComment string) error {
	comm := tables.Comment{}
	comm.ID = id
	comm.Username = username
	comm.Timestamp = timestamp

	return update.ModifyComment(comm, newComment)

}

func get(id int) ([]tables.Comment, error) {
	return query.CommentsByID(id)
}
