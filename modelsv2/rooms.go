package modelsv2

import "time"

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

// Create a new room
func (r *Room) Create() error {
	stmt, err := db.Prepare("INSERT INTO tbl_rooms(id_conversation, id_account, created_at, updated_at) VALUES (?, ?, ?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(r.IDConversation, r.IDAccount, time.UTC, time.UTC)
	if err != nil {
		return err
	}

	return nil
}
