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
	1. Logs in as cristiano
	2. Makes a request to leave a given friend group
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

type groupForm struct {
	// Fields related to the tag
	Creator     string `json:"Creator"`
	GroupName   string `json:"GroupName"`
	Description string `json:"Description"`
	Intent      string `json:"Intent"` //"mk" , "rm" for make and delete
	// Fields related to the requestor
	Token string `json:"Token"`
	User  string `json:"User"`
}

type groupResponse struct {
	Successful bool  `json:"Successful"`
	ErrMsg     error `json:"ErrMsg"`
}

func main() {

	// SEND A REQUEST TO LOGIN, NEEDED FOR TOKEN
	challenge := loginForm{
		Username: "cristiano",
		Password: "izthebest",
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

	// SEND A REQUEST TO LEAVE A FRIEND GROUP

	group := groupForm{
		Creator:   "adamsans",
		GroupName: "the ballerz",
		Intent:    "leave",
		Token:     body.Token,
		User:      "cristiano",
	}

	j, err = json.Marshal(group)
	if err != nil {
		panic("YIKES")
	}

	req, err = http.NewRequest("POST", "http://localhost:3000/group", bytes.NewBuffer(j))

	req.Header.Set("Content-Type", "application/json")

	resp, err = client.Do(req)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	bdy, _ := ioutil.ReadAll(resp.Body)

	fmt.Println("Response to join friend group request: ", string(bdy))
}
