package remove

import (
	"github.com/adamsanghera/mysqlBus"
)

//Content ...
//Sets username of poster to '[DELETED_USER]'
func Content(username string) error {
	//stmt, err := mysqlBus.DB.Prepare("UPDATE Content SET username = '[DELETED_USER]'  WHERE username IN (SELECT username FROM Content WHERE username = ? ) ")
	stmt, err := mysqlBus.DB.Prepare("UPDATE Content SET username = '[DELETED_USER]'  WHERE username = ?")

	if err != nil {
		return err
	}

	_, err = stmt.Exec(username)
	if err != nil {
		return err
	}
	return nil

}
