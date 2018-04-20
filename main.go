package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"webrtc-video-chat/models"
	"webrtc-video-chat/routes"
	"webrtc-video-chat/ws"

	"github.com/rs/cors"
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

	db, err := models.Connect(url)

	if err != nil {
		log.Fatalf("Connection error: %s", err.Error())
	}

	models.SetDatabase(db)
}

func main() {
	go connectDatabase()

	hub := ws.H

	go hub.Run()

	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
	})

	addr, _ := determineListenAddress()
	routes := routes.NewRoutes()
	n := negroni.Classic()
	n.Use(c)
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
