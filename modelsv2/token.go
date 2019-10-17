package modelsv2

import "time"

// Token table users models
type Token struct {
	ID           int
	UserID       int
	Token        string
	RefreshToken string
	FlagDelete   int
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

// FindToken token
func FindToken(id int) (*Token, error) {
	token := Token{}
	result, err := db.Query("SELECT id, user_id, token, refresh_token, flag_delete, created_at, updated_at FROM access_token WHERE id = ? LIMIT 1", id)
	if err != nil {
		return &token, err
	}
	defer result.Close()
	for result.Next() {
		err := result.Scan(&token.ID, &token.UserID, &token.Token, &token.RefreshToken, &token.FlagDelete, &token.CreatedAt, &token.UpdatedAt)
		if err != nil {
			panic(err.Error())
		}
	}
	return &token, nil
}

// FindTokenByToken token
func FindTokenByToken(char string) (*Token, error) {
	token := Token{}
	result, err := db.Query("SELECT id, user_id, token, refresh_token, flag_delete, created_at, updated_at FROM access_token WHERE token = ? LIMIT 1", char)
	if err != nil {
		return &token, err
	}
	defer result.Close()
	for result.Next() {
		err := result.Scan(&token.ID, &token.UserID, &token.Token, &token.RefreshToken, &token.FlagDelete, &token.CreatedAt, &token.UpdatedAt)
		if err != nil {
			panic(err.Error())
		}
	}
	return &token, nil
}

// FindAllToken users
func FindAllToken() ([]*Token, error) {
	tokens := []*Token{}
	result, err := db.Query("SELECT id, user_id, token, refresh_token, flag_delete, created_at, updated_at FROM access_token")
	if err != nil {
		return tokens, err
	}
	defer result.Close()
	for result.Next() {
		token := Token{}
		err := result.Scan(&token.ID, &token.UserID, &token.Token, &token.RefreshToken, &token.FlagDelete, &token.CreatedAt, &token.UpdatedAt)
		if err != nil {
			panic(err.Error())
		}
		tokens = append(tokens, &token)
	}
	return tokens, nil
}

// Update a user
func (t *Token) Update() error {
	stmt, err := db.Prepare("UPDATE access_token SET user_id = ?, token = ?, refresh_token = ?, flag_delete = ?, updated_at = ? WHERE id = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(t.UserID, t.Token, t.RefreshToken, t.FlagDelete, time.Now(), t.ID)
	if err != nil {
		return err
	}

	return nil
}

// Create a new user
func (t *Token) Create() error {
	stmt, err := db.Prepare("INSERT INTO users(user_id, token, refresh_token, flag_delete, created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(t.UserID, t.Token, t.RefreshToken, t.FlagDelete, time.Now(), time.Now())
	if err != nil {
		return err
	}

	return nil
}
