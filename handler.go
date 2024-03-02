package main

import (
	"html/template"
	"log"
	"net/http"
	"sort"

	"github.com/xlund/badminton-tracker/game"
)

func (a *App) Home(games game.GameList) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tmpl, err := template.ParseFiles("templates/index.html", "templates/games.html", "templates/partials/result.html", "templates/partials/team.html")
		if err != nil {
			log.Fatal(err)
		}
		err = tmpl.Execute(w, games)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func (a *App) GetTiebreaks(games game.GameList) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tmpl, err := template.ParseFiles("templates/games.html", "templates/partials/team.html", "templates/partials/result.html")
		if err != nil {
			log.Fatal(err)
		}
		sort.Slice(games, func(i, j int) bool {
			return games[i].Date.Before(games[j].Date)
		})
		err = tmpl.Execute(w, games)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func (app *App) GamesByPlayer(games game.GameList) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tmpl, err := template.ParseFiles("templates/player.html", "templates/partials/team.html")
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
