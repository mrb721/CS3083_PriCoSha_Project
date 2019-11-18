package nuke

import (
	"errors"

	"../../db/update"
)

//User ...
//Called when user deactivates account
func User(string username) error {
	var usr tables.Person
	usr.Username = username

	//set all tag permissions to false, if taggee
	taggeeErr := update.TaggeePermissions(usr.Username, false)
	if taggeeErr != nil {
		return errors.New("Unable to change tag permissions")
	}

	//set all tags to [deleted user], if tagger
	taggerErr := remove.Tagger(usr.Username)
	if taggerErr != nil {
		return errors.New("Unable to change tagger username")
	}

	//update posts(content + comments) to be from anonymous source, '[DELETED_USER]'
	contentErr := remove.Content(usr.Username)
	if contentErr != nil {
		return errors.New("Unable to set to '[DELETED_USER]'")
	}

	commentErr := remove.Comment(usr.Username)
	if commentErr != nil {
		return errors.New("Unable to set to '[DELETED_USER]'")
	}

	//remove from all friend groups
	memberErr := remove.MemberFromAllGroups(usr.Username)
	if memberErr != nil {
		return errors.New("Unable to remove from friend group")
	}
	groupErr := remove.AllGroups(usr.Username)
	if groupErr != nil {
		return errors.New("Unable to remove all friend groups")
	}
	//delete user last
	usrErr := remove.User(usr.Username)
	if usrErr != nil {
		return errors.New("Unable to remove user")
	}
}
