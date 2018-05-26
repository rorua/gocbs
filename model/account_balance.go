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
func AccountBalanceAccount(accountId uint32, date string) (AccountBalance, error) {
	var err error
	t, _ := time.Parse("2006-01-02", date)
	t = t.AddDate(0, 0, -1)
	result := AccountBalance{}
	err = database.SQL.Get(&result, `
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
		where date <= ?
		and account_id = ?
		orderBy date
		Limit 1
		`, date, accountId)
	return result, standardizeError(err)
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
	//1. Создать массив с структурой, совпадающей с таблицей account_balances
	var accountBalances = []AccountBalance{}

	//3. Удалить из таблицы account_balances все записи с выбранной даты и больше;
	_, err = database.SQL.Exec("DELETE FROM account_balances WHERE date >= ?", date)
	fmt.Println("deleting...")

	//4. Достать все счета из таблицы accounts;
	accounts, err := AccountsAll()
	if err != nil {
		log.Println(err)
		accounts = []Account{}
	}
	//fmt.Print(accounts)
	//fmt.Print(accountBalances)

	//5. Для каждого счета, посчитать баланс до выбранной даты;
	for i, account := range accounts {
		fmt.Print(i)
		t, _ := time.Parse("2006-01-02", date)

		startBalance, creditSum, debitSum, endBalance := calculateAccountBalance(date, account)
		balance := AccountBalance {
			StartBalance	: startBalance,
			EndBalance		: endBalance,
			CreditSum		: creditSum,
			DebitSum		: debitSum,
			Date			: t,
			AccountID		: account.ID,
		}
		//fmt.Println(balance)
		accountBalances = append(accountBalances, balance)
	}

	fmt.Println(accountBalances)
	//6. Полученный массив балансов счетов, вставить в таблицу account_balance, SQL запросом INSERT INTO….

	return standardizeError(err)
}

/**
1.	Вытащить из БД баланс счета на дату, за день до выбранной;
2.	Пройтись по проводкам (таблица transactions) выбранной даты и посчитать сумму всех проводок, используя функцию SUM(amount), где дебетовый счет проводки равен нашему счету (WHERE debit_account_id=account_id). На этом шаге найдем обороты по дебиту;
3.	Пройтись по проводкам (таблица transactions) выбранной даты и посчитать сумму всех проводок, используя функцию SUM(amount), где кредитовый счет проводки равен нашему счету (WHERE cred-it_account_id = account_id). На этом шаге найдем обороты по кредиту;
4.	Посмотреть тип счета (полу type): Если счет пассивный (passive) перейти к шагу (5), если счет активный (active) перейти к шагу (6);
5.	Конечное сальдо равно начальному сальдо плюс обороты по кредиту минус обороты по дебиту;
6.	Конечное сальдо равно начальному сальдо плюс обороты по дебиту минус обороты по кредиту;
7.	Вернуть массив
 */
func calculateAccountBalance(date string, account Account) (float64, float64, float64, float64) {
	//1.Вытащить из БД баланс счета на дату, за день до выбранной;
	balance, err := AccountBalanceAccount(account.ID, date)
	if err != nil {
		balance.EndBalance = 0
	}
	startBalance 	:= balance.EndBalance
	//2. Пройтись по проводкам (таблица transactions) выбранной даты и посчитать сумму всех проводок
	creditSum 		:= CreditAccountSum(account.ID, date)
	//3. Пройтись по проводкам (таблица transactions) выбранной даты и посчитать сумму всех проводок
	debitSum 		:= DebitAccountSum(account.ID, date)

	endBalance := 0.
	//4. Посмотреть тип счета (полу type): Если счет пассивный (passive) перейти к шагу (5), если счет активный (active) перейти к шагу (6);
	if account.Type == "passive" {
		//5. Конечное сальдо равно начальному сальдо плюс обороты по кредиту минус обороты по дебиту;
		endBalance = startBalance + creditSum - debitSum
	} else {
		//6. Конечное сальдо равно начальному сальдо плюс обороты по дебиту минус обороты по кредиту;
		endBalance = startBalance + debitSum - creditSum
	}
	//fmt.Println(balance)
	return startBalance, creditSum, debitSum, endBalance
}
