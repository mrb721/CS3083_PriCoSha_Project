package content

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"

	"../session"
)

// In make, we need to create a file that matches our content.
// In get, we need to retrieve this file, in byte string format.
// In edit, we need to overwrite the previous file.
// In delete we need to delete the previous file.

// Handler responds to http requests about content.
func Handler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("HANDLING A CONTENT REQUEST")
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
	resp.update(false, []byte(""), "", "", 0, "", json.NewDecoder(r.Body).Decode(&data))

	fmt.Println("Obtained following data: ")
	fmt.Printf("%+v\n", data)

	// Validate requestor token
	valid := true
	valid, err := session.Validate(data.User, data.Token)
	resp.update(false, []byte(""), "", "", 0, "", err)

	if valid {
		fmt.Println("Was given a valid token for the requestor")
		switch data.Intent {
		case "mk":
			fmt.Println("Trying to post new content")
			err = make(data.User, time.Now(), data.ContentName, data.ContentType, data.Content)
			fmt.Println("Post attempt was made")
			resp.update(err == nil, []byte(""), "", "", 0, "", err)
		case "rm":
			fmt.Println("Trying to remove content")
			err = remove(data.ID, data.User)
			fmt.Println("Remove attempt was made")
			resp.update(err == nil, []byte(""), "", "", 0, "", err)
		case "ed":
			fmt.Println("Trying to edit content")
			err = edit(data.ID, data.User, data.Content, data.ContentName, data.ContentType)
			fmt.Println("Edit attempt was made")
			resp.update(err == nil, []byte(""), "", "", 0, "", err)
		case "get":
			content, cType, row, err := get(data.ID)
			resp.update(err == nil, content, row.ContentName, cType, row.ID, row.Username, err)
		case "publish":
			err = publish(data.ID, data.User)
			resp.update(err == nil, []byte(""), "", "", 0, "", err)
		case "privatize":
			err = privatize(data.ID, data.User)
			resp.update(err == nil, []byte(""), "", "", 0, "", err)
		}
	}
}
