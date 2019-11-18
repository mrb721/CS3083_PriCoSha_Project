package file

type response struct {
	Successful bool  `json:"Successful"`
	ErrMsg     error `json:"ErrMsg"`
}
