package models

import (
	"time"
)

// Token access_token
type Token struct {
	ID           int
	UserID       int
	Token        string
	RefreshToken string
	FlagDelete   int
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

// FindOneByToken token from accessToken
func (token *Token) FindOneByToken() error {
	err := db.QueryRow("SELECT * FROM access_token WHERE token = ? LIMIT 1", token.Token).Scan(&token.ID, &token.UserID, &token.Token, &token.RefreshToken, &token.FlagDelete, &token.CreatedAt, &token.UpdatedAt)
	if err != nil {
		return err
	}

	return nil
}
