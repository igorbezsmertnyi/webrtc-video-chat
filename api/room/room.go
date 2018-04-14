package room

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"webrtc-video-chat/utils"
)

var db *sql.DB

type e map[string]string

// Room row structure
type Room struct {
	Id        int64     `json:"id"`
	Slug      string    `json:"slug"`
	Peer      string    `json:"peer"`
	CreatedAt time.Time `json:"created_at"`
}

// Handler for creating new room
func Handler(w http.ResponseWriter, r *http.Request) {
	room, err := unmarshalRoom(r)

	if err != nil {
		errorResponse(r, w, err)
		return
	}

	created, err := createRoom(db, room.Peer)

	if err != nil {
		errorResponse(r, w, err)
		return
	}

	respond(r, w, http.StatusOK, created)
}

func unmarshalRoom(r *http.Request) (*Room, error) {
	fmt.Printf("%w", r.Body)
	defer r.Body.Close()

	var room Room

	bytes, err := ioutil.ReadAll(r.Body)

	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(bytes, &room)

	if err != nil {
		return nil, err
	}

	return &room, nil
}

func createRoom(db *sql.DB, peer string) (*Room, error) {
	created := Room{}
	slug := utils.Slug()

	row := db.QueryRow(
		`INSERT INTO rooms (slug, peer) VALUES ($1, $2) RETURNING id, slug, peer, created_at`,
		slug, peer,
	)

	err := row.Scan(&created.Id, &created.Slug, &created.Peer, &created.CreatedAt)

	if err != nil {
		return nil, err
	}

	return &created, nil
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
