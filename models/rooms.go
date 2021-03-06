package models

import (
	"awise-messenger/helpers"
	"strconv"
	"sync"
	"time"
)

// Room table model
type Room struct {
	ID             int       `json:"id"`
	IDConversation int       `json:"idConversation"`
	IDAccount      int       `json:"idAccount"`
	Token          string    `json:"token"`
	CreatedAt      time.Time `json:"createdAt"`
	UpdatedAt      time.Time `json:"cupdatedAt"`
}

// FindRoom for find one room by id
func FindRoom(id int) (*Room, error) {
	room := Room{}
	result, err := db.Query("SELECT id, id_conversation, id_account, token, created_at, updated_at FROM tbl_rooms WHERE id = ? LIMIT 1", id)
	if err != nil {
		return &room, err
	}
	defer result.Close()
	for result.Next() {
		err := result.Scan(&room.ID, &room.IDConversation, &room.IDAccount, &room.Token, &room.CreatedAt, &room.UpdatedAt)
		if err != nil {
			panic(err.Error())
		}
	}
	return &room, nil
}

// FindRoomByToken for find one room by token
func FindRoomByToken(token string) (*Room, error) {
	room := Room{}
	result, err := db.Query("SELECT id, id_conversation, id_account, token, created_at, updated_at FROM tbl_rooms WHERE token = ? LIMIT 1", token)
	if err != nil {
		return &room, err
	}
	defer result.Close()
	for result.Next() {
		err := result.Scan(&room.ID, &room.IDConversation, &room.IDAccount, &room.Token, &room.CreatedAt, &room.UpdatedAt)
		if err != nil {
			panic(err.Error())
		}
	}
	return &room, nil
}

// FindAllRooms for find all rooms in the database
func FindAllRooms() ([]*Room, error) {
	rooms := []*Room{}
	result, err := db.Query("SELECT id, id_conversation, id_account, token, created_at, updated_at FROM tbl_rooms")
	if err != nil {
		return rooms, err
	}
	defer result.Close()
	for result.Next() {
		room := Room{}
		err := result.Scan(&room.ID, &room.IDConversation, &room.IDAccount, &room.Token, &room.CreatedAt, &room.UpdatedAt)
		if err != nil {
			panic(err.Error())
		}
		rooms = append(rooms, &room)
	}
	return rooms, nil
}

// FindRoomByIDConversationAndIDAccount for find all rooms by id_conversation in the database
func FindRoomByIDConversationAndIDAccount(IDConversation int, IDAccount int) (*Room, error) {
	room := Room{}
	result, err := db.Query("SELECT id, id_conversation, id_account, token, created_at, updated_at FROM tbl_rooms WHERE id_conversation = ? AND id_account = ? LIMIT 1", IDConversation, IDAccount)
	if err != nil {
		return &room, err
	}
	defer result.Close()
	for result.Next() {
		err := result.Scan(&room.ID, &room.IDConversation, &room.IDAccount, &room.Token, &room.CreatedAt, &room.UpdatedAt)
		if err != nil {
			panic(err.Error())
		}
	}
	return &room, nil
}

// FindAllRoomsByIDAccount for find all rooms by id_account in the database
func FindAllRoomsByIDAccount(IDAccount int) ([]*Room, error) {
	rooms := []*Room{}
	result, err := db.Query("SELECT id, id_conversation, id_account, token, created_at, updated_at FROM tbl_rooms WHERE id_account = ?", IDAccount)
	if err != nil {
		return rooms, err
	}
	defer result.Close()
	for result.Next() {
		room := Room{}
		err := result.Scan(&room.ID, &room.IDConversation, &room.IDAccount, &room.Token, &room.CreatedAt, &room.UpdatedAt)
		if err != nil {
			panic(err.Error())
		}
		rooms = append(rooms, &room)
	}
	return rooms, nil
}

// FindAllRoomsByIDConversation for find all rooms by id_account in the database
func FindAllRoomsByIDConversation(IDConversation int) ([]*Room, []int, error) {
	rooms := []*Room{}
	targets := []int{}
	result, err := db.Query("SELECT id, id_conversation, id_account, token, created_at, updated_at FROM tbl_rooms WHERE id_conversation = ?", IDConversation)
	if err != nil {
		return rooms, targets, err
	}
	defer result.Close()
	for result.Next() {
		room := Room{}
		err := result.Scan(&room.ID, &room.IDConversation, &room.IDAccount, &room.Token, &room.CreatedAt, &room.UpdatedAt)
		if err != nil {
			panic(err.Error())
		}
		rooms = append(rooms, &room)
		targets = append(targets, room.IDAccount)
	}
	return rooms, targets, nil
}

// Update a room
func (r *Room) Update() error {
	stmt, err := db.Prepare("UPDATE tbl_rooms SET id_conversation = ?, id_account = ?, token = ?, updated_at = ? WHERE id = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(r.IDConversation, r.IDAccount, r.Token, time.UTC, r.ID)
	if err != nil {
		return err
	}

	return nil
}

// CreateRoom new conversation
func CreateRoom(IDConversation int, IDAccount int, token string) (*Room, error) {
	room := &Room{}
	stmt, err := db.Prepare("INSERT INTO tbl_rooms(id_conversation, id_account, token, created_at, updated_at) VALUES (?, ?, ?, ?, ?)")
	if err != nil {
		return room, err
	}
	defer stmt.Close()

	utc := helpers.GetUtc()

	result, err := stmt.Exec(IDConversation, IDAccount, token, utc, utc)
	if err != nil {
		return room, err
	}

	ID, err := result.LastInsertId()
	if err != nil {
		return room, err
	}

	room, _ = FindRoom(int(ID))

	return room, nil
}

// CreateRoomForMultipleAccount dsq
func CreateRoomForMultipleAccount(IDConversation int, IDAccounts ...int) chan error {
	var wg sync.WaitGroup
	wg.Add(len(IDAccounts))

	errors := make(chan error, len(IDAccounts))
	defer close(errors)

	for _, IDAccount := range IDAccounts {
		go func(IDConversation int, IDAccount int) {
			defer wg.Done()
			_, err := CreateRoom(IDConversation, IDAccount, helpers.Token(strconv.Itoa(IDConversation)+":"+strconv.Itoa(IDAccount)+"randomstring"))
			if err != nil {
				errors <- err
			}
		}(IDConversation, IDAccount)
	}
	wg.Wait()

	return errors
}
