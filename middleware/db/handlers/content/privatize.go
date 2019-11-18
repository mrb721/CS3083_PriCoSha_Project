package content

import (
	"errors"

	"../../db/query"
	"../../db/update"
)

func privatize(id int, user string) error {
	// Find the content
	cont, err := query.Content(id)
	if err != nil {
		return err
	}

	if cont.Username != user {
		return errors.New("User is not authorized to privatize this content")
	}

	return update.ContentPublicity(id, false)
}
