package insert

import (
	"errors"

	"../tables"
	"github.com/adamsanghera/mysqlBus"
)

//Share ...
// Adds a share to the database; returns nil if it works ok.
func Share(share tables.Share) error {
	stmt, err := mysqlBus.DB.Prepare("INSERT INTO Share VALUES(?, ?, ?)")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(share.ID, share.GroupName, share.Username)
	if err != nil {
		return errors.New("unable to share this content")
	}
	return nil
}
