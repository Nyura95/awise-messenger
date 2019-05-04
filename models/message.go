package models

import "time"

// Message tbl_messages model
type Message struct {
	IDMessage      int
	IDUser         int
	IDConversation int
	Message        string
	IDStatus       int
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

// FindOne message
func (message *Message) FindOne() error {
	err := db.QueryRow("SELECT * FROM tbl_messages WHERE id_message = ? LIMIT 1", message.IDMessage).Scan(&message.IDMessage, &message.IDUser, &message.IDConversation, &message.Message, &message.IDStatus, &message.CreatedAt, &message.UpdatedAt)
	if err != nil {
		return err
	}

	return nil
}

// FindLastMessageSendByIDConversation message
func (message *Message) FindLastMessageSendByIDConversation() error {
	err := db.QueryRow("SELECT * FROM tbl_messages WHERE id_conversation = ? AND id_user = ? ORDER BY id_message DESC LIMIT 1", message.IDConversation, message.IDUser).Scan(&message.IDMessage, &message.IDUser, &message.IDConversation, &message.Message, &message.IDStatus, &message.CreatedAt, &message.UpdatedAt)
	if err != nil {
		return err
	}

	return nil
}

// FindAllMessageByIDConversation from tbl_message
func FindAllMessageByIDConversation(idConversation int) ([]*Message, error) {
	messages := make([]*Message, 0)
	rows, err := db.Query("SELECT * FROM tbl_messages WHERE id_conversation = ? ORDER BY id_message ASC", idConversation)
	if err != nil {
		return messages, err
	}
	defer rows.Close()

	for rows.Next() {
		message := new(Message)
		err := rows.Scan(&message.IDMessage, &message.IDUser, &message.IDConversation, &message.Message, &message.IDStatus, &message.CreatedAt, &message.UpdatedAt)
		if err != nil {
			return messages, err
		}
		messages = append(messages, message)
	}
	if err = rows.Err(); err != nil {
		return messages, err
	}
	return messages, nil
}

// FindAllMessageNotRead from tbl_message
func FindAllMessageNotRead(idConversation int, IDUser int) ([]*Message, error) {
	messages := make([]*Message, 0)
	rows, err := db.Query("SELECT * FROM tbl_messages WHERE id_conversation = ? AND id_status = 1 and id_user = ? ORDER BY id_message ASC", idConversation, IDUser)
	if err != nil {
		return messages, err
	}
	defer rows.Close()

	for rows.Next() {
		message := new(Message)
		err := rows.Scan(&message.IDMessage, &message.IDUser, &message.IDConversation, &message.Message, &message.IDStatus, &message.CreatedAt, &message.UpdatedAt)
		if err != nil {
			return messages, err
		}
		messages = append(messages, message)
	}
	if err = rows.Err(); err != nil {
		return messages, err
	}
	return messages, nil
}

// Create the message
func (message *Message) Create() error {
	insert, err := db.Exec("INSERT INTO tbl_messages(id_user, id_conversation, message, id_status, created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?)", message.IDUser, message.IDConversation, message.Message, message.IDStatus, time.Now(), time.Now())
	if err != nil {
		return err
	}

	lastID, err := insert.LastInsertId()
	if err != nil {
		return err
	}

	message.IDMessage = int(lastID)

	err = message.FindOne()
	if err != nil {
		return err
	}

	return nil
}

// Update the message
func (message *Message) Update() error {
	_, err := db.Exec("UPDATE tbl_messages SET message = ?, id_status = ?, updated_at = ? WHERE id_message = ?", message.Message, message.IDStatus, time.Now(), message.IDMessage)
	if err != nil {
		return err
	}

	return nil
}

// UpdateMessageRead the message
func (message *Message) UpdateMessageRead() error {
	_, err := db.Exec("UPDATE tbl_messages SET id_status = 2, updated_at = ? WHERE id_conversation = ? and id_user = ?", time.Now(), message.IDConversation, message.IDUser)
	if err != nil {
		return err
	}

	return nil
}
