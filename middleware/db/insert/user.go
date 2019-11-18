package insert

import (
	"fmt"

	"../tables"
	"github.com/adamsanghera/mysqlBus"
)

//User ...
// Adds a user to the database; returns nil if it works ok.
func User(user tables.Person) error {
	fmt.Println("preparing")
	stmt, err := mysqlBus.DB.Prepare("INSERT INTO Person VALUES(?, ?, ?, ?, ?, ?)")
	fmt.Println("prepared")
	fmt.Println(err)
	if err != nil {
		return err
	}
	fmt.Println("executing")
	_, err = stmt.Exec(user.Username, user.HashedPassword, user.Salt, user.Fname, user.Lname, user.ColorPalette)
	fmt.Println("executed")
	fmt.Println(err)
	if err != nil {
		return err
	}
	fmt.Println("no error")
	fmt.Println(err)
	return nil
}
