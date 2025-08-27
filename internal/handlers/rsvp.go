package handlers

import (
	"database/sql"
	"net/http"
)

func RSVPPostHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Méthode non autorisée", http.StatusMethodNotAllowed)
			return
		}

		if err := r.ParseForm(); err != nil {
			http.Error(w, "Formulaire invalide", http.StatusBadRequest)
			return
		}

		name := r.FormValue("name")
		email := r.FormValue("email")
		attending := r.FormValue("attending")
		diet := r.FormValue("diet")
		brunch := r.FormValue("brunch")
		gite := r.FormValue("gite")
		message := r.FormValue("message")

		// No duplicate email check

		if name == "" || email == "" || attending == "" || brunch == "" || gite == "" {
			http.Error(w, "Missing required fields", http.StatusBadRequest)
			return
		}

		if attending != "yes" && attending != "no" {
			http.Error(w, "Invalid value for attending", http.StatusBadRequest)
			return
		}
		if brunch != "yes" && brunch != "no" {
			http.Error(w, "Invalid value for brunch participation", http.StatusBadRequest)
			return
		}
		if gite != "yes" && gite != "no" {
			http.Error(w, "Invalid value for gite participation", http.StatusBadRequest)
			return
		}

		attending_bool := attending == "yes"
		brunch_bool := brunch == "yes"
		gite_bool := gite == "yes"

		_, err := db.Exec(`
			INSERT INTO rsvps (name, email, attending, diet, brunch_participation, gite, comment)
			VALUES ($1, $2, $3, $4, $5, $6, $7)`,
			name, email, attending_bool, diet, brunch_bool, gite_bool, message)
		if err != nil {
			http.Error(w, "Server error for RSVP submission: "+err.Error()+"\n"+
				"name: "+name+
				"\nemail: "+email+
				"\nattending: "+attending+
				"\ndiet: "+diet+
				"\nbrunch: "+brunch+
				"\ngite: "+gite+
				"\nmessage: "+message,
				http.StatusInternalServerError)
			return
		}

		http.Redirect(w, r, "/merci", http.StatusSeeOther)
	}
}
