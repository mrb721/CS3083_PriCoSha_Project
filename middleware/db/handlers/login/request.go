package login

type form struct {
	Username string `json:"Username"`
	Password string `json:"Password"`
}

// func parseRequest(req *http.Request) (*form, error) {
// 	var data form
// 	return &data, json.NewDecoder(req.Body).Decode(&data)
// }
