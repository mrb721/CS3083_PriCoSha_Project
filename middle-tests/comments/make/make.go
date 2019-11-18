package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

/*
	This test:
	1. Logs in adamsans
	2. Makes a comment on post with content ID 1
*/

type Comment struct {
	ID          int
	Username    string
	Timestamp   time.Time
	CommentText string
}

type loginResponse struct {
	Token          string        `json:"Token"`
	ExpirationTime time.Duration `json:"ExpirationTime"`
	ErrMsg         error         `json:"ErrMsg"`
}

type loginForm struct {
	Username string `json:"Username"`
	Password string `json:"Password"`
}

type response struct {
	Successful bool      `json:"Successful"`
	ErrMsg     error     `json:"ErrMsg"`
	Comments   []Comment `json:"Comments"`
}

type commentForm struct {
	// Fields related to the tag
	ID          int    `json:"ID"`
	CommentText string `json:"CommentText"`
	Intent      string `json:"Intent"` // Can be "rm" "mk", or "ed"  for remove, make, and edit

	// Fields related to the requestor
	Token string `json:"Token"`
	User  string `json:"User"`
}

func main() {
	// SEND A REQUEST TO LOGIN, NEEDED FOR TOKEN
	challenge := loginForm{
		Username: "adamsans",
		Password: "password",
	}
	j, err := json.Marshal(challenge)

	fmt.Println(string(j))

	if err != nil {
		panic("YIKES")
	}

	req, err := http.NewRequest("POST", "http://localhost:3000/login/user", bytes.NewBuffer(j))

	if err != nil {
		panic("OH MY GOD")
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	body := loginResponse{}
	err = json.NewDecoder(resp.Body).Decode(&body)
	if err != nil {
		panic(err)
	}

	// SEND A REQUEST TO MAKE A COMMENT ON THE CONTENT

	comment := commentForm{
		ID:          1,
		CommentText: "Ah man this is cool bro gj",
		Intent:      "mk",
		Token:       body.Token,
		User:        "adamsans",
	}

	j, err = json.Marshal(comment)
	if err != nil {
		panic("YIKES")
	}

	req, err = http.NewRequest("POST", "http://localhost:3000/comment", bytes.NewBuffer(j))

	req.Header.Set("Content-Type", "application/json")

	resp, err = client.Do(req)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	bdy, _ := ioutil.ReadAll(resp.Body)

	fmt.Println("Response to content edit: ", string(bdy))
}
