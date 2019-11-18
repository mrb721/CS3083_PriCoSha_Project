package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

/*
	This test
	1. Logs in adamsans
	2. Gets all content from FriendGroup X, which adamsans has already created and joined.
*/

type loginResponse struct {
	Token          string        `json:"Token"`
	ExpirationTime time.Duration `json:"ExpirationTime"`
	ErrMsg         error         `json:"ErrMsg"`
}

type loginForm struct {
	Username string `json:"Username"`
	Password string `json:"Password"`
}

type rawContent struct {
	Content         []byte `json:"Content"`
	ContentName     string `json:"ContentName"`
	ContentType     string `json:"ContentType"`
	ID              int    `json:"ID"`
	UsernameCreator string `json:"UsernameCreator"`
}

type groupResponse struct {
	Successful  bool         `json:"Successful"`
	RawContents []rawContent `json:"RawContents"`
	ErrMsg      error        `json:"ErrMsg"`
}
type groupForm struct {
	GroupName string `json:"GroupName"`

	// For validation
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

	// SEND A REQUEST TO GET ALL THE CONTENTS FOR GROUP X

	contentRequest := groupForm{
		GroupName: "the ballerz",
		Token:     body.Token,
		User:      "adamsans",
	}

	j, err = json.Marshal(contentRequest)
	if err != nil {
		panic("YIKES")
	}

	req, err = http.NewRequest("POST", "http://localhost:3000/contentByGroup", bytes.NewBuffer(j))

	req.Header.Set("Content-Type", "application/json")

	resp, err = client.Do(req)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	bdy := &groupResponse{}

	err = json.NewDecoder(resp.Body).Decode(&bdy)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Response to content Request: %v\n", bdy)
}
