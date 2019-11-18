package main

import (
	"net/http"

	"./db/query"
	"./handlers/comment"
	"./handlers/content"
	"./handlers/feed"
	"./handlers/friendgroup"
	"./handlers/login"
	"./handlers/register"
	"./handlers/search"
	"./handlers/settings"
	"./handlers/share"
	"./handlers/tag"

	"github.com/adamsanghera/mysqlBus"
)

func getAllUsers(w http.ResponseWriter, r *http.Request) {
	people, err := query.AllUsers()

	output := ""
	if err != nil {
		output += err.Error()
	} else {
		for _, people := range people {
			output += people.Username + "\n"
		}
	}

	w.Write([]byte(output))
}

func main() {
	defer mysqlBus.DB.Close()

	http.HandleFunc("/get/users/all", getAllUsers)
	http.HandleFunc("/login/user", login.Handler)
	http.HandleFunc("/register/user", register.Handler)
	http.HandleFunc("/content", content.Handler)
	http.HandleFunc("/contentByGroup", content.GroupHandler)
	http.HandleFunc("/feed", feed.Handler)
	http.HandleFunc("/comment", comment.Handler)
	http.HandleFunc("/group", friendgroup.Handler)
	http.HandleFunc("/tag", tag.Handler)
	http.HandleFunc("/share", share.Handler)
	http.HandleFunc("/search", search.Handler)
	http.HandleFunc("/settings", settings.Handler)

	http.ListenAndServe(":3000", nil)
}
