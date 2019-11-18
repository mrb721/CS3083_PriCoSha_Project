package search

type userResult struct {
	Username  string `json:"Username"`
	FirstName string `json:"FirstName"`
	LastName  string `json:"LastName"`
}

type groupResult struct {
	Creator     string `json:"Creator"`
	GroupName   string `json:"GroupName"`
	Description string `json:"Description"`
}

type response struct {
	Successful   bool          `json:"Successful"`
	ErrMsg       error         `json:"ErrMsg"`
	UserResults  []userResult  `json:"UserResults"`
	GroupResults []groupResult `json:"GroupResults"`
}

func (r *response) update(s bool, e error, usrs []userResult, groups []groupResult) {
	r.Successful = s
	r.ErrMsg = e
	r.UserResults = usrs
	r.GroupResults = groups
	if e != nil {
		panic(e)
	}
}
