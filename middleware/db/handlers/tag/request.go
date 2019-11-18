package tag

import (
	"encoding/json"
	"io"
)

type form struct {
	// Fields related to the tag
	ID             int    `json:"ID"`
	UsernameTagger string `json:"UsernameTagger"`
	UsernameTaggee string `json:"UsernameTaggee"`
	Intent         string `json:"Intent"` // Can be "rm" "mk", "ap" or "de" for remove, make, approve or deny

	// Fields related to the requestor
	Token string `json:"Token"`
	User  string `json:"User"`
}

func parseRequest(rawData io.ReadCloser) (form, error) {
	var data form
	err := json.NewDecoder(rawData).Decode(&data)
	return data, err
}
