package helpers

import "messenger/models"

// GetUser chan
func GetUser(userID int, user chan models.User) {
	userSearch := models.User{UserID: userID}
	if err := userSearch.FindOne(); err != nil {
		user <- userSearch
		return
	}
	user <- userSearch
}
