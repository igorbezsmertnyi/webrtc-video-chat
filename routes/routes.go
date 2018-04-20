package routes

import (
	"io/ioutil"
	"net/http"
	"webrtc-video-chat/api/room"
	"webrtc-video-chat/ws"

	"github.com/gorilla/mux"
)

func emptyHandler(w http.ResponseWriter, r *http.Request) {
	f, _ := ioutil.ReadFile("./client/dist/index.html")

	w.WriteHeader(http.StatusOK)
	w.Write(f)
}

// NewRoutes builds the routes for the api
func NewRoutes() *mux.Router {
	mux := mux.NewRouter().StrictSlash(true)

	// client static files
	mux.Handle("/", http.FileServer(http.Dir("./client/dist/"))).Methods("GET")
	mux.HandleFunc("/room/{slug}", emptyHandler).Methods("GET")
	mux.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./client/dist/static/"))))

	// client websocket connection
	mux.HandleFunc("/channel/{slug}", ws.Handler).Methods("GET")

	// create new room request
	api := mux.PathPrefix("/api/").Subrouter()

	api.HandleFunc("/create_room", room.HandlerCreate).Methods("POST")
	api.HandleFunc("/get_room/{slug}", room.HandlerGet).Methods("GET")

	return mux
}
