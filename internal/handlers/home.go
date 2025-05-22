package handlers

import (
	"net/http"

	"github.com/amphigourik/longbridge/internal/views"
)

func Hero(w http.ResponseWriter, r *http.Request) {
	views.RenderTemplate(w, "hero.html", nil)
}

func Home(w http.ResponseWriter, r *http.Request) {
	views.RenderTemplate(w, "home.html", nil)
}

func RSVP(w http.ResponseWriter, r *http.Request) {
	views.RenderTemplate(w, "rsvp.html", nil)
}
