package settings

import (
	"../../db/query"
	"../../db/tables"
	"github.com/adamsanghera/mysqlBus"
)

//Set ...
//sets a custom color palette for the user
func Set(username string, newColor string) error {
	stmt, err := mysqlBus.DB.Prepare("UPDATE Person SET color_palette = ? WHERE username = ?")

	if err != nil {
		return err
	}

	_, err = stmt.Exec(newColor, username)
	if err != nil {
		return err
	}
	return nil
}

//Get ...
//Retreives the current color palette of a user
func Get(username string) (string, error) {
	usr := tables.Person{}

	usr, err := query.User(username)
	if err != nil {
		return "", err
	}

	return usr.ColorPalette, err
}
