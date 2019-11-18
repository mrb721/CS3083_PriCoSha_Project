package tag

import (
	"../../db/tables"
)

type response struct {
	Successful bool         `json:"Successful"`
	ErrMsg     error        `json:"ErrMsg"`
	Tags       []tables.Tag `json:"Tags"`
}

func (resp *response) updateResp(successful bool, t []tables.Tag, err error) {
	resp.Successful = successful
	resp.ErrMsg = err
	resp.Tags = t
	if err != nil && err.Error() != "Token invalid" {
		panic(err)
	}
}
