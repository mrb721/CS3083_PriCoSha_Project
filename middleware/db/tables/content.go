package tables

import "time"

//Content ...
// Struct representing the SQL table Content
type Content struct {
	ID          int
	Username    string
	Timestamp   time.Time
	FilePath    string
	ContentName string
	Public      bool
}
