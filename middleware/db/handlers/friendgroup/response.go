package friendgroup

import "../../db/tables"

type response struct {
	Successful bool                 `json:"Successful"`
	ErrMsg     error                `json:"ErrMsg"`
	Members    []tables.Member      `json:"Members"`
	Groups     []tables.FriendGroup `json:"Groups"`
}

func (r *response) update(s bool, m []tables.Member, f []tables.FriendGroup, e error) {
	r.Successful = s
	r.ErrMsg = e
	r.Members = m
	r.Groups = f
	if e != nil {
		panic(e)
	}
}
