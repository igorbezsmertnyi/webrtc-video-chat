package room

import (
	"database/sql"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"time"
	"webrtc-video-chat/models"
	"webrtc-video-chat/utils"

	"github.com/gorilla/mux"
)

type e map[string]string

// Room row structure
type Room struct {
	Id        int64     `json:"id"`
	Slug      string    `json:"slug"`
	CreatedAt time.Time `json:"created_at"`
}

var db *sql.DB

// Handler for creating new room
func HandlerCreate(w http.ResponseWriter, r *http.Request) {
	db := models.DB
	created, err := createRoom(db)

	if err != nil {
		errorResponse(r, w, err)
		return
	}

	respond(r, w, http.StatusOK, created)
}

// Handler for getting room
func HandlerGet(w http.ResponseWriter, r *http.Request) {
	db := models.DB
	slug := mux.Vars(r)["slug"]

	room, err := selectRoom(db, slug)

	if err != nil {
		errorResponse(r, w, err)
		return
	}

	respond(r, w, http.StatusOK, room)
}

func unmarshalRoom(r *http.Request) (*Room, error) {
	defer r.Body.Close()

	var room Room

	bytes, err := ioutil.ReadAll(r.Body)

	if err != nil {
		return nil, err
	}

	if err = json.Unmarshal(bytes, &room); err != nil {
		return nil, err
	}

	return &room, nil
}

func createRoom(db *sql.DB) (*Room, error) {
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

func selectRoom(db *sql.DB, slug string) (*Room, error) {
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

func response(r *http.Request, w http.ResponseWriter, status int, bytes []byte) {
	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(status)

	if bytes != nil {
		bytes = append(bytes, '\n')
		w.Write(bytes)
	}

	log.Printf(
		"\"%s %s %s\" %d %d\n",
		r.Method, r.URL.Path, r.Proto, status, len(bytes),
	)
}

func errorResponse(r *http.Request, w http.ResponseWriter, err error) {
	bytes, _ := json.Marshal(e{"message": err.Error()})

	response(r, w, http.StatusInternalServerError, bytes)
}

func respond(r *http.Request, w http.ResponseWriter, status int, data interface{}) {
	if data != nil {
		bytes, err := json.Marshal(data)

		if err != nil {
			errorResponse(r, w, err)
			return
		}

		response(r, w, status, bytes)
	} else {
		response(r, w, status, nil)
	}
}
