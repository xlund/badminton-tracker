package main

import (
	"html/template"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/passageidentity/passage-go"
)

func (app *App) dashboardHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tmpl, err := template.New("index.html").ParseFiles("index.html")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		user := getUser(r.Context())
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		err = tmpl.ExecuteTemplate(w, "index.html", user)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

func (app *App) signOutHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		uid := getUser(r.Context())
		println(os.Getenv("PASSAGE_API_KEY"))
		psg, err := passage.New("bYQcMjnqyyPFrud9PPbobIV7", &passage.Config{
			APIKey: os.Getenv("PASSAGE_API_KEY"),
		})
		if err != nil {
			log.Default().Println("Failed to create Passage client: ", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		_, err = psg.SignOut(uid)
		if err != nil {
			log.Default().Println("Failed to sign out user: ", uid, "\n with error: ", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		c := &http.Cookie{
			Name:     "psg_auth_token",
			Value:    "",
			Expires:  time.Unix(0, 0),
			HttpOnly: true,
		}
		http.SetCookie(w, c)
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}

func (app *App) loginHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tmpl, err := template.New("login.html").ParseFiles("login.html")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		err = tmpl.ExecuteTemplate(w, "login.html", nil)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

func (app *App) favicon() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "static/favicon.ico")
	}
}
