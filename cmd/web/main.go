package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/amphigourik/longbridge/internal/db"
	"github.com/amphigourik/longbridge/internal/handlers"
	"github.com/go-chi/chi/v5"
)

func main() {
	isDev := os.Getenv("APP_ENV") == "development"

	var database *sql.DB
	var err error

	log.Println(isDev)

	if !isDev {
		database, err = db.Connect()
		if err != nil {
			log.Fatal("Database connection failed:", err)
		}
		defer database.Close()
	} else {
		log.Println("Running in development mode, skipping database connection.")
	}

	r := chi.NewRouter()

	fs := http.FileServer(http.Dir("./static"))
	r.Handle("/static/*", http.StripPrefix("/static/", fs))

	r.Get("/", handlers.Home)
	r.Get("/rsvp", handlers.RSVP)
	r.Post("/rsvp", handlers.RSVPPostHandler(database))
	r.Get("/merci", handlers.ThankYou)

	r.Get("/infos", handlers.Info)
	r.Get("/programme", handlers.Agenda)
	r.Get("/contact", handlers.Contact)

	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})

	log.Println("Listening on :8080...")
	err = http.ListenAndServe(":8080", r)
	if err != nil {
		log.Fatal(err)
	}
}
