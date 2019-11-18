package feed

type form struct {
	// Fields related to the user
	// Intent string `json:"Intent"` // Can be "ret" or "get" to retreive or get files
	// Fields related to the requestor
	Token string `json:"Token"`
	User  string `json:"User"`
}
