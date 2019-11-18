package login

import (
	"errors"
	"fmt"
	"time"
)

type response struct {
	Token          string        `json:"Token"`
	ExpirationTime time.Duration `json:"ExpirationTime"`
	ErrMsg         error         `json:"ErrMsg"`
}

func setupResp() *response {
	resp := &response{
		Token:  "",
		ErrMsg: errors.New("Unknown error"),
	}
	return resp
}

func (r *response) updateResp(t string, et time.Duration, em error) {
	r.Token = t
	r.ExpirationTime = et
	r.ErrMsg = em
	if em != nil && em.Error() != "Password does not match our records" {
		fmt.Println("Error: " + em.Error())
		panic(em)
	}
}
