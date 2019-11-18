package share

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"../../db/query"

	"../session"
)

// Handler responds to http requests about content.
func Handler(w http.ResponseWriter, r *http.Request) {
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
	resp.update(false, json.NewDecoder(r.Body).Decode(&data))

	fmt.Println("Obtained following data: ")
	fmt.Printf("%+v\n", data)

	// Validate requestor token
	valid := true
	valid, err := session.Validate(data.User, data.Token)
	resp.update(false, err)

	// Get the Post we're trying to share
	postToShare, err := query.Content(data.ID)
	resp.update(false, err)

	// Get all the groups that the user is a member in.
	groups, err := query.GroupsByMember(data.User)
	resp.update(false, err)

	validGroup := false
	for _, b := range groups {
		if b.GroupName == data.GroupName {
			validGroup = true
		}
	}

	if !validGroup {
		resp.update(false, errors.New("Tried to share to a group that user is not a member of"))
	}

	if data.User != postToShare.Username {
		resp.update(false, errors.New("User does not have authority to share"))
	}

	if valid {
		switch data.Intent {
		case "mk":
			err = post(data.ID, data.GroupName, data.User)
		case "rm":
			err = remove(data.ID, data.GroupName, data.User)
		}
		resp.update(err == nil, err)
	} else {
		resp.update(false, errors.New("Token invalid"))
	}
}
