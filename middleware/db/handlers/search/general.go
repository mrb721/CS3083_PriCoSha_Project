package search

import (
	"../../db/methods"
	"../../db/tables"
	"github.com/adamsanghera/mysqlBus"
)

// User ...
// searches by user using username, first name, and/or last name
// func User(request string) ([]tables.Person, error) {}
func User(request string) ([]userResult, error) {
	//better method:
	//mysqlBus.DB.Query("SELECT DISTINCT * FROM Person WHERE first_name COLLATE UTF8_GENERAL_CI LIKE '%?' OR  last_name COLLATE UTF_GENERAL_CI LIKE '%?", request, request)

	results := []userResult{}
	tableResults := []tables.Person{}

	fnames, fnErr := FirstName(request)

	if fnErr == nil {
		tableResults = append(tableResults, fnames...)
	} else {
		return results, fnErr
	}
	lnames, lnErr := LastName(request)
	if lnErr == nil {
		tableResults = append(tableResults, lnames...)
	} else {
		return results, lnErr
	}

	usrnames, unErr := Username(request)
	if unErr == nil {
		tableResults = append(tableResults, usrnames...)
	} else {
		return results, unErr
	}

	uniqueUsers := methods.RemoveDupUsers(tableResults)

	for i := range tableResults {
		var usr userResult
		usr.Username = uniqueUsers[i].Username
		usr.FirstName = uniqueUsers[i].Fname
		usr.LastName = uniqueUsers[i].Lname

		results = append(results, usr)
	}
	return results, nil
}

// Group ...
// Allows users to search through friend groups
func Group(request string) ([]groupResult, error) {
	rows, err := mysqlBus.DB.Query("SELECT * FROM FriendGroup WHERE group_name COLLATE UTF8_GENERAL_CI  LIKE '% ?'", request)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	allGroups := []tables.FriendGroup{}

	for rows.Next() {
		fg1 := tables.FriendGroup{}
		if err := rows.Scan(&fg1.GroupName, &fg1.Username, &fg1.Description); err != nil {
			return nil, err
		}

		allGroups = append(allGroups, fg1)
	}

	var results []groupResult

	for i := range allGroups {
		var group groupResult
		group.Creator = allGroups[i].Username
		group.GroupName = allGroups[i].GroupName
		group.Description = allGroups[i].Description

		results = append(results, group)
	}
	return results, nil
}
