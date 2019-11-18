package settings

type response struct {
	Successful   bool   `json:"Successful"`
	ErrMsg       error  `json:"ErrMsg"`
	ColorPalette string `json:"ColorPalette"`
}

func (r *response) update(s bool, palette string, e error) {
	r.Successful = s
	r.ErrMsg = e
	r.ColorPalette = palette
	if e != nil {
		panic(e)
	}
}
