package methods

import (
	"../tables"
)

//RemoveDupUsers ...
//Removes duplicate users from array
func RemoveDupUsers(arr []tables.Person) []tables.Person {
	encountered := map[tables.Person]bool{}
	result := []tables.Person{}

	for v := range arr {
		if encountered[arr[v]] == true {

		} else {
			encountered[arr[v]] = true
			result = append(result, arr[v])
		}
	}

	return result

}

//RemoveDupPosts ...
//Removes duplicate users from array
func RemoveDupPosts(arr []tables.Content) []tables.Content {
	encountered := map[tables.Content]bool{}
	result := []tables.Content{}

	for v := range arr {
		if encountered[arr[v]] == true {

		} else {
			encountered[arr[v]] = true
			result = append(result, arr[v])
		}
	}

	return result

}
