package comment

import "../../db/tables"

type response struct {
	Successful bool             `json:"Successful"`
	ErrMsg     error            `json:"ErrMsg"`
	Comments   []tables.Comment `json:"Comments"`
}

func (r *response) update(s bool, t []tables.Comment, e error) {
	r.Successful = s
	r.ErrMsg = e
	r.Comments = t

	if e != nil {
		panic(e)
	}
}
