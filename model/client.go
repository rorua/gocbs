package model

import (
	"time"

	"fmt"
	"gocbs/app/database"
)

// *****************************************************************************
// Client
// *****************************************************************************
type Client struct {
	ID           uint32    `db:"id"`
	Name         string    `db:"name"`
	FullName     string    `db:"full_name"`
	Phone        string    `db:"phone_number"`
	Email        string    `db:"email"`
	Address      string    `db:"address"`
	Type         string    `db:"type"`
	ClientTypeId string    `db:"client_type_id"`
	CreatedAt    time.Time `db:"created_at"`
}

// ClientID returns the client id
func (u *Client) ClientID() string {
	r := ""
	r = fmt.Sprintf("%v", u.ID)
	return r
}

func (u *Client) ClientName() string {
	r := ""
	r = fmt.Sprintf("%v", u.Name)
	return r
}

func (u *Client) ClientFullName() string {
	r := ""
	r = fmt.Sprintf("%v", u.FullName)
	return r
}

// ClientByID gets client by ID
func ClientByID(clientID string) (Client, error) {
	var err error
	result := Client{}
	err = database.SQL.Get(&result, "SELECT clients.id, clients.name as name, full_name, phone_number, email, address, t.name as type, client_type_id, created_at FROM clients inner join client_types t on clients.client_type_id = t.id WHERE clients.id = ? LIMIT 1", clientID)
	return result, standardizeError(err)
}

// NotesByUserID gets all notes for a user
func ClientsAll() ([]Client, error) {
	var err error
	var result []Client
	err = database.SQL.Select(&result,
		`SELECT c.id, c.name as name, full_name, phone_number, email, address, t.short_name as type, client_type_id, created_at 
				FROM clients c
				inner join client_types t on c.client_type_id = t.id
	`)
	//fmt.Println(result)
	return result, standardizeError(err)
}

// ClientCreate creates a client
func ClientCreate(name, fullName, phone, email, address, clientTypeId string) error {
	var err error
	now := time.Now()

	_, err = database.SQL.Exec(
		`INSERT INTO clients (name, full_name, phone_number, email, address, client_type_id, created_at, updated_at) 
			VALUES (?,?,?,?,?,?,?,?)`,
		name, fullName, phone, email, address, clientTypeId, now, now)
	return standardizeError(err)
}
