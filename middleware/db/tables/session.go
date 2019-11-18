package tables

import "time"

//Session ...
// Struct implementing the Session table
// Represents a login session of a user
type Session struct {
	Username       string
	Token          string
	ExpirationTime time.Time
}
