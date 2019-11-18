package remove

import (
	"github.com/adamsanghera/mysqlBus"
)

//Tagger ...
//Sets all tags of a deleted user to '[DELETED_USER]'
func Tagger(tagger string) error {
	stmt, err := mysqlBus.DB.Prepare("UPDATE Tag SET tagger  = '[DELETED_USER]' WHERE username_tagger = ? ")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(tagger)
	if err != nil {
		return err
	}
	return nil
}


