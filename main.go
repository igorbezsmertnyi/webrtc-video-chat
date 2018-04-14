package main

import (
	"fmt"
	"os"

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

func main() {
	addr, _ := determineListenAddress()
	routes := routes.NewRoutes()
	n := negroni.Classic()
	n.UseHandler(routes)
	n.Run(addr)
}
