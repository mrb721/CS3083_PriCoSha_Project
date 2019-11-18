package file

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"../../db/update"
	"../session"
)

var path = "/files/static"

//Handler ...
//Handles request and stores content being sent to the database
func WriteHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("HANDLING A FILE REQUEST")
	// Setup the response
	fmt.Println("Received a content request.  Preparing response")
	resp := response{
		Successful: false,
		ErrMsg:     errors.New("Unknown failure"),
	}
	writer := json.NewEncoder(w)

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

	defer writer.Encode(resp)

	// Parse the request.
	fmt.Println("Parsing request")
	data := form{}
	resp.ErrMsg = json.NewDecoder(r.Body).Decode(&data)

	fmt.Println("Obtained following data: ")
	fmt.Printf("%+v\n", data)

	// Validate requestor token
	valid := true
	valid, resp.ErrMsg = session.Validate(data.User, data.Token)

	filepath = path + data.ID + data.ContentType

	err = update.ModifyContentFilePath(filepath)

	createErr = create(filepath)
	writeErr = write(filepath)

	if createErr != nil {
		resp.ErrMsg = createErr
	} else {
		resp.ErrMsg = writeErr
	}

}

//GetHandler ...
//handler function to retreive file
func GetHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Received a content request.  Preparing response")
	resp := response{
		Successful: false,
		ErrMsg:     errors.New("Unknown failure"),
	}
	writer := json.NewEncoder(w)
	defer writer.Encode(resp)

	// Parse the request.
	fmt.Println("Parsing request")
	data := form{}
	resp.ErrMsg = json.NewDecoder(r.Body).Decode(&data)

	fmt.Println("Obtained following data: ")
	fmt.Printf("%+v\n", data)

	// Validate requestor token
	valid := true
	valid, resp.ErrMsg = session.Validate(data.User, data.Token)

	filepath = path + data.ID + data.ContentType

	resp.ErrMsg = read(filepath)

}

//EditHandler ...
//handler function to edit file
func EditHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Received a content request.  Preparing response")
	resp := response{
		Successful: false,
		ErrMsg:     errors.New("Unknown failure"),
	}
	writer := json.NewEncoder(w)
	defer writer.Encode(resp)

	// Parse the request.
	fmt.Println("Parsing request")
	data := form{}
	resp.ErrMsg = json.NewDecoder(r.Body).Decode(&data)

	fmt.Println("Obtained following data: ")
	fmt.Printf("%+v\n", data)

	// Validate requestor token
	valid := true
	valid, resp.ErrMsg = session.Validate(data.User, data.Token)

	filepath = path + data.ID + data.ContentType

	modErr = update.ModifyContentFilePath(filepath)
	writeErr = write(filepath, data.Content)
	if modErr != nil {
		resp.ErrMsg = modErr
	} else {
		resp.ErrMsg = writeErr
	}

}

//EditHandler ...
//handler function to edit file
func deleteHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Received a content request.  Preparing response")
	resp := response{
		Successful: false,
		ErrMsg:     errors.New("Unknown failure"),
	}
	writer := json.NewEncoder(w)
	defer writer.Encode(resp)

	// Parse the request.
	fmt.Println("Parsing request")
	data := form{}
	resp.ErrMsg = json.NewDecoder(r.Body).Decode(&data)

	fmt.Println("Obtained following data: ")
	fmt.Printf("%+v\n", data)

	// Validate requestor token
	valid := true
	valid, resp.ErrMsg = session.Validate(data.User, data.Token)

	filepath = path + data.ID + data.ContentType

	modErr = update.ModifyContentFilePath("")
	deleteErr = delete(filepath)

	if modErr != nil {
		resp.ErrMsg = modErr
	} else {
		resp.ErrMsg = deleteErr
	}

}
