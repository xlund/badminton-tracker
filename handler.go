package main

import (
	"html/template"
	"log"
	"net/http"

	"github.com/xlund/badminton-tracker/game"
)

func (a *App) Home(games game.GameList) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tmpl, err := template.ParseFiles("templates/index.html")
		if err != nil {
			log.Fatal(err)
		}
		err = tmpl.Execute(w, games)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func (app *App) favicon() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "static/favicon.ico")
	}
}
