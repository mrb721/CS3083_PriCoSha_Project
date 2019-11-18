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
	1. Logs in as adamsans
	2. Makes a request to post new content
*/

type loginResponse struct {
	Token          string        `json:"Token"`
	ExpirationTime time.Duration `json:"ExpirationTime"`
	ErrMsg         error         `json:"ErrMsg"`
}

type logInfo struct {
	Username string `json:"Username"`
	Password string `json:"Password"`
}

type form struct {
	// Fields related to the tag
	ID          int    `json:"ID"`
	Username    string `json:"Username"`
	ContentName string `json:"ContentName"`
	ContentType string `json:"ContentType"`
	Content     []byte `json:"Content"`
	Intent      string `json:"Intent"` // Can be "rm" "mk", or "ed" for remove, make, or edit

	// Fields related to the requestor
	Token string `json:"Token"`
	User  string `json:"User"`
}

func main() {
	// SEND A REQUEST TO LOGIN, NEEDED FOR TOKEN
	challenge := logInfo{
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

	// SEND A REQUEST TO POST A CONTENT

	content := form{
		ContentName: "My first post",
		ContentType: "txt",
		Content:     []byte("The quick fox jumped over the lazy dog"),
		Intent:      "mk",
		Token:       body.Token,
		User:        "adamsans",
	}

	j, err = json.Marshal(content)
	if err != nil {
		panic("YIKES")
	}

	req, err = http.NewRequest("POST", "http://localhost:3000/content", bytes.NewBuffer(j))

	req.Header.Set("Content-Type", "application/json")

	resp, err = client.Do(req)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	bdy, _ := ioutil.ReadAll(resp.Body)

	fmt.Println("Response to content post: ", string(bdy))

}
