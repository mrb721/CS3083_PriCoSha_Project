package login

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"../../db/query"
	"../session"
	"github.com/adamsanghera/hashing"
)

//Handler ...
//  Handles login requests
func Handler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("HANDLING A LOGIN REQUEST")
	// Setup the response
	resp := setupResp()

	if acrh, ok := r.Header["Access-Control-Request-Headers"]; ok {
		w.Header().Set("Access-Control-Allow-Headers", acrh[0])
	}
	w.Header().Set("Access-Control-Allow-Credentials", "True")
	if acao, ok := r.Header["Access-Control-Allow-Origin"]; ok {
		w.Header().Set("Access-Control-Allow-Origin", acao[0])
	} else {
		if _, oko := r.Header["Origin"]; oko {
			w.Header().Set("Access-Control-Allow-Origin", r.Header["Origin"][0])
		} else {
			w.Header().Set("Access-Control-Allow-Origin", "*")
		}
	}
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE")
	w.Header().Set("Connection", "Close")

	defer json.NewEncoder(w).Encode(resp)

	// Parse the request.
	data := form{}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&data)

	fmt.Println("Got our hands on some request: ")
	fmt.Printf("%v\n", data)

	resp.updateResp("", 0, err)

	// Get the hashedPass from the DB.
	hashedPass, salt, err := query.LoginInfo(data.Username)
	resp.updateResp("", 0, err)

	// Check if the challenge is valid
	if !hashing.IsValidChallenge(data.Password, salt, hashedPass) {
		fmt.Println("Challenge ain't valid")
		resp.updateResp("", 0, errors.New("Password does not match our records"))
	} else {
		// Create a new session token for the user.
		fmt.Println("Challenge sure is valid")
		resp.updateResp(session.Create(data.Username))
	}
}
