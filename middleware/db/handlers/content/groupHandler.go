package content

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"../../db/query"
	"../../db/tables"
	"../session"
)

func GroupHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("HANDLING A CONTENT REQUEST")
	// Setup the response
	resp := &groupResponse{
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
	data := &groupForm{}
	resp.update(false, []rawContent{}, json.NewDecoder(r.Body).Decode(data))

	fmt.Println("Obtained following data: ")
	fmt.Printf("%+v\n", data)

	// Validate requestor token
	valid := true
	valid, err := session.Validate(data.User, data.Token)
	resp.update(false, []rawContent{}, err)

	if valid {
		fmt.Println("Was given a valid token for the requestor")
		contents := []tables.Content{}
		switch data.Intent {
		case "public":
			contents, err = getPublic()
		default:
			// Check to make sure user is in group
			groups, err := query.GroupsByMember(data.User)
			resp.update(false, []rawContent{}, err)

			validGroup := false
			for _, b := range groups {
				if b.GroupName == data.GroupName {
					validGroup = true
				}
			}
			if !validGroup {
				resp.update(false, []rawContent{}, errors.New("Tried to get contents from a group that the user was not a member of"))
			}
			contents, err = getByGroup(data.GroupName)
		}
		rawContents := []rawContent{}
		for c := range contents {
			r, rct, table, err := get(contents[c].ID)
			resp.update(false, []rawContent{}, err)
			rc := rawContent{
				Content:         r,
				ContentType:     rct,
				ContentName:     table.ContentName,
				ID:              table.ID,
				UsernameCreator: table.Username,
			}
			//fmt.Printf("%+v\n", rc)
			rawContents = append(rawContents, rc)
		}
		resp.update(err == nil, rawContents, err)
	}

}
