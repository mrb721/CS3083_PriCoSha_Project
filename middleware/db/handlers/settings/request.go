package settings

type form struct {
	// Fields related to the tag
	Username     string `json:"Username"`
	ColorPalette string `json:"ColorPalette"`
	Intent       string `json:"Intent"` // Can be "ccp" or "gcp" or "del" for change color palette, get color palette and delete

	// Fields related to the requestor
	Token string `json:"Token"`
	User  string `json:"User"`
}
