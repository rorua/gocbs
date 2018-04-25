package controller

import (
	"fmt"
	//"log"
	"net/http"

	//"app/model"
	"app/shared/session"
	"app/shared/view"

	//"github.com/gorilla/context"
	//"github.com/josephspurrier/csrfbanana"
	//"github.com/julienschmidt/httprouter"
)

// AccountReadGET displays the notes in the notepad
func AccountReadGET(w http.ResponseWriter, r *http.Request) {
	// Get session
	sess := session.Instance(r)

	userID := fmt.Sprintf("%s", sess.Values["id"])

	// Display the view
	v := view.New(r)
	v.Name = "accounts/read"
	v.Vars["first_name"] = sess.Values["first_name"]
	v.Vars["user_id"] = userID
	v.Render(w)
}
