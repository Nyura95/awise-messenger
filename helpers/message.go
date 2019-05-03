package helpers

import "awise-messenger/models"

// CountMessageNotReadByIDUser  count the row by iduser and idstatus 1
func CountMessageNotReadByIDUser(IDUser int, messages []*models.Message) int {
	var count int
	for _, message := range messages {
		if message.IDUser == IDUser && message.IDStatus == 1 {
			count++
		}
	}
	return count
}
