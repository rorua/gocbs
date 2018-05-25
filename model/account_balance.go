package model

import (
	"time"

	"fmt"
	"app/shared/database"
	"log"
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

// AccountBalanceID returns the account id
func AccountBalanceBalanceAccount(accountId int32) AccountBalance {
	var a = AccountBalance{}
	fmt.Print(accountId)
	return a
}

// AccountBalanceCreate creates a account
/*
	1.	Создать массив с структурой, совпадающей с таблицей account_balances, которая хранит обороты по дебиту (поле debit_sum в БД), обороты по кредиту (credit_sum), начальное сальдо (start_balance), конечное сальдо (end_balance), а также ID текущего счета из таблицы accounts;
	2.	Выбрать дату, на которую нужно вычислить ОСВ;
	3.	Удалить из таблицы account_balances все записи с выбранной даты и больше;
	4.	Достать все счета из таблицы accounts;
	5.	Для каждого счета, посчитать баланс до выбранной даты;
	6.	Полученный массив балансов счетов, вставить в таблицу account_balance, SQL запросом INSERT INTO….
*/
func AccountBalanceCreate(date string) error {

	var err error
	var accountBalances = []AccountBalance{}
	_, err = database.SQL.Exec("DELETE FROM account_balances WHERE date >= ?", date)
	fmt.Println("deleting...")

	accounts, err := AccountsAll()
	if err != nil {
		log.Println(err)
		accounts = []Account{}
	}
	fmt.Print(accounts)
	fmt.Print(accountBalances)



	//now := time.Now()
	//
	//_, err = database.SQL.Exec("INSERT INTO accounts (name, type, number, created_at, updated_at) VALUES (?,?,?,?,?)", name, typeName, number, now, now)
	return standardizeError(err)
}

//func AccountBalanceDelete(date string)  error {
//	var err error
//	_, err = database.SQL.Exec("DELETE FROM account_balance WHERE date >= ?", date)
//	err = ErrCode
//	fmt.Println("deleting...")
//	return standardizeError(err)
//}