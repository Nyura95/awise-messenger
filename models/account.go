package models

import (
	"awise-messenger/helpers"
	"crypto/md5"
	"encoding/hex"
	"sync"
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

// FindAccountByPassword for find one account by password
func FindAccountByPassword(password string) (*Account, error) {
	account := Account{}
	result, err := db.Query("SELECT id, avatars, firstname, lastname, username, password, id_scope, created_at, updated_at FROM tbl_account WHERE password = ?", password)
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

// CheckAccountExist check if account exist
func CheckAccountExist(IDAccounts ...int) bool {
	jobs := make(chan bool, len(IDAccounts))
	defer close(jobs)

	var wg sync.WaitGroup
	wg.Add(len(IDAccounts))

	for _, IDAccount := range IDAccounts {
		go func(IDAccount int) {
			defer wg.Done()
			account, err := FindAccount(IDAccount)
			if err != nil || account.ID == 0 {
				jobs <- false
			}
		}(IDAccount)
	}

	wg.Wait()

	return len(jobs) == 0
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

// CreateAccount create a new user
func CreateAccount(avatars string, firstname string, lastname string, username string, password string, idScope int) (*Account, error) {
	account := &Account{}
	stmt, err := db.Prepare("INSERT INTO tbl_account(avatars, firstname, lastname, username, password, id_scope, created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?, ?, ?)")
	if err != nil {
		return account, err
	}
	defer stmt.Close()

	utc := helpers.GetUtc()

	hasher := md5.New()
	hasher.Write([]byte(password))

	result, err := stmt.Exec(avatars, firstname, lastname, username, hex.EncodeToString(hasher.Sum(nil)), idScope, utc, utc)
	if err != nil {
		return account, err
	}

	ID, err := result.LastInsertId()
	if err != nil {
		return account, err
	}

	account, _ = FindAccount(int(ID))

	return account, nil
}
