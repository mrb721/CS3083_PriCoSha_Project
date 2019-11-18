package update

import (
	"github.com/adamsanghera/mysqlBus"
)

//TaggeePermissions ...
//Change all the tag permissions of a user
func TaggeePermissions(taggee string, switchTo bool) error {
	stmt, err := mysqlBus.DB.Prepare("UPDATE Tag SET status  = ? WHERE username_taggee = ? ")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(switchTo, taggee)
	if err != nil {
		return err
	}
	return nil

}
