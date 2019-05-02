package models

// User table users models
type User struct {
	UserID  int
	Avatars string
	Fname   string
	Lname   string
}

// FindOne user
func (user *User) FindOne() error {
	err := db.QueryRow("SELECT userID, fname, lname, avatars FROM users WHERE userID = ? LIMIT 1", user.UserID).Scan(&user.UserID, &user.Fname, &user.Lname, &user.Avatars)
	if err != nil {
		return err
	}

	return nil
}
