package models

import (
	"database/sql"
	"time"
	"webrtc-video-chat/utils"
)

// Room row structure
type Room struct {
	Id        int64     `json:"id"`
	Slug      string    `json:"slug"`
	CreatedAt time.Time `json:"created_at"`
}

// Create Room Action
func CreateRoom(db *sql.DB) (*Room, error) {
	created := Room{}
	slug := utils.RandStringBytes(64)
	createdAt := time.Now()

	row, _ := db.Query(
		"INSERT INTO rooms (slug, created_at) VALUES ($1, $2) RETURNING *;",
		slug,
		createdAt,
	)

	row.Next()

	if err := row.Scan(&created.Id, &created.Slug, &created.CreatedAt); err != nil {
		return nil, err
	}

	return &created, nil
}

// Slect Room Action
func SelectRoom(db *sql.DB, slug string) (*Room, error) {
	room := Room{}

	row, _ := db.Query(
		"SELECT * FROM rooms WHERE slug LIKE $1;",
		slug,
	)

	row.Next()

	if err := row.Scan(&room.Id, &room.Slug, &room.CreatedAt); err != nil {
		return nil, err
	}

	return &room, nil
}
