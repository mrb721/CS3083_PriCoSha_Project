package register

import (
	"encoding/json"
	"io"
)

type form struct {
	Username string `json:"Username"`
	Password string `json:"Password"`
	Fname    string `json:"Fname"`
	Lname    string `json:"Lname"`
}

func parseRequest(rawData io.ReadCloser) (form, error) {
	var data form
	err := json.NewDecoder(rawData).Decode(&data)
	return data, err
}
