package remove

import (
	"github.com/adamsanghera/mysqlBus"
)

//Comment ...
//Sets username of poster to '[DELETED_USER]'
func Comment(username string) error {
	//stmt, err := mysqlBus.DB.Prepare("UPDATE Comment SET username = '[DELETED_USER]'  WHERE username IN (SELECT username FROM Comment WHERE username = ? ) ")
	stmt, err := mysqlBus.DB.Prepare("UPDATE Comment SET username = '[DELETED_USER]'  WHERE username = ?")

	if err != nil {
		return err
	}

	_, err = stmt.Exec(username)
	if err != nil {
		return err
	}
	return nil

}
