package search

type form struct {
	// Fields related to the tag
	//GroupName string `json:"GroupName"`
	Term   string `json:"Term"`   //can either be username or user's name or group name
	Intent string `json:"Intent"` // Can be "usr" or "fg" for user and friendgroup

	// Fields related to the requestor
	Token string `json:"Token"`
	User  string `json:"User"`
}
