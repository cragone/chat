package main

import (
	"chat/handlers"
	"log"
	"net/http"

	"github.com/bmizerany/pat"
)

func main() {
	mux := routes()

	log.Println("Starting channel listener")
	go handlers.ListenToWSChannel()

	log.Println("Starting web server on port 80")

	_ = http.ListenAndServe(":80", mux)
}

func routes() http.Handler {
	mux := pat.New()

	mux.Get("/", http.HandlerFunc(handlers.Home))
	mux.Get("/ws", http.HandlerFunc(handlers.WsEndpoint))

	fileServer := http.FileServer(http.Dir("/static"))
	mux.Get("/static/", http.StripPrefix("/static", fileServer))

	return mux
}
