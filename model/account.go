package model

import (
	"time"

	"fmt"
	"app/shared/database"
)

// *****************************************************************************
// Account
// *****************************************************************************
type Account struct {
	ID        uint32        `db:"id"`
	Number 	  string        `db:"number"`
	Name 	  string        `db:"name"`
	Type 	  string        `db:"type"`
	CreatedAt time.Time     `db:"created_at"`
	UpdatedAt time.Time     `db:"updated_at"`
	Deleted   uint8         `db:"deleted"`
}

// AccountID returns the account id
func (u *Account) AccountID() string {
	r := ""
	r = fmt.Sprintf("%v", u.ID)
	return r
}

func (u *Account) AccountNumber() string {
	r := ""
	r = fmt.Sprintf("%v", u.Number)
	return r
}

func (u *Account) AccountName() string {
	r := ""
	r = fmt.Sprintf("%v", u.Name)
	return r
}

// AccountByID gets account by ID
func AccountByID(accountID string) (Account, error) {
	var err error
	result := Account{}
	err = database.SQL.Get(&result, "SELECT id, number, name, type, created_at FROM accounts WHERE id = ? LIMIT 1", accountID)
	return result, standardizeError(err)
}

// NotesByUserID gets all notes for a user
func AccountsAll() ([]Account, error) {
	var err error
	var result []Account
	err = database.SQL.Select(&result, "SELECT id, number, name, type, created_at FROM accounts")
	return result, standardizeError(err)
}

// AccountCreate creates a account
func AccountCreate(name string, number string, typeName string) error {
	var err error
	now := time.Now()

	_, err = database.SQL.Exec("INSERT INTO accounts (name, type, number, created_at, updated_at) VALUES (?,?,?,?,?)", name, typeName, number, now, now)
	return standardizeError(err)
}
//
//// NoteUpdate updates a note
//func NoteUpdate(content string, userID string, noteID string) error {
//	var err error
//
//	now := time.Now()
//
//	switch database.ReadConfig().Type {
//	case database.TypeMySQL:
//		_, err = database.SQL.Exec("UPDATE note SET content=? WHERE id = ? AND user_id = ? LIMIT 1", content, noteID, userID)
//	case database.TypeMongoDB:
//		if database.CheckConnection() {
//			// Create a copy of mongo
//			session := database.Mongo.Copy()
//			defer session.Close()
//			c := session.DB(database.ReadConfig().MongoDB.Database).C("note")
//			var note Note
//			note, err = NoteByID(userID, noteID)
//			if err == nil {
//				// Confirm the owner is attempting to modify the note
//				if note.UserID.Hex() == userID {
//					note.UpdatedAt = now
//					note.Content = content
//					err = c.UpdateId(bson.ObjectIdHex(noteID), &note)
//				} else {
//					err = ErrUnauthorized
//				}
//			}
//		} else {
//			err = ErrUnavailable
//		}
//	case database.TypeBolt:
//		var note Note
//		note, err = NoteByID(userID, noteID)
//		if err == nil {
//			// Confirm the owner is attempting to modify the note
//			if note.UserID.Hex() == userID {
//				note.UpdatedAt = now
//				note.Content = content
//				err = database.Update("note", userID+note.ObjectID.Hex(), &note)
//			} else {
//				err = ErrUnauthorized
//			}
//		}
//	default:
//		err = ErrCode
//	}
//
//	return standardizeError(err)
//}
//
//// NoteDelete deletes a note
//func NoteDelete(userID string, noteID string) error {
//	var err error
//
//	switch database.ReadConfig().Type {
//	case database.TypeMySQL:
//		_, err = database.SQL.Exec("DELETE FROM note WHERE id = ? AND user_id = ?", noteID, userID)
//	case database.TypeMongoDB:
//		if database.CheckConnection() {
//			// Create a copy of mongo
//			session := database.Mongo.Copy()
//			defer session.Close()
//			c := session.DB(database.ReadConfig().MongoDB.Database).C("note")
//
//			var note Note
//			note, err = NoteByID(userID, noteID)
//			if err == nil {
//				// Confirm the owner is attempting to modify the note
//				if note.UserID.Hex() == userID {
//					err = c.RemoveId(bson.ObjectIdHex(noteID))
//				} else {
//					err = ErrUnauthorized
//				}
//			}
//		} else {
//			err = ErrUnavailable
//		}
//	case database.TypeBolt:
//		var note Note
//		note, err = NoteByID(userID, noteID)
//		if err == nil {
//			// Confirm the owner is attempting to modify the note
//			if note.UserID.Hex() == userID {
//				err = database.Delete("note", userID+note.ObjectID.Hex())
//			} else {
//				err = ErrUnauthorized
//			}
//		}
//	default:
//		err = ErrCode
//	}
//
//	return standardizeError(err)
//}
