package models

import (
	"awise-messenger/helpers"
	"time"
)

// Message table model
type Message struct {
	ID             int    `json:"id"`
	IDAccount      int    `json:"idAccount"`
	IDConversation int    `json:"idConversation"`
	Message        string `json:"message"`
	IDStatus       int    `json:"idStatus"`
	delete         int
	CreatedAt      time.Time `json:"createdAt"`
	UpdatedAt      time.Time `json:"updatedAt"`
}

// FindMessage for find one access_token by id
func FindMessage(id int) (*Message, error) {
	message := Message{}
	result, err := db.Query("SELECT id, id_account, id_conversation, message, id_status, created_at, updated_at FROM tbl_messages WHERE id = ? LIMIT 1", id)
	if err != nil {
		return &message, err
	}
	defer result.Close()
	for result.Next() {
		err := result.Scan(&message.ID, &message.IDAccount, &message.IDConversation, &message.Message, &message.IDStatus, &message.CreatedAt, &message.UpdatedAt)
		if err != nil {
			panic(err.Error())
		}
	}
	return &message, nil
}

// FindAllMessage for find all message in the database
func FindAllMessage() ([]*Message, error) {
	messages := []*Message{}
	result, err := db.Query("SELECT id, id_account, id_conversation, message, id_status, created_at, updated_at FROM tbl_messages")
	if err != nil {
		return messages, err
	}
	defer result.Close()
	for result.Next() {
		message := Message{}
		err := result.Scan(&message.ID, &message.IDAccount, &message.IDConversation, &message.Message, &message.IDStatus, &message.CreatedAt, &message.UpdatedAt)
		if err != nil {
			panic(err.Error())
		}
		messages = append(messages, &message)
	}
	return messages, nil
}

// FindAllMessageByIDConversation for find all message in the database
func FindAllMessageByIDConversation(IDConversation int, nb int, page int) ([]*Message, error) {
	messages := []*Message{}

	var count int
	err := db.QueryRow("SELECT COUNT(*) FROM tbl_messages WHERE id_conversation = ? AND `delete` = 0", IDConversation).Scan(&count)
	if err != nil {
		return messages, err
	}

	nbMaxPage := count / nb
	if count%nb > 0 {
		nbMaxPage = nbMaxPage + 1
	}
	if page == 0 {
		page = 1
	}

	if page > nbMaxPage {
		page = nbMaxPage
	}
	offset := page*nb - nb

	if offset < 0 {
		offset = 0
	}

	result, err := db.Query("SELECT id, id_account, id_conversation, message, id_status, created_at, updated_at FROM tbl_messages WHERE id_conversation = ? AND `delete` = 0 ORDER BY id DESC LIMIT ? OFFSET ?", IDConversation, nb, offset)
	if err != nil {
		return messages, err
	}
	defer result.Close()
	for result.Next() {
		message := Message{}
		err := result.Scan(&message.ID, &message.IDAccount, &message.IDConversation, &message.Message, &message.IDStatus, &message.CreatedAt, &message.UpdatedAt)
		if err != nil {
			panic(err.Error())
		}
		messages = append(messages, &message)
	}

	return messages, nil
}

// Update this message
func (m *Message) Update() error {
	stmt, err := db.Prepare("UPDATE tbl_messages SET id_account = ?, id_conversation = ?, message = ?, id_status = ?, updated_at = ? WHERE id = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(m.IDAccount, m.IDConversation, m.Message, m.IDStatus, helpers.GetUtc(), m.ID)
	if err != nil {
		return err
	}

	return nil
}

// Delete this message
func (m *Message) Delete() error {
	stmt, err := db.Prepare("UPDATE tbl_messages SET `delete` = ?,  updated_at = ? WHERE id = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(1, helpers.GetUtc(), m.ID)
	if err != nil {
		return err
	}

	return nil
}

// CreateMessage for insert a new message into the database
func CreateMessage(IDAccount int, IDConversation int, msg string, IDStatus int) (*Message, error) {
	message := &Message{}
	stmt, err := db.Prepare("INSERT INTO tbl_messages(id_account, id_conversation, message, id_status, `delete`, created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?, ?)")
	if err != nil {
		return message, err
	}
	defer stmt.Close()

	utc := helpers.GetUtc()
	result, err := stmt.Exec(IDAccount, IDConversation, msg, IDStatus, 0, utc, utc)
	if err != nil {
		return message, err
	}

	ID, err := result.LastInsertId()
	if err != nil {
		return message, err
	}

	message, err = FindMessage(int(ID))
	if err != nil {
		return message, err
	}

	return message, nil
}
