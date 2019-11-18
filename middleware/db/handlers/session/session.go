package session

import (
	"crypto/rand"
	"encoding/hex"
	"errors"
	"fmt"
	"time"

	"../../db/insert"
	"../../db/query"
	"../../db/tables"
	"../../db/update"
)

const (
	tokenLength    = 256
	expirationTime = time.Duration(time.Second * 1000)
)

func genToken() string {
	token := make([]byte, tokenLength)
	_, err := rand.Read(token)
	if err != nil {
		panic(err)
	}
	return string(hex.EncodeToString(token))
}

//Create ...
// Possibilities:
// 1. User has no entry in Session table
// 2. User has an entry in Session table
//   2.1 Session is still valid
//   2.2 Session is not valid
//
// 2.1 and 2.2 should be treated the same way. Update the existing token with a valid one.
// 1 can be solved by inserting a new, valid token.
func Create(uname string) (string, time.Duration, error) {
	fmt.Println("Trying to create a new session token")
	// Check to branch between 1 and 2.
	_, err := query.SessionToken(uname)

	fmt.Println("Found a session token")

	// If 1 or 2
	if err == nil || err.Error() == "no entry for username" {
		// Make a new session row
		newSesh := tables.Session{
			Username:       uname,
			Token:          genToken(),
			ExpirationTime: time.Now().Add(expirationTime),
		}
		// If 2, else 1
		if err == nil {
			fmt.Println("Updating session table")
			err = update.Session(newSesh)
		} else {
			fmt.Println("Inserting into session table")
			err = insert.Session(newSesh)
		}
		if err != nil { // If update or insert fails.
			return "", 0, err
		}
		return newSesh.Token, expirationTime, nil
	}
	// Should never really get here, barring disaster.

	fmt.Println("We have an error!")

	return "", 0, err
}

// Validate validates a token against the database.
func Validate(uname string, challengeToken string) (bool, error) {
	seshData, err := query.SessionToken(uname)
	if err != nil {
		return false, errors.New("Can't retrieve session info for " + uname)
	}

	if seshData.Token == challengeToken {
		if time.Now().After(seshData.ExpirationTime) {
			return false, errors.New("Expired token")
		}
		return true, nil
	}
	return false, errors.New("Bad token")
}

func Renew(uname string, challengeToken string) (string, time.Duration, error) {
	seshData, err := query.SessionToken(uname)
	if err != nil {
		return "", 0, errors.New("Can't retrieve session info")
	}

	if seshData.Token == challengeToken {
		if time.Now().After(seshData.ExpirationTime) {
			return "", 0, errors.New("Expired token")
		}
		return Create(uname)
	}

	return "", 0, errors.New("Bad token")
}
