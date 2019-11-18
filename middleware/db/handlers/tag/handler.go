package tag

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"../../db/tables"
	"../session"
)

//Handler ...
//Handles requests anc calls functions pertaining to tagging
func Handler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("HANDLING A REQUEST TO ADD A TAG TO A CONTENT")
	// Setup the response
	resp := &response{
		Successful: false,
		ErrMsg:     errors.New("Unknown error"),
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
	data, err := parseRequest(r.Body)
	resp.updateResp(false, []tables.Tag{}, err)

	// Validate requestor token
	valid, err := session.Validate(data.User, data.Token)
	resp.updateResp(false, []tables.Tag{}, err)

	if valid {
		fmt.Println("Token sent is valid for the requestor")
		switch data.Intent {
		case "mk":
			fmt.Println("Trying to add a tag")
			err = add(data.ID, data.UsernameTagger, data.UsernameTaggee, data.UsernameTaggee == data.UsernameTagger)
			resp.updateResp(err == nil, []tables.Tag{}, err)
		case "rm":
			fmt.Println("Trying to remove a tag")
			if data.User == data.UsernameTaggee || data.User == data.UsernameTagger {
				fmt.Println("The user has the right to remove")
				err = remove(data.ID, data.UsernameTagger, data.UsernameTaggee)
				resp.updateResp(err == nil, []tables.Tag{}, err)
			}
		case "ap":
			fmt.Println("Trying to approve a pending tag")
			if data.User == data.UsernameTaggee {
				fmt.Println("User has the right to approve")
				err = Permission(data.ID, data.UsernameTagger, data.UsernameTaggee, true)
				resp.updateResp(err == nil, []tables.Tag{}, err)
			}
		case "get":
			tags, err := get(data.ID)
			resp.updateResp(err == nil, tags, err)
		}
	} else {
		resp.updateResp(false, []tables.Tag{}, errors.New("Token invalid"))
	}
}
