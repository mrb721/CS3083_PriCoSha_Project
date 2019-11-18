package delete

import (
	"../query"
	"github.com/adamsanghera/mysqlBus"
)

//Content ...
// Deletes a single content item from the database; returns nil if it works ok.
func Content(contentID int) error {
	// Get all associated comments
	comments, err := query.CommentsByID(contentID)
	if err != nil {
		return err
	}

	// Delete all associated comments
	for c := range comments {
		if err = Comment(comments[c]); err != nil {
			return err
		}
	}

	// Get all associated tags
	tags, err := query.TagsByID(contentID)
	if err != nil {
		return err
	}

	// Delete all associated tags
	for t := range tags {
		if err = Tag(tags[t]); err != nil {
			return err
		}
	}

	// Get all associated shares
	shares, err := query.SharesByID(contentID)
	if err != nil {
		return err
	}

	// Delete the shares themselves
	for s := range shares {
		if err = Share(shares[s].ID, shares[s].GroupName, shares[s].Username); err != nil {
			return err
		}
	}

	// Delete the content
	stmt, err := mysqlBus.DB.Prepare("DELETE FROM Content WHERE id = ?")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(contentID)
	if err != nil {
		return err
	}
	return nil
}
