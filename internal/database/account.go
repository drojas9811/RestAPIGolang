package database

import (
	model "RestAPIGolang/internal/models"
	"database/sql"
	"fmt"
)

func GetAccounts() ([]*model.Account, error) {
	//Get database instance
	s := GetDB()
	//GetProcess
	rows, err := s.Query("select * from account")
	if err != nil {
		return nil, err
	}

	accounts := []*model.Account{}
	for rows.Next() {
		account, err := scanIntoAccount(rows)
		if err != nil {
			return nil, err
		}
		accounts = append(accounts, account)
	}

	return accounts, nil
}

func GetAccountByNumber(number int) (*model.Account, error) {
	//Get database instance
	s := GetDB()
	//GetProcess
	rows, err := s.Query("select * from account where number = $1", number)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		return scanIntoAccount(rows)
	}

	return nil, fmt.Errorf("account with number [%d] not found", number)
}

func GetAccountByID(id int) (*model.Account, error) {
	//Get database instance
	s := GetDB()
	//GetProcess
	rows, err := s.Query("select * from account where id = $1", id)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		return scanIntoAccount(rows)
	}

	return nil, fmt.Errorf("account %d not found", id)
}

func CreateAccount(acc *model.Account) error {
	query := `insert into account 
	(first_name, last_name, number, encrypted_password, balance, created_at)
	values ($1, $2, $3, $4, $5, $6)`
	//Get database instance
	s := GetDB()
	//GetProcess
	_, err := s.Query(
		query,
		acc.FirstName,
		acc.LastName,
		acc.Number,
		acc.EncryptedPassword,
		acc.Balance,
		acc.CreatedAt)

	if err != nil {
		return err
	}

	return nil
}

func DeleteAccount(id int) error {
	//Get database instance
	s := GetDB()
	//GetProcess
	_, err := s.Query("delete from account where id = $1", id)
	return err
}

func scanIntoAccount(rows *sql.Rows) (*model.Account, error) {
	account := new(model.Account)
	err := rows.Scan(
		&account.ID,
		&account.FirstName,
		&account.LastName,
		&account.Number,
		&account.EncryptedPassword,
		&account.Balance,
		&account.CreatedAt)

	return account, err
}
