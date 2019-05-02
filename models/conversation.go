package models

import (
	"log"
	"time"
)

// Conversation tbl_conversation models
type Conversation struct {
	IDConversation int
	UniqHash       string
	Token          string
	Title          string
	IDCreator      int
	IDReceiver     int
	IDLastMessage  int
	IDFirstMessage int
	IDStatus       int
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

// FindAllConversationByIDUser show all conversation per user
func FindAllConversationByIDUser(IDUser int) ([]*Conversation, error) {
	conversations := make([]*Conversation, 0)
	rows, err := db.Query("SELECT * FROM tbl_conversations WHERE (id_creator = ? OR id_receiver = ?)", IDUser, IDUser)
	if err != nil {
		return conversations, err
	}
	defer rows.Close()

	for rows.Next() {
		conversation := new(Conversation)
		err := rows.Scan(&conversation.IDConversation, &conversation.UniqHash, &conversation.Token, &conversation.Title, &conversation.IDCreator, &conversation.IDReceiver, &conversation.IDLastMessage, &conversation.IDFirstMessage, &conversation.IDStatus, &conversation.CreatedAt, &conversation.UpdatedAt)
		if err != nil {
			return conversations, err
		}
		conversations = append(conversations, conversation)
	}
	if err = rows.Err(); err != nil {
		return conversations, err
	}
	return conversations, nil
}

// FindOne a conversation
func (conversation *Conversation) FindOne() error {
	err := db.QueryRow("SELECT * FROM tbl_conversations WHERE id_conversation = ? LIMIT 1", conversation.IDConversation).Scan(&conversation.IDConversation, &conversation.UniqHash, &conversation.Token, &conversation.Title, &conversation.IDCreator, &conversation.IDReceiver, &conversation.IDLastMessage, &conversation.IDFirstMessage, &conversation.IDStatus, &conversation.CreatedAt, &conversation.UpdatedAt)
	if err != nil {
		return err
	}

	return nil
}

// FindOneByHash a conversation
func (conversation *Conversation) FindOneByHash() error {
	err := db.QueryRow("SELECT * FROM tbl_conversations WHERE uniq_hash = ? LIMIT 1", conversation.UniqHash).Scan(&conversation.IDConversation, &conversation.UniqHash, &conversation.Token, &conversation.Title, &conversation.IDCreator, &conversation.IDReceiver, &conversation.IDLastMessage, &conversation.IDFirstMessage, &conversation.IDStatus, &conversation.CreatedAt, &conversation.UpdatedAt)
	if err != nil {
		return err
	}

	return nil
}

// FindOneByToken a conversationn
func (conversation *Conversation) FindOneByToken() error {
	err := db.QueryRow("SELECT * FROM tbl_conversations WHERE token = ? LIMIT 1", conversation.Token).Scan(&conversation.IDConversation, &conversation.UniqHash, &conversation.Token, &conversation.Title, &conversation.IDCreator, &conversation.IDReceiver, &conversation.IDLastMessage, &conversation.IDFirstMessage, &conversation.IDStatus, &conversation.CreatedAt, &conversation.UpdatedAt)
	if err != nil {
		return err
	}

	return nil
}

// Create the conversation
func (conversation *Conversation) Create() error {
	insert, err := db.Exec("INSERT INTO tbl_conversations(uniq_hash, token, title, id_creator, id_receiver, id_last_message, id_first_message, id_status, created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", conversation.UniqHash, conversation.Token, conversation.Title, conversation.IDCreator, conversation.IDReceiver, conversation.IDLastMessage, conversation.IDFirstMessage, conversation.IDStatus, time.Now(), time.Now())
	if err != nil {
		return err
	}

	lastID, err := insert.LastInsertId()
	if err != nil {
		return err
	}

	conversation.IDConversation = int(lastID)

	err = conversation.FindOne()
	if err != nil {
		return err
	}

	return nil
}

// Update the message
func (conversation *Conversation) Update() error {
	_, err := db.Exec("UPDATE tbl_conversations SET title = ?, id_last_message = ?, id_first_message = ?, id_status = ?, updated_at = ? WHERE id_conversation = ?", conversation.Title, conversation.IDLastMessage, conversation.IDFirstMessage, conversation.IDStatus, time.Now(), conversation.IDConversation)
	if err != nil {
		log.Panic(err)
		return err
	}

	return nil
}
