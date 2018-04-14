package routes

import (
	"net/http"

	"github.com/igorbezsmertnyi/webrtc-video-chat/ws"

	"github.com/gorilla/mux"
)

// NewRoutes builds the routes for the api
func NewRoutes() *mux.Router {
	mux := mux.NewRouter()

	// client static files
	mux.Handle("/", http.FileServer(http.Dir("./client/dist/"))).Methods("GET")
	mux.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./client/dist/static/"))))

	// client websocket connection
	mux.HandleFunc("/channel", ws.Handler).Methods("GET")

	return mux
}
