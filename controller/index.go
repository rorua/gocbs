package controller

import (
	"net/http"

	"fmt"
	"gocbs/app/session"
	"gocbs/app/view"
	"gocbs/model"
)

// IndexGET displays the home page
func IndexGET(w http.ResponseWriter, r *http.Request) {
	// Get session
	session := session.Instance(r)

	if session.Values["id"] != nil {
		// Display the view
		var user_id string
		user_id = fmt.Sprintf("%s", session.Values["id"])
		user, _ := model.UserById(user_id)
		v := view.New(r)
		v.Name = "index/auth"
		v.Vars["first_name"] = session.Values["first_name"]
		v.Vars["last_name"] = user.LastName
		v.Render(w)
	} else {
		// Display the view
		v := view.New(r)
		v.Name = "index/anon"
		v.Render(w)
		return
	}
}
