package content

type form struct {
	// Fields related to the tag
	ID          int    `json:"ID"`
	Username    string `json:"Username"`
	ContentName string `json:"ContentName"`
	ContentType string `json:"ContentType"` // Can be "ret" or "get" to retreive or get files
	Content     []byte `json:"Content"`
	Intent      string `json:"Intent"`
	// Fields related to the requestor
	Token string `json:"Token"`
	User  string `json:"User"`
}

type groupForm struct {
	GroupName string `json:"GroupName"`
	Intent    string `json:"Intent"`

	// For validation
	Token string `json:"Token"`
	User  string `json:"User"`
}
