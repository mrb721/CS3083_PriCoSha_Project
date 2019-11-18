package search

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"../session"
)

// Handler responds to http requests about content.
func Handler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("HANDLING A SEARCH REQUEST")
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
	resp.update(false, json.NewDecoder(r.Body).Decode(&data), []userResult{}, []groupResult{})

	fmt.Println("Obtained following data: ")
	fmt.Printf("%+v\n", data)

	// Validate requestor token
	valid := true
	valid, err := session.Validate(data.User, data.Token)

	resp.update(false, err, nil, nil)

	if valid {
		switch data.Intent {
		case "usr":
			usrs, err := User(data.Term)
			resp.update(err == nil, err, usrs, nil)
		case "fg":
			fgs, err := Group(data.Term)
			resp.update(err == nil, err, nil, fgs)
		}

	} else {
		resp.update(false, errors.New("Token invalid"), []userResult{}, []groupResult{})
	}
}
