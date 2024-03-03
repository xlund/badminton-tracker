package handlers

import (
	"net/http"
	"text/template"

	"github.com/xlund/badminton-tracker/internal/game"
)

func LoadGames() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		games := game.GameList{
			game.Game{ID: 1, GameType: game.Singles, Result: game.Result{Winner: game.Team{Score: 21}, Loser: game.Team{Score: 12}}},
			game.Game{ID: 2, GameType: game.Doubles, Result: game.Result{Winner: game.Team{Score: 21}, Loser: game.Team{Score: 12}}},
			game.Game{ID: 3, GameType: game.Singles, Result: game.Result{Winner: game.Team{Score: 21}, Loser: game.Team{Score: 12}}},
		}
		println(games)
		tmpl, err := template.ParseFiles("web/templates/index.html", "web/templates/games.html", "web/templates/partials/team.html")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		tmpl.Execute(w, games)
	}
}
