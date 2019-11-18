package settings

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"../session"
)

// Handler responds to http requests about content.
func Handler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("HANDLING A SETTINGS REQUEST")
	// Setup the response
	resp := &response{
		Successful:   false,
		ColorPalette: "000000",
		ErrMsg:       errors.New("Unknown failure"),
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
	resp.update(false, "000000", json.NewDecoder(r.Body).Decode(&data))

	fmt.Println("Obtained following data: ")
	fmt.Printf("%+v\n", data)

	// Validate requestor token
	valid := true
	valid, err := session.Validate(data.User, data.Token)

	resp.update(false, "000000", err)

	if valid {
		switch data.Intent {
		case "ccp":
			if data.User == data.Username {
				err = Set(data.Username, data.ColorPalette)
				resp.update(err == nil, data.ColorPalette, err)
			}
		case "gcp":
			if data.User == data.Username {
				col, err := Get(data.Username)
				resp.update(err == nil, col, err)
			}

		}
	} else {
		resp.update(false, "000000", errors.New("Token invalid"))
	}
}
