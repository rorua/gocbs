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
	"gocbs/model"
	"log"
)

// AccountGET displays the accounts
func AccountBalanceIndexGET(w http.ResponseWriter, r *http.Request) {
	// Get session
	sess := session.Instance(r)

	userID := fmt.Sprintf("%s", sess.Values["id"])

	accounts, err := model.AccountsAll()
	if err != nil {
		log.Println(err)
		accounts = []model.Account{}
	}


	// Display the view
	v := view.New(r)
	v.Name = "account_balances/index"
	v.Vars["first_name"] = sess.Values["first_name"]
	v.Vars["user_id"] = userID
	v.Vars["accounts"] = accounts
	v.Render(w)
}
