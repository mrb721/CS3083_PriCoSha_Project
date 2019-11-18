package delete

import (
	"../tables"
	"github.com/adamsanghera/mysqlBus"
)

//User ...
// Deletes a single user from the database; returns nil if it works ok.
func User(user tables.Person) error {
	stmt, err := mysqlBus.DB.Prepare("DELETE FROM Person WHERE username = ?")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(user.Username)
	if err != nil {
		return err
	}
	return nil
}
