package friendgroup

import (
	"../../db/insert"
	"../../db/query"
	"../../db/remove"
	"../../db/tables"
)

//make ...
// make creates new friend group
func make(username string, groupName string, description string) error {
	newGroup := tables.FriendGroup{}
	newGroup.Username = username
	newGroup.GroupName = groupName
	newGroup.Description = description

	err := insert.FriendGroup(newGroup)
	return err
}

//delete ...
//removes existing friendgroup
func delete(username string, groupName string) error {
	return remove.Group(username, groupName)
}

//get ...
//retreives group
func get(username string, groupName string) (tables.FriendGroup, error) {
	group, err := query.Group(username, groupName)
	if err != nil {
		return group, err
	}
	return group, err
}

//getAllMemberOf ...
//retrieves all groups a person is a member of
func getAllMemberOf(username string) ([]tables.FriendGroup, error) {
	groups, err := query.AllJoinedGroups(username)

	return groups, err

}

func join(username string, groupName string, creator string) error {
	return insert.Member(tables.Member{
		Username:        username,
		GroupName:       groupName,
		UsernameCreator: creator,
	})
}

func getMembers(groupName string, creator string) ([]tables.Member, error) {
	return query.MembersByGroup(groupName, creator)
}

// TODO: delete all content made by the user the group
func leave(username string, groupName string, creator string) error {
	return remove.Member(username, creator, groupName)
}
