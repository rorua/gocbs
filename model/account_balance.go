package model

import (
	"time"

	"fmt"
	"app/shared/database"
)

// *****************************************************************************
// AccountBalance
// *****************************************************************************
type AccountBalance struct {
	ID           uint32    `db:"id"`
	AccountID    uint32    `db:"account_id"`
	Number       string    `db:"number"`
	DebitSum     float64   `db:"debit_sum"`
	CreditSum    float64   `db:"credit_sum"`
	EndBalance   float64   `db:"end_balance"`
	StartBalance float64   `db:"start_balance"`
	Date    	 time.Time `db:"date"`
}

// AccountBalanceID returns the account id
func (u *AccountBalance) AccountBalanceBalanceID() string {
	r := ""
	r = fmt.Sprintf("%v", u.ID)
	return r
}

//  gets all
func AccountBalancesByDate(date string) ([]AccountBalance, error) {
	var err error
	var result []AccountBalance
	err = database.SQL.Select(&result, `
		select
  			ab.id as id,
  			a.number as number,
  			date,
  			credit_sum,
  			debit_sum,
  			start_balance,
  			end_balance,
  			ab.account_id as account_id
		from account_balances ab
  		inner join accounts a on ab.account_id = a.id
		where date = ?
	`, date)

	return result, standardizeError(err)
}

// AccountBalanceCreate creates a account
func AccountBalanceCreate(date string) error {
	var err error
	//now := time.Now()
	//
	//_, err = database.SQL.Exec("INSERT INTO accounts (name, type, number, created_at, updated_at) VALUES (?,?,?,?,?)", name, typeName, number, now, now)
	return standardizeError(err)
}