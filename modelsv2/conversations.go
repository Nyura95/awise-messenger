package modelsv2

import (
	"awise-messenger/helpers"
	"time"
)

// Conversation table model
type Conversation struct {
	ID                int
	UniqHash          string
	Title             string
	TokenConversation string
	IDLastMessage     int
	IDFirstMessage    int
	IDStatus          int
	CreatedAt         time.Time
	UpdatedAt         time.Time
}

// FindConversation for find one conversation by id
func FindConversation(id int) (*Conversation, error) {
	conversation := Conversation{}
	result, err := db.Query("SELECT id, uniq_hash, title, token_conversation, id_last_message, id_first_message, id_status, created_at, updated_at FROM tbl_conversations WHERE id = ? LIMIT 1", id)
	if err != nil {
		return &conversation, err
	}
	defer result.Close()
	for result.Next() {
		err := result.Scan(&conversation.ID, &conversation.UniqHash, &conversation.Title, &conversation.TokenConversation, &conversation.IDLastMessage, &conversation.IDFirstMessage, &conversation.IDStatus, &conversation.CreatedAt, &conversation.UpdatedAt)
		if err != nil {
			panic(err.Error())
		}
	}
	return &conversation, nil
}

// FindAllConversation for find all conversations in the database
func FindAllConversation() ([]*Conversation, error) {
	conversations := []*Conversation{}
	result, err := db.Query("SELECT id, uniq_hash, title, token_conversation, id_last_message, id_first_message, id_status, created_at, updated_at FROM tbl_conversations")
	if err != nil {
		return conversations, err
	}
	defer result.Close()
	for result.Next() {
		conversation := Conversation{}
		err := result.Scan(&conversation.ID, &conversation.UniqHash, &conversation.Title, &conversation.TokenConversation, &conversation.IDLastMessage, &conversation.IDFirstMessage, &conversation.IDStatus, &conversation.CreatedAt, &conversation.UpdatedAt)
		if err != nil {
			panic(err.Error())
		}
		conversations = append(conversations, &conversation)
	}
	return conversations, nil
}

// Update a conversation
func (c *Conversation) Update() error {
	stmt, err := db.Prepare("UPDATE tbl_conversations SET uniq_hash = ?, title = ?, token_conversation = ?, id_last_message = ?, id_first_message = ?, id_status = ?, updated_at = ? WHERE id = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(c.UniqHash, c.Title, c.TokenConversation, c.IDLastMessage, c.IDFirstMessage, c.IDStatus, helpers.GetUtc(), c.ID)
	if err != nil {
		return err
	}

	return nil
}

// Create new conversation
func (c *Conversation) Create() error {
	stmt, err := db.Prepare("INSERT INTO tbl_conversations(uniq_hash, title, token_conversation, id_last_message, id_first_message, id_status, created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?, ?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	utc := helpers.GetUtc()
	c.CreatedAt = utc
	c.UpdatedAt = utc

	result, err := stmt.Exec(c.UniqHash, c.Title, c.TokenConversation, c.IDLastMessage, c.IDFirstMessage, c.IDStatus, c.CreatedAt, c.UpdatedAt)
	if err != nil {
		return err
	}

	ID, err := result.LastInsertId()
	if err != nil {
		return err
	}

	c, _ = FindConversation(int(ID))

	return nil
}
