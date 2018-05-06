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
	"app/shared/session"
	"github.com/josephspurrier/csrfbanana"
)

// TransactionIndexGET displays the Transactions
func TransactionIndexGET(w http.ResponseWriter, r *http.Request) {
	// Get session
	//sess := session.Instance(r)

	transactions, err := model.TransactionsAll()
	if err != nil {
		log.Println(err)
		transactions = []model.Transaction{}
	}
	// Display the view
	v := view.New(r)
	v.Name = "transactions/index"
	v.Vars["transactions"] = transactions
	v.Render(w)
}

//
func TransactionCreateGET(w http.ResponseWriter, r *http.Request) {
	// Get session
	sess := session.Instance(r)

	accounts, err := model.AccountsAll()
	if err != nil {
		log.Println(err)
		accounts = []model.Account{}
	}
	// Display the view
	v := view.New(r)
	v.Name = "transactions/create"
	v.Vars["accounts"] = accounts
	v.Vars["token"] = csrfbanana.Token(w, r, sess)
	v.Render(w)
}

func TransactionCreatePOST(w http.ResponseWriter, r *http.Request) {
	// Get session
	sess := session.Instance(r)
	data := []string{"credit_account_id", "debit_account_id", "amount", "date", "description"}

	// Validate with required fields
	if validate, missingField := view.Validate(r, data); !validate {
		sess.AddFlash(view.Flash{"Field missing: " + missingField, view.FlashError})
		sess.Save(r, w)
		TransactionCreateGET(w, r)
		return
	}

	amount := r.FormValue("amount")
	description := r.FormValue("description")
	clients := r.FormValue("clients")
	creditAccountId := r.FormValue("credit_account_id")
	debitAccountId := r.FormValue("debit_account_id")
	date := r.FormValue("date")

	// Get database result
	err := model.TransactionCreate(creditAccountId, debitAccountId, description, clients, date, amount)
	// Will only error if there is a problem with the query
	if err != nil {
		log.Println(err)
		sess.AddFlash(view.Flash{"An error occurred on the server. Please try again later.", view.FlashError})
		sess.Save(r, w)
	} else {
		sess.AddFlash(view.Flash{"Проводка Добавлена!", view.FlashSuccess})
		sess.Save(r, w)
		http.Redirect(w, r, "/transactions", http.StatusFound)
		return
	}

	// Display the same page
	TransactionCreateGET(w, r)

}
