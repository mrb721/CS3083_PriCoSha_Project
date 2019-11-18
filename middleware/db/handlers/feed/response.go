package feed

import (
	"../../db/tables"
)

type response struct {
	Successful bool             `json:"Successful"`
	Posts      []tables.Content `json:"Posts"`
	ErrMsg     error            `json:"ErrMsg"`
}

func (r *response) update(s bool, p []tables.Content, e error) {
	r.Successful = s
	r.Posts = p
	r.ErrMsg = e

	if e != nil {
		panic(e)
	}
}
