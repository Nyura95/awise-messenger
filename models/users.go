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

// FindAllUsers from users
func FindAllUsers() ([]*User, error) {
	users := make([]*User, 0)
	rows, err := db.Query("SELECT userID, fname, lname, avatars FROM users")
	if err != nil {
		return users, err
	}
	defer rows.Close()

	for rows.Next() {
		user := new(User)
		err := rows.Scan(&user.UserID, &user.Lname, &user.Fname, &user.Avatars)
		if err != nil {
			return users, err
		}
		users = append(users, user)
	}
	if err = rows.Err(); err != nil {
		return users, err
	}
	return users, nil
}
