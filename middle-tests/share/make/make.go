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
	2. Shares a given content to a given group
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

type shareForm struct {
	// Fields related to the share
	ID        int    `json:"ID"`
	GroupName string `json:"GroupName"`
	Creator   string `json:"Creator"`
	Intent    string `json:"Intent"` // Can be "rm" or "mk" for remove and make

	// Fields related to the requestor
	Token string `json:"Token"`
	User  string `json:"User"`
}

type shareResponse struct {
	Successful bool  `json:"Successful"`
	ErrMsg     error `json:"ErrMsg"`
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

	bodyLogin := loginResponse{}
	err = json.NewDecoder(resp.Body).Decode(&bodyLogin)
	if err != nil {
		panic(err)
	}

	// SEND A REQUEST TO SHARE THE POST

	share := shareForm{
		ID:        4,
		Creator:   "adamsans",
		GroupName: "the ballerz",
		Intent:    "mk",
		Token:     bodyLogin.Token,
		User:      "adamsans",
	}

	j, err = json.Marshal(share)
	if err != nil {
		panic("YIKES")
	}

	req, err = http.NewRequest("POST", "http://localhost:3000/share", bytes.NewBuffer(j))

	req.Header.Set("Content-Type", "application/json")

	resp, err = client.Do(req)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	bdy, _ := ioutil.ReadAll(resp.Body)

	fmt.Println("Response to content edit: ", string(bdy))
}
