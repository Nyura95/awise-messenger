package modelsv2

import "time"

// Message table model
type Message struct {
	ID             int
	IDAccount      int
	IDConversation int
	Message        string
	IDStatus       int
	CreatedAt      time.Time
	UpdatedAt      time.Time
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

// Update a message
func (m *Message) Update() error {
	stmt, err := db.Prepare("UPDATE tbl_messages SET id_account = ?, id_conversation = ?, message = ?, id_status = ?, updated_at = ? WHERE id = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(m.IDAccount, m.IDAccount, m.Message, m.IDStatus, time.UTC, m.ID)
	if err != nil {
		return err
	}

	return nil
}

// Create a new message
func (m *Message) Create() error {
	stmt, err := db.Prepare("INSERT INTO tbl_messages(id_account, id_conversation, message, id_status, created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(m.IDAccount, m.IDConversation, m.Message, m.IDStatus, time.UTC, time.UTC)
	if err != nil {
		return err
	}

	return nil
}
