package handlers

import (
	"net/http"

	"github.com/amphigourik/longbridge/internal/views"
)

func Home(w http.ResponseWriter, r *http.Request) {
	views.RenderTemplate(w, "home.html", nil)
}

func Info(w http.ResponseWriter, r *http.Request) {
	views.RenderTemplate(w, "infos.html", nil)
}

func Agenda(w http.ResponseWriter, r *http.Request) {
	views.RenderTemplate(w, "agenda.html", nil)
}

func RSVP(w http.ResponseWriter, r *http.Request) {
	views.RenderTemplate(w, "rsvp.html", nil)
}

func ThankYou(w http.ResponseWriter, r *http.Request) {
	views.RenderTemplate(w, "thankyou.html", nil)
}
