package comment

import "time"

type form struct {
	// Fields related to the tag
	ID          int       `json:"ID"`
	CommentText string    `json:"CommentText"`
	Intent      string    `json:"Intent"` // 'mk' 'rm' 'ed' for make, remove, edit
	Timestamp   time.Time `json:"Timestamp"`

	// Fields related to the requestor
	Token string `json:"Token"`
	User  string `json:"User"`
}
