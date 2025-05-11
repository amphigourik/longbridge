package main

import (
	"log"
	"net/http"

	"github.com/amphigourik/longbridge/internal/handlers"
	"github.com/go-chi/chi/v5"
)

func main() {
	r := chi.NewRouter()

	fs := http.FileServer(http.Dir("./static"))
	r.Handle("/static/*", http.StripPrefix("/static/", fs))

	r.Get("/", handlers.Home)
	r.Get("/rsvp", handlers.RSVP)

	log.Println("Listening on :8080...")
	err := http.ListenAndServe(":8080", r)
	if err != nil {
		log.Fatal(err)
	}
}
