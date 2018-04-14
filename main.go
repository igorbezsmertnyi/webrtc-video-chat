package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/igorbezsmertnyi/webrtc-video-chat/db"
	"github.com/igorbezsmertnyi/webrtc-video-chat/routes"
	"github.com/urfave/negroni"
)

func determineListenAddress() (string, error) {
	port := os.Getenv("PORT")
	if port == "" {
		return ":3000", fmt.Errorf("$PORT not set")
	}
	return ":" + port, nil
}

func connectDatabase() {
	url := os.Getenv("DATABASE_URL")

	if url == "" {
		url = "user=postgres dbname=postgres sslmode=disable"
	}

	_, err := db.Connect(url)

	if err != nil {
		log.Fatalf("Connection error: %s", err.Error())
	}
}

func main() {
	connectDatabase()

	addr, _ := determineListenAddress()
	routes := routes.NewRoutes()
	n := negroni.Classic()
	n.UseHandler(routes)

	s := &http.Server{
		Addr:           addr,
		Handler:        n,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	log.Fatal(s.ListenAndServe())
}
