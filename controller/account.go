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
	"github.com/julienschmidt/httprouter"
	"github.com/gorilla/context"
	"github.com/josephspurrier/csrfbanana"
)

// AccountGET displays the accounts
func AccountIndexGET(w http.ResponseWriter, r *http.Request) {
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
	v.Name = "accounts/index"
	v.Vars["first_name"] = sess.Values["first_name"]
	v.Vars["user_id"] = userID
	v.Vars["accounts"] = accounts
	v.Render(w)
}

func AccountShowGET(w http.ResponseWriter, r *http.Request) {
	// Get session
	sess := session.Instance(r)

	userID := fmt.Sprintf("%s", sess.Values["id"])
	var params httprouter.Params
	params = context.Get(r, "params").(httprouter.Params)
	accountID := params.ByName("id")

	account, err := model.AccountByID(accountID)
	if err != nil {
		log.Println(err)
		account = model.Account{}
	}

	// Display the view
	v := view.New(r)
	v.Name = "accounts/show"
	v.Vars["first_name"] = sess.Values["first_name"]
	v.Vars["user_id"] = userID
	v.Vars["account"] = account
	v.Render(w)
}

func AccountCreateGET(w http.ResponseWriter, r *http.Request) {
	// Get session
	sess := session.Instance(r)

	// Display the view
	v := view.New(r)
	v.Name = "accounts/create"
	v.Vars["token"] = csrfbanana.Token(w, r, sess)
	v.Render(w)
}

func AccountCreatePOST(w http.ResponseWriter, r *http.Request)  {
	sess := session.Instance(r)

	// Validate with required fields
	if validate, missingField := view.Validate(r, []string{"number", "name", "type"}); !validate {
		sess.AddFlash(view.Flash{"Field missing: " + missingField, view.FlashError})
		sess.Save(r, w)
		AccountCreateGET(w, r)
		return
	}

	// Get form values
	name := r.FormValue("name")
	number := r.FormValue("number")
	type_name := r.FormValue("type")

	// Get database result
	err := model.AccountCreate(name, number, type_name)
	// Will only error if there is a problem with the query
	if err != nil {
		log.Println(err)
		sess.AddFlash(view.Flash{"An error occurred on the server. Please try again later.", view.FlashError})
		sess.Save(r, w)
	} else {
		sess.AddFlash(view.Flash{"Счет  Добавлен!", view.FlashSuccess})
		sess.Save(r, w)
		http.Redirect(w, r, "/accounts", http.StatusFound)
		return
	}

	// Display the same page
	AccountCreateGET(w, r)
}
