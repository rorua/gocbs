package model

import (
	"time"

	"fmt"
	"app/shared/database"
)

// *****************************************************************************
// Transaction
// *****************************************************************************
type Transaction struct {
	ID              	uint32 		`db:"id"`
	CreditAccountId 	uint32 		`db:"credit_account_id"`
	CreditAccount	 	string 		`db:"credit_account"`
	DebitAccountId 		uint32 		`db:"debit_account_id"`
	DebitAccount	 	string 		`db:"debit_account"`
	Description		 	string 		`db:"description"`
	Clients			 	string 		`db:"clients"`
	Amount	  			float64 	`db:"amount"`
	Date	  			time.Time   `db:"date"`
	CreatedAt 			time.Time   `db:"created_at"`
	UpdatedAt 			time.Time   `db:"updated_at"`
	Deleted   			uint8       `db:"deleted"`
}

// TransactionID returns the account id
func (u *Transaction) TransactionID() string {
	r := ""
	r = fmt.Sprintf("%v", u.ID)
	return r
}

// TransactionByID gets transaction by ID
func TransactionByID(userID string, transactionID string) (Transaction, error) {
	var err error
	result := Transaction{}
	err = database.SQL.Get(&result, "SELECT id, credit_account_id, debit_account_id, amount, date, created_at, description, clients FROM transactions WHERE id = ? LIMIT 1", transactionID)
	return result, standardizeError(err)
}

// transactions by account id
func TransactionsByAccountId(accountID string) ([]Transaction, error) {
	var err error
	var result []Transaction
	err = database.SQL.Select(&result, `
		select 
			t.id as id, 
			credit_account_id, 
			c.number as credit_account, 
			debit_account_id, 
			d.number as debit_account, 
			amount, 
			clients, 
			date, 
			t.created_at as created_at, 
			t.updated_at as updated_at, 
			description  
		from transactions t
		inner join accounts c on t.credit_account_id = c.id
		inner join accounts d on t.debit_account_id = d.id
		where t.credit_account_id = ?
			or t.debit_account_id = ?
	`, accountID, accountID)
	return result, standardizeError(err)
}

// transactions by account id
func CreditAccountSum(accountID uint32, date string) float64 {
	var err error
	var result []Transaction
	err = database.SQL.Select(&result, `
		select 
			t.id as id, 
			credit_account_id, 
			c.number as credit_account, 
			debit_account_id, 
			d.number as debit_account, 
			amount, 
			clients, 
			date, 
			t.created_at as created_at, 
			t.updated_at as updated_at, 
			description  
		from transactions t
		inner join accounts c on t.credit_account_id = c.id
		inner join accounts d on t.debit_account_id = d.id
		where t.credit_account_id = ? and t.date < ?
	`, accountID, date)
	if err != nil {
		fmt.Println("---")
	}
	sum := 0.
	for _, n := range result  {
		sum += n.Amount
	}
	return sum
}

// transactions by account id
func DebitAccountSum(accountID uint32, date string) float64 {
	var err error
	var result []Transaction
	err = database.SQL.Select(&result, `
		select 
			t.id as id, 
			credit_account_id, 
			c.number as credit_account, 
			debit_account_id, 
			d.number as debit_account, 
			amount, 
			clients, 
			date, 
			t.created_at as created_at, 
			t.updated_at as updated_at, 
			description  
		from transactions t
		inner join accounts c on t.credit_account_id = c.id
		inner join accounts d on t.debit_account_id = d.id
		where t.debit_account_id = ? and t.date < ?
	`, accountID, date)
	if err != nil {
		fmt.Println("---")
	}
	sum := 0.
	for _, n := range result  {
		sum += n.Amount
	}
	return sum
}

// Gets all transactions
func TransactionsAll() ([]Transaction, error) {
	var err error
	var result []Transaction
	err = database.SQL.Select(&result, `
		select 
			t.id as id, 
			credit_account_id, 
			c.number as credit_account, 
			debit_account_id, 
			d.number as debit_account, 
			amount, 
			clients, 
			date, 
			t.created_at as created_at, 
			t.updated_at as updated_at, 
			description  
		from transactions t
		inner join accounts c on t.credit_account_id = c.id
		inner join accounts d on t.debit_account_id = d.id
	`)
	return result, standardizeError(err)
}
//
// NoteCreate creates a note
func TransactionCreate(creditAccountId, debitAccountId, description, clients, date, amount string) error {

	var err error
	now := time.Now()

	_, err = database.SQL.Exec(`
			INSERT INTO transactions (
				credit_account_id, debit_account_id, amount, date, description, clients, created_at, updated_at
			)
			VALUES (?,?,?,?,?,?,?,?)
		`, creditAccountId, debitAccountId, amount, date, description, clients, now, now)

	return standardizeError(err)
}