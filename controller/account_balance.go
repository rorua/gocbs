package controller

import (
	//"log"
	"net/http"


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
	//sess := session.Instance(r)

	date, ok := r.URL.Query()["date"]
	dat := "2018-05-05"
	if !ok {
		log.Println("Url Param 'date' is missing")
	} else {
		dat = date[0]
	}

	accounts, err := model.AccountBalancesByDate(dat)
	if err != nil {
		log.Println(err)
		accounts = []model.AccountBalance{}
	}
	// Display the view
	v := view.New(r)
	v.Name = "account_balances/index"
	v.Vars["date"] = dat
	v.Vars["accounts"] = accounts
	v.Render(w)
}
