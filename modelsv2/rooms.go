package modelsv2

import (
	"awise-messenger/helpers"
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

// FindRoomBetweenTwoAccount for find the rooms between two account
func FindRoomBetweenTwoAccount(IDAccount1 int, IDAccount2 int) (int, error) {
	var IDConversation int
	result, err := db.Query("SELECT id_conversation from tbl_rooms WHERE id_account = ? or id_account = ? GROUP BY id_conversation HAVING count(id_conversation) > 1 LIMIT 1;", IDAccount1, IDAccount2)
	if err != nil {
		return IDConversation, err
	}
	defer result.Close()
	for result.Next() {
		err := result.Scan(&IDConversation)
		if err != nil {
			panic(err.Error())
		}
	}
	return IDConversation, nil
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

// Create new conversation
func (r *Room) Create() error {
	stmt, err := db.Prepare("INSERT INTO tbl_rooms(id_conversation, id_account, created_at, updated_at) VALUES (?, ?, ?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	utc := helpers.GetUtc()
	r.CreatedAt = utc
	r.UpdatedAt = utc

	result, err := stmt.Exec(r.IDConversation, r.IDAccount, r.CreatedAt, r.UpdatedAt)
	if err != nil {
		return err
	}

	ID, err := result.LastInsertId()
	if err != nil {
		return err
	}

	r, _ = FindRoom(int(ID))

	return nil
}

// CreateRoomForMultipleAccount dsq
func CreateRoomForMultipleAccount(IDConversation int, IDAccounts ...int) error {
	for _, IDAccount := range IDAccounts {
		room := Room{IDAccount: IDAccount, IDConversation: IDConversation}
		if err := room.Create(); err != nil {
			return err
		}
	}
	return nil
}
