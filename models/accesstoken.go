package models

import "time"

// AccessToken table model
type AccessToken struct {
	ID           int
	IDAccount    int
	Token        string
	RefreshToken string
	FlagDelete   int
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

// FindAccessToken for find one access_token by id
func FindAccessToken(id int) (*AccessToken, error) {
	accessToken := AccessToken{}
	result, err := db.Query("SELECT id, id_account, token, refresh_token, flag_delete, created_at, updated_at FROM tbl_access_token WHERE id = ? LIMIT 1", id)
	if err != nil {
		return &accessToken, err
	}
	defer result.Close()
	for result.Next() {
		err := result.Scan(&accessToken.ID, &accessToken.IDAccount, &accessToken.Token, &accessToken.RefreshToken, &accessToken.FlagDelete, &accessToken.CreatedAt, &accessToken.UpdatedAt)
		if err != nil {
			panic(err.Error())
		}
	}
	return &accessToken, nil
}

// FindAccessTokenByToken for find one access_token by token
func FindAccessTokenByToken(token string) (*AccessToken, error) {
	accessToken := AccessToken{}
	result, err := db.Query("SELECT id, id_account, token, refresh_token, flag_delete, created_at, updated_at FROM tbl_access_token WHERE token = ? LIMIT 1", token)
	if err != nil {
		return &accessToken, err
	}
	defer result.Close()
	for result.Next() {
		err := result.Scan(&accessToken.ID, &accessToken.IDAccount, &accessToken.Token, &accessToken.RefreshToken, &accessToken.FlagDelete, &accessToken.CreatedAt, &accessToken.UpdatedAt)
		if err != nil {
			panic(err.Error())
		}
	}
	return &accessToken, nil
}

// FindAllAccessToken for find all access_token in the database
func FindAllAccessToken() ([]*AccessToken, error) {
	accessTokens := []*AccessToken{}
	result, err := db.Query("SELECT id, id_account, token, refresh_token, flag_delete, created_at, updated_at FROM tbl_access_token")
	if err != nil {
		return accessTokens, err
	}
	defer result.Close()
	for result.Next() {
		accessToken := AccessToken{}
		err := result.Scan(&accessToken.ID, &accessToken.IDAccount, &accessToken.Token, &accessToken.RefreshToken, &accessToken.FlagDelete, &accessToken.CreatedAt, &accessToken.UpdatedAt)
		if err != nil {
			panic(err.Error())
		}
		accessTokens = append(accessTokens, &accessToken)
	}
	return accessTokens, nil
}

// Update a access_token
func (a *AccessToken) Update() error {
	stmt, err := db.Prepare("UPDATE tbl_access_token SET id_account = ?, token = ?, refresh_token = ?, flag_delete = ?, updated_at = ? WHERE id = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(a.IDAccount, a.Token, a.RefreshToken, a.FlagDelete, time.UTC, a.ID)
	if err != nil {
		return err
	}

	return nil
}

// Create a new access_token
func (a *AccessToken) Create() error {
	stmt, err := db.Prepare("INSERT INTO tbl_access_token(id_account, token, refresh_token, flag_delete, created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(a.IDAccount, a.Token, a.RefreshToken, a.FlagDelete, time.UTC, time.UTC)
	if err != nil {
		return err
	}

	return nil
}
