package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type form struct {
	Username string `json:"Username"`
	Password string `json:"Password"`
	Fname    string `json:"Fname"`
	Lname    string `json:"Lname"`
}

/*
	This test:
	Registers three users, adamsans, cristiano, and sanders.
*/

func regPerson(username string, pass string, fname string, lname string) {
	challenge := form{
		Username: username,
		Password: pass,
		Fname:    fname,
		Lname:    lname,
	}
	j, err := json.Marshal(challenge)

	fmt.Println(string(j))

	if err != nil {
		panic("YIKES")
	}

	req, err := http.NewRequest("POST", "http://localhost:3000/register/user", bytes.NewBuffer(j))

	if err != nil {
		panic("OH MY GOD")
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)

	fmt.Println("Body: ", string(body))
}

func main() {
	regPerson("adamsans", "password", "adam", "sanghera")
	// regPerson("sanders", "password", "sanders", "sanghera")
	// regPerson("cristiano", "izthebest", "cristiano", "penaldo")
}
