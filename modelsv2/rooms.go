package modelsv2

import (
	"awise-messenger/helpers"
	"errors"
	"time"
)

// Room table model
type Room struct {
	ID             int
	IDConversation int
	IDAccount      int
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

// FindRoom for find one room by id
func FindRoom(id int) (*Room, error) {
	room := Room{}
	result, err := db.Query("SELECT id, id_conversation, id_account, created_at, updated_at FROM tbl_rooms essages WHERE id = ? LIMIT 1", id)
	if err != nil {
		return &room, err
	}
	defer result.Close()
	for result.Next() {
		err := result.Scan(&room.ID, &room.IDConversation, &room.IDAccount, &room.CreatedAt, &room.UpdatedAt)
		if err != nil {
			panic(err.Error())
		}
	}
	return &room, nil
}

// FindAllRooms for find all rooms in the database
func FindAllRooms() ([]*Room, error) {
	rooms := []*Room{}
	result, err := db.Query("SELECT id, id_conversation, id_account, created_at, updated_at FROM tbl_rooms")
	if err != nil {
		return rooms, err
	}
	defer result.Close()
	for result.Next() {
		room := Room{}
		err := result.Scan(&room.ID, &room.IDConversation, &room.IDAccount, &room.CreatedAt, &room.UpdatedAt)
		if err != nil {
			panic(err.Error())
		}
		rooms = append(rooms, &room)
	}
	return rooms, nil
}

// FindAllRoomsByIDConversation for find all rooms by id_conversation in the database
func FindAllRoomsByIDConversation(IDConversation int) ([]*Room, error) {
	rooms := []*Room{}
	result, err := db.Query("SELECT id, id_conversation, id_account, created_at, updated_at FROM tbl_rooms WHERE id_conversation = ?", IDConversation)
	if err != nil {
		return rooms, err
	}
	defer result.Close()
	for result.Next() {
		room := Room{}
		err := result.Scan(&room.ID, &room.IDConversation, &room.IDAccount, &room.CreatedAt, &room.UpdatedAt)
		if err != nil {
			panic(err.Error())
		}
		rooms = append(rooms, &room)
	}
	return rooms, nil
}

// FindAllRoomsByIDAccount for find all rooms by id_account in the database
func FindAllRoomsByIDAccount(IDAccount int) ([]*Room, error) {
	rooms := []*Room{}
	result, err := db.Query("SELECT id, id_conversation, id_account, created_at, updated_at FROM tbl_rooms WHERE id_account = ?", IDAccount)
	if err != nil {
		return rooms, err
	}
	defer result.Close()
	for result.Next() {
		room := Room{}
		err := result.Scan(&room.ID, &room.IDConversation, &room.IDAccount, &room.CreatedAt, &room.UpdatedAt)
		if err != nil {
			panic(err.Error())
		}
		rooms = append(rooms, &room)
	}
	return rooms, nil
}

// Update a room
func (r *Room) Update() error {
	stmt, err := db.Prepare("UPDATE tbl_rooms SET id_conversation = ?, id_account = ?, updated_at = ? WHERE id = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(r.IDConversation, r.IDAccount, time.UTC, r.ID)
	if err != nil {
		return err
	}

	return nil
}

// CreateRoom new conversation
func CreateRoom(IDConversation int, IDAccount int) (*Room, error) {
	room := &Room{}
	stmt, err := db.Prepare("INSERT INTO tbl_rooms(id_conversation, id_account, created_at, updated_at) VALUES (?, ?, ?, ?)")
	if err != nil {
		return room, err
	}
	defer stmt.Close()

	utc := helpers.GetUtc()

	result, err := stmt.Exec(IDConversation, IDAccount, utc, utc)
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
func CreateRoomForMultipleAccount(IDConversation int, IDAccounts ...int) error {
	for _, IDAccount := range IDAccounts {
		room, err := CreateRoom(IDConversation, IDAccount)
		if err != nil {
			return err
		}
		if room.ID == 0 {
			return errors.New("Error during creating a new room")
		}
	}
	return nil
}
