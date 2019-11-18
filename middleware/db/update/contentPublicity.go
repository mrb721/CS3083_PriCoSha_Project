package update

import (
	"github.com/adamsanghera/mysqlBus"
)

//ContentPublicity ...
//Allows user to make posts public or private
func ContentPublicity(id int, publicBit bool) error {
	stmt, err := mysqlBus.DB.Prepare("UPDATE Content SET public = ? WHERE id = ? ")
	if err != nil {
		return err
	}

	_, err = stmt.Exec(publicBit, id)
	if err != nil {
		return err
	}
	return nil

}
