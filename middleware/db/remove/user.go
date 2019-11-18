package remove

import (
	"github.com/adamsanghera/mysqlBus"
)

//User ...
//Deletes user from person table
func User(username string) error {
	stmt, err := mysqlBus.DB.Prepare("DELETE FROM Person WHERE username = ?")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(username)
	if err != nil {
		return err
	}
	return nil
}
