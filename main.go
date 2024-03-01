package main

import (
	"flag"
	"net/http"

	"github.com/xlund/badminton-tracker/game"
)

type App struct {
}

func main() {
	app := App{}

	file := flag.String("file", "data.csv", "The file to parse")
	flag.Parse()

	games := game.CsvParser(*file)
	result := games.Filter(hasResult)

	mux := http.NewServeMux()
	mux.HandleFunc("GET /", app.Home(result))

	println("Application launched and running on http://localhost:3000")
	http.ListenAndServe("localhost:3000", mux)

}

func hasResult(g game.Game) bool {
	return g.HasResult()
}
