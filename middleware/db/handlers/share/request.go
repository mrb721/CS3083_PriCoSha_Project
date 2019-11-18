package share

type form struct {
	// Fields related to the share
	ID        int    `json:"ID"`
	GroupName string `json:"GroupName"`
	Creator   string `json:"Creator"`
	Intent    string `json:"Intent"` // Can be "rm" or "mk" for remove and make

	// Fields related to the requestor
	Token string `json:"Token"`
	User  string `json:"User"`
}
