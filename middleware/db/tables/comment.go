package tables

import "time"

//Comment ...
// Represents the Comment Table in mysql
type Comment struct {
	ID          int
	Username    string
	Timestamp   time.Time
	CommentText string
}
