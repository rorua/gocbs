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
	"github.com/josephspurrier/csrfbanana"
	"gocbs/model"
	"log"
)

// ClientReadGET displays the notes in the notepad
func ClientIndexGET(w http.ResponseWriter, r *http.Request) {
	// Get session
	sess := session.Instance(r)

	userID := fmt.Sprintf("%s", sess.Values["id"])

	clients, err := model.ClientsAll()
	if err != nil {
		log.Println(err)
		clients = []model.Client{}
	}
	// Display the view
	v := view.New(r)
	v.Name = "clients/index"
	v.Vars["first_name"] = sess.Values["first_name"]
	v.Vars["user_id"] = userID
	v.Vars["clients"] = clients
	v.Render(w)
}

func ClientCreateGET(w http.ResponseWriter, r *http.Request) {
	// Get session
	sess := session.Instance(r)

	// Display the view
	v := view.New(r)
	v.Name = "clients/create"
	v.Vars["token"] = csrfbanana.Token(w, r, sess)
	v.Render(w)
}

func ClientCreatePOST(w http.ResponseWriter, r *http.Request) {
	// Get session
	sess := session.Instance(r)

	// Validate with required fields
	if validate, missingField := view.Validate(r, []string{"name", "full_name", "client_type_id", "address", "phone", "email"}); !validate {
		sess.AddFlash(view.Flash{"Field missing: " + missingField, view.FlashError})
		sess.Save(r, w)
		ClientCreateGET(w, r)
		return
	}

	// Get form values
	fullName := r.FormValue("full_name")
	name := r.FormValue("name")
	clientTypeId := r.FormValue("client_type_id")
	address := r.FormValue("address")
	phone := r.FormValue("phone")
	email := r.FormValue("email")


	// Get database result
	err := model.ClientCreate(name, fullName, phone, email, address, clientTypeId)
	// Will only error if there is a problem with the query
	if err != nil {
		log.Println(err)
		sess.AddFlash(view.Flash{"An error occurred on the server. Please try again later.", view.FlashError})
		sess.Save(r, w)
	} else {
		sess.AddFlash(view.Flash{"Клиент  Добавлен!", view.FlashSuccess})
		sess.Save(r, w)
		http.Redirect(w, r, "/clients", http.StatusFound)
		return
	}

	// Display the same page
	AccountCreateGET(w, r)
}
