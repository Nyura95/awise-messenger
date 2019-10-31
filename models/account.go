package models

import (
	"crypto/md5"
	"encoding/hex"
	"time"
)

// Account table models
type Account struct {
	ID        int    `json:"id"`
	Avatars   string `json:"avatars"`
	Firstname string `json:"firstName"`
	Lastname  string `json:"lastName"`
	Username  string `json:"username"`
	password  string
	IDScope   int       `json:"idScope"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

// AccountInfos it's the account with info for the front
type AccountInfos struct {
	*Account
	Online bool `json:"online"`
}

// FindAccount for find one account by id
func FindAccount(id int) (*Account, error) {
	account := Account{}
	result, err := db.Query("SELECT id, avatars, firstname, lastname, username, password, id_scope, created_at, updated_at FROM tbl_account WHERE id = ?", id)
	if err != nil {
		return &account, err
	}
	defer result.Close()
	for result.Next() {
		err := result.Scan(&account.ID, &account.Avatars, &account.Firstname, &account.Lastname, &account.Username, &account.password, &account.IDScope, &account.CreatedAt, &account.UpdatedAt)
		if err != nil {
			panic(err.Error())
		}
	}
	return &account, nil
}

// FindAllAccount for find all accounts in the database
func FindAllAccount() ([]*Account, error) {
	accounts := []*Account{}
	result, err := db.Query("SELECT id, avatars, firstname, lastname, username, password, id_scope, created_at, updated_at FROM tbl_account")
	if err != nil {
		return accounts, err
	}
	defer result.Close()
	for result.Next() {
		account := Account{}
		err := result.Scan(&account.ID, &account.Avatars, &account.Firstname, &account.Lastname, &account.Username, &account.password, &account.IDScope, &account.CreatedAt, &account.UpdatedAt)
		if err != nil {
			panic(err.Error())
		}
		accounts = append(accounts, &account)
	}
	return accounts, nil
}

// Update a user
func (a *Account) Update() error {
	stmt, err := db.Prepare("UPDATE tbl_account SET avatars = ?, firstname = ?, lastname = ?, username = ?, password = ?, id_scope = ?, updated_at = ? WHERE id = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(a.Avatars, a.Firstname, a.Lastname, a.Username, a.password, a.IDScope, time.UTC, a.ID)
	if err != nil {
		return err
	}

	return nil
}

// Create a new user
func (a *Account) Create(password string) error {
	stmt, err := db.Prepare("INSERT INTO tbl_account(avatars, firstname, lastname, username, password, id_scope, created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	hasher := md5.New()
	hasher.Write([]byte(password))

	_, err = stmt.Exec(a.Avatars, a.Firstname, a.Lastname, a.Username, hex.EncodeToString(hasher.Sum(nil)), a.IDScope, time.UTC, time.UTC)
	if err != nil {
		return err
	}

	return nil
}
