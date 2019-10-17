package modelsv2

// User table users models
type User struct {
	UserID  int
	Avatars string
	Fname   string
	Lname   string
}

// FindUser user
func FindUser(id int) (*User, error) {
	user := User{}
	result, err := db.Query("SELECT userID, fname, lname, avatars FROM users WHERE userID = ?", id)
	if err != nil {
		return &user, err
	}
	defer result.Close()
	for result.Next() {
		err := result.Scan(&user.UserID, &user.Fname, &user.Lname, &user.Avatars)
		if err != nil {
			panic(err.Error())
		}
	}
	return &user, nil
}

// FindAllUsers users
func FindAllUsers() ([]*User, error) {
	users := []*User{}
	result, err := db.Query("SELECT userID, fname, lname, avatars FROM users")
	if err != nil {
		return users, err
	}
	defer result.Close()
	for result.Next() {
		user := User{}
		err := result.Scan(&user.UserID, &user.Fname, &user.Lname, &user.Avatars)
		if err != nil {
			panic(err.Error())
		}
		users = append(users, &user)
	}
	return users, nil
}

// Update a user
func (u *User) Update() error {
	stmt, err := db.Prepare("UPDATE users SET Avatars = ?, Fname = ?, Lname = ? WHERE id = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(u.Avatars, u.Fname, u.Lname, u.UserID)
	if err != nil {
		return err
	}

	return nil
}

// Create a new user
func (u *User) Create() error {
	stmt, err := db.Prepare("INSERT INTO users(avatars, fname, lname) VALUES (?, ?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(u.Avatars, u.Fname, u.Lname)
	if err != nil {
		return err
	}

	return nil
}
