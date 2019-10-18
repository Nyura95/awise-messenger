package models

import (
	"log"
	"time"
)

// Conversation tbl_conversation models
type Conversation struct {
	IDConversation int
	UniqHash       string
	Title          string
	IDCreator      int
	TokenCreator   string
	IDReceiver     int
	TokenReceiver  string
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
		err := rows.Scan(&conversation.IDConversation, &conversation.UniqHash, &conversation.Title, &conversation.IDCreator, &conversation.TokenCreator, &conversation.IDReceiver, &conversation.TokenReceiver, &conversation.IDLastMessage, &conversation.IDFirstMessage, &conversation.IDStatus, &conversation.CreatedAt, &conversation.UpdatedAt)
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
	err := db.QueryRow("SELECT * FROM tbl_conversations WHERE id_conversation = ? LIMIT 1", conversation.IDConversation).Scan(&conversation.IDConversation, &conversation.UniqHash, &conversation.Title, &conversation.IDCreator, &conversation.TokenCreator, &conversation.IDReceiver, &conversation.TokenReceiver, &conversation.IDLastMessage, &conversation.IDFirstMessage, &conversation.IDStatus, &conversation.CreatedAt, &conversation.UpdatedAt)
	if err != nil {
		return err
	}

	return nil
}

// FindOneByHash a conversation
func (conversation *Conversation) FindOneByHash() error {
	err := db.QueryRow("SELECT * FROM tbl_conversations WHERE uniq_hash = ? LIMIT 1", conversation.UniqHash).Scan(&conversation.IDConversation, &conversation.UniqHash, &conversation.Title, &conversation.IDCreator, &conversation.TokenCreator, &conversation.IDReceiver, &conversation.TokenReceiver, &conversation.IDLastMessage, &conversation.IDFirstMessage, &conversation.IDStatus, &conversation.CreatedAt, &conversation.UpdatedAt)
	if err != nil {
		return err
	}

	return nil
}

// FindOneByToken a conversationn
func (conversation *Conversation) FindOneByToken() error {
	err := db.QueryRow("SELECT * FROM tbl_conversations WHERE token_creator = ? OR token_receiver = ? LIMIT 1", conversation.TokenCreator, conversation.TokenReceiver).Scan(&conversation.IDConversation, &conversation.UniqHash, &conversation.Title, &conversation.IDCreator, &conversation.TokenCreator, &conversation.IDReceiver, &conversation.TokenReceiver, &conversation.IDLastMessage, &conversation.IDFirstMessage, &conversation.IDStatus, &conversation.CreatedAt, &conversation.UpdatedAt)
	if err != nil {
		return err
	}

	return nil
}

// FindOneByTokenCreator a conversationn
func (conversation *Conversation) FindOneByTokenCreator() error {
	err := db.QueryRow("SELECT * FROM tbl_conversations WHERE token_creator = ? LIMIT 1", conversation.TokenCreator).Scan(&conversation.IDConversation, &conversation.UniqHash, &conversation.Title, &conversation.IDCreator, &conversation.TokenCreator, &conversation.IDReceiver, &conversation.TokenReceiver, &conversation.IDLastMessage, &conversation.IDFirstMessage, &conversation.IDStatus, &conversation.CreatedAt, &conversation.UpdatedAt)
	if err != nil {
		return err
	}

	return nil
}

// FindOneByTokenReceiver a conversationn
func (conversation *Conversation) FindOneByTokenReceiver() error {
	err := db.QueryRow("SELECT * FROM tbl_conversations WHERE token_receiver = ? LIMIT 1", conversation.TokenReceiver).Scan(&conversation.IDConversation, &conversation.UniqHash, &conversation.Title, &conversation.IDCreator, &conversation.TokenCreator, &conversation.IDReceiver, &conversation.TokenReceiver, &conversation.IDLastMessage, &conversation.IDFirstMessage, &conversation.IDStatus, &conversation.CreatedAt, &conversation.UpdatedAt)
	if err != nil {
		return err
	}

	return nil
}

// Create the conversation
func (conversation *Conversation) Create() error {
	insert, err := db.Exec("INSERT INTO tbl_conversations(uniq_hash, title, id_creator, token_creator, id_receiver, token_receiver, id_last_message, id_first_message, id_status, created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", conversation.UniqHash, conversation.Title, conversation.IDCreator, conversation.TokenCreator, conversation.IDReceiver, conversation.TokenReceiver, conversation.IDLastMessage, conversation.IDFirstMessage, conversation.IDStatus, time.UTC, time.UTC)
	if err != nil {
		log.Panicln(err)
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
	_, err := db.Exec("UPDATE tbl_conversations SET title = ?, id_last_message = ?, id_first_message = ?, id_status = ?, updated_at = ? WHERE id_conversation = ?", conversation.Title, conversation.IDLastMessage, conversation.IDFirstMessage, conversation.IDStatus, time.UTC, conversation.IDConversation)
	if err != nil {
		log.Panic(err)
		return err
	}

	return nil
}
