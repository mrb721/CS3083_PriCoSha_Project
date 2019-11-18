package share

type response struct {
	Successful bool  `json:"Successful"`
	ErrMsg     error `json:"ErrMsg"`
}

func (r *response) update(s bool, e error) {
	r.Successful = s
	r.ErrMsg = e
	if e != nil {
		panic(e)
	}
}
