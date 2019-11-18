package comment

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"

	"../../db/tables"
	"../session"
)

//Handler ...
//Handles requests anc calls functions pertaining to tagging
func Handler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("HANDLING A COMMENT REQUEST")

	// Setup the response
	resp := &response{
		Successful: false,
		ErrMsg:     errors.New("Unknown failure"),
	}

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
	fmt.Println("Parsing request")
	data := form{}
	resp.update(false, []tables.Comment{}, json.NewDecoder(r.Body).Decode(&data))

	fmt.Println("Obtained following data: ")
	fmt.Printf("%+v\n", data)

	// Validate requestor token
	valid := true
	valid, err := session.Validate(data.User, data.Token)
	resp.update(false, []tables.Comment{}, err)

	if valid {
		fmt.Println("Token for given user is valid")
		// err = errors.New("Invalid intent")
		switch data.Intent {
		case "mk":
			fmt.Println("Trying to make a new comment")
			err = post(data.ID, data.User, time.Now(), data.CommentText)
			resp.update(err == nil, []tables.Comment{}, err)
		case "rm":
			fmt.Println("Trying to remove a comment")
			err = remove(data.ID, data.User)
			resp.update(err == nil, []tables.Comment{}, err)
		case "ed":
			fmt.Println("Trying to edit a comment")
			err = edit(data.ID, data.User, data.Timestamp, data.CommentText)
			resp.update(err == nil, []tables.Comment{}, err)
		case "get":
			fmt.Println("Trying to get a comment")
			comments, err := get(data.ID)
			resp.update(err == nil, comments, err)
		}
	}
}
