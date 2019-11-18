package friendgroup

type form struct {
	// Fields related to the tag
	Creator     string `json:"Creator"`
	GroupName   string `json:"GroupName"`
	Description string `json:"Description"`
	Intent      string `json:"Intent"` //"mk" , "rm" for make and delete
	// Fields related to the requestor
	Token string `json:"Token"`
	User  string `json:"User"`
}
