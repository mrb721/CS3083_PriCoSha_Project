package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type loginResponse struct {
	Token          string        `json:"Token"`
	ExpirationTime time.Duration `json:"ExpirationTime"`
	ErrMsg         error         `json:"ErrMsg"`
}

type loginForm struct {
	Username string `json:"Username"`
	Password string `json:"Password"`
}

type tagResponse struct {
	Successful bool  `json:"Successful"`
	ErrMsg     error `json:"ErrMsg"`
}

type tagForm struct {
	// Fields related to the tag
	ID             int    `json:"ID"`
	UsernameTagger string `json:"UsernameTagger"`
	UsernameTaggee string `json:"UsernameTaggee"`
	Intent         string `json:"Intent"` // Can be "rm" "mk", "ap" or "de" for remove, make, approve or deny

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

	// SEND A REQUEST TO TAG A USER

	tag := tagForm{
		ID:             4,
		UsernameTagger: "adamsans",
		UsernameTaggee: "cristiano",
		Intent:         "mk",
		Token:          body.Token,
		User:           "adamsans",
	}

	j, err = json.Marshal(tag)
	if err != nil {
		panic("YIKES")
	}

	req, err = http.NewRequest("POST", "http://localhost:3000/tag", bytes.NewBuffer(j))

	req.Header.Set("Content-Type", "application/json")

	resp, err = client.Do(req)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	bdy, _ := ioutil.ReadAll(resp.Body)

	fmt.Println("Response to tag make: ", string(bdy))
}
