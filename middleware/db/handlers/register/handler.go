package register

import (
	"encoding/json"
	"fmt"
	"net/http"

	"../../db/insert"
	"../../db/tables"
	"github.com/adamsanghera/hashing"
)

//Handler ...
//  Handles registration requests.
func Handler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("HANDLING A REGISTER REQUEST")
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

	// Parse the request, make sure it's A-OK
	data, err := parseRequest(r.Body)
	resp.updateResp(false, err)

	// TODO: Validate challenge

	// Hash the password, with a generated salt.
	salt, hashedPassword := hashing.WithNewSalt(data.Password)

	// Create a new Person row!
	newPerson := tables.Person{
		Username:       data.Username,
		HashedPassword: hashedPassword,
		Salt:           salt,
		Fname:          data.Fname,
		Lname:          data.Lname,
		ColorPalette:   "ffffff",
	}

	// Insert this row into our database, make sure we're good!
	err = insert.User(newPerson)
	resp.updateResp(err == nil, err)
}
