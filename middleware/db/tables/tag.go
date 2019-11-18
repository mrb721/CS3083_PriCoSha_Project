package tables

import "time"

//Tag ...
// Struct representing the Tag table
type Tag struct {
	ID             int
	UsernameTagger string
	UsernameTaggee string
	Timestamp      time.Time
	Status         bool
}
