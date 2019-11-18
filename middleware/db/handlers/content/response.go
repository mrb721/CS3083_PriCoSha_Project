package content

type response struct {
	Successful      bool   `json:"Successful"`
	Content         []byte `json:"Content"`
	ContentName     string `json:"ContentName"`
	ContentType     string `json:"ContentType"`
	ID              int    `json:"ID"`
	UsernameCreator string `json:"UsernameCreator"`
	ErrMsg          error  `json:"ErrMsg"`
}

type rawContent struct {
	Content         []byte `json:"Content"`
	ContentName     string `json:"ContentName"`
	ContentType     string `json:"ContentType"`
	ID              int    `json:"ID"`
	UsernameCreator string `json:"UsernameCreator"`
}

type groupResponse struct {
	Successful  bool         `json:"Successful"`
	RawContents []rawContent `json:"RawContents"`
	ErrMsg      error        `json:"ErrMsg"`
}

func (r *groupResponse) update(s bool, rc []rawContent, e error) {
	r.Successful = s
	r.RawContents = rc
	r.ErrMsg = e

	if e != nil {
		panic(e)
	}
}

func (r *response) update(s bool, c []byte, cName string, cType string, id int, uname string, e error) {
	r.Successful = s
	r.Content = c
	r.ContentName = cName
	r.ContentType = cType
	r.ID = id
	r.UsernameCreator = uname
	r.ErrMsg = e

	if e != nil {
		panic(e)
	}
}
