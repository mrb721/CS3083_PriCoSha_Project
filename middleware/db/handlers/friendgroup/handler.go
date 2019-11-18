package friendgroup

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"../../db/tables"
	"../session"
)

// In make, we need to create a file that matches our content.
// In get, we need to retrieve this file, in byte string format.
// In edit, we need to overwrite the previous file.
// In delete we need to delete the previous file.

// Handler responds to http requests about content.
func Handler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("HANDLING A FRIEND GROUP REQUEST")
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
	resp.update(false, []tables.Member{}, []tables.FriendGroup{}, json.NewDecoder(r.Body).Decode(&data))

	fmt.Println("Obtained following data: ")
	fmt.Printf("%+v\n", data)

	// Validate requestor token
	valid := true
	valid, err := session.Validate(data.User, data.Token)
	resp.update(false, []tables.Member{}, []tables.FriendGroup{}, err)

	if valid {
		fmt.Println("The token was valid")
		switch data.Intent {
		case "mk":
			fmt.Println("Trying to make a friend group")
			if data.User == data.Creator {
				err = make(data.Creator, data.GroupName, data.Description)
				resp.update(err == nil, []tables.Member{}, []tables.FriendGroup{}, err)
				err = join(data.User, data.GroupName, data.Creator)
				resp.update(err == nil, []tables.Member{}, []tables.FriendGroup{}, err)
			}
		case "rm":
			fmt.Println("Trying to delete a friend group")
			if data.User == data.Creator {
				err = delete(data.Creator, data.GroupName)
			}
			resp.update(err == nil, []tables.Member{}, []tables.FriendGroup{}, err)
		case "get":
			fmt.Println("Trying to get all members of a friend group")
			members, err := getMembers(data.GroupName, data.Creator)
			resp.update(err == nil, members, []tables.FriendGroup{}, err)
		case "join":
			fmt.Println("Trying to join a friend group")
			err = join(data.User, data.GroupName, data.Creator)
			resp.update(err == nil, []tables.Member{}, []tables.FriendGroup{}, err)
		case "leave":
			fmt.Println("Trying to leave a friend group")
			err = leave(data.User, data.GroupName, data.Creator)
			resp.update(err == nil, []tables.Member{}, []tables.FriendGroup{}, err)
		case "memof":
			fmt.Println("Trying to retrieve all friend groups this user is a member of")
			groups, err := getAllMemberOf(data.User)
			resp.update(err == nil, []tables.Member{}, groups, err)
		}

	}
}
