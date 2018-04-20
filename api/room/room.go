package room

import (
	"database/sql"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"webrtc-video-chat/models"

	"github.com/gorilla/mux"
)

type e map[string]string

var db *sql.DB

// Handler for creating new room
func HandlerCreate(w http.ResponseWriter, r *http.Request) {
	db := models.DB
	created, err := models.CreateRoom(db)

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

	room, err := models.SelectRoom(db, slug)

	if err != nil {
		errorResponse(r, w, err)
		return
	}

	respond(r, w, http.StatusOK, room)
}

func unmarshalRoom(r *http.Request) (*models.Room, error) {
	defer r.Body.Close()

	var room models.Room

	bytes, err := ioutil.ReadAll(r.Body)

	if err != nil {
		return nil, err
	}

	if err = json.Unmarshal(bytes, &room); err != nil {
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
