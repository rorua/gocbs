package controller

import (
	//"log"
	"net/http"

	"gocbs/app/view"

	//"github.com/gorilla/context"
	//"github.com/josephspurrier/csrfbanana"
	//"github.com/julienschmidt/httprouter"
	"fmt"
	"github.com/josephspurrier/csrfbanana"
	"gocbs/app/session"
	"gocbs/model"
	"log"
)

// AccountGET displays the accounts
func AccountBalanceIndexGET(w http.ResponseWriter, r *http.Request) {
	// Get session
	sess := session.Instance(r)

	date, ok := r.URL.Query()["date"]
	dat := "2018-05-05"
	if !ok {
		log.Println("Url Param 'date' is missing")
	} else {
		dat = date[0]
	}

	query := r.URL.Query()
	qCount := query.Get("date")
	fmt.Println(qCount)

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
	v.Vars["token"] = csrfbanana.Token(w, r, sess)
	v.Render(w)
}

// AccountGET displays the accounts
func AccountBalanceRestartPOST(w http.ResponseWriter, r *http.Request) {

	// Get session
	sess := session.Instance(r)
	// Validate with required fields
	if validate, missingField := view.Validate(r, []string{"date"}); !validate {
		sess.AddFlash(view.Flash{"Field missing: " + missingField, view.FlashError})
		sess.Save(r, w)
		AccountBalanceIndexGET(w, r)
		return
	}

	//fmt.Println(dat)
	date := r.FormValue("date")
	fmt.Println(date)
	fmt.Println("=======")

	err := model.AccountBalanceCreate(date)
	if err != nil {
		log.Println(err)
		sess.AddFlash(view.Flash{"An error occurred on the server. Please try again later.", view.FlashError})
		sess.Save(r, w)
	} else {
		sess.AddFlash(view.Flash{"ОСВ Перезапущен!", view.FlashSuccess})
		sess.Save(r, w)
		http.Redirect(w, r, "/account-balances?date="+date, http.StatusFound)
		return
	}
	// Display the same page
	AccountBalanceIndexGET(w, r)
}
