package main

import (
	"flag"
	"net/http"

	"github.com/xlund/badminton-tracker/game"
)

type App struct{}

func main() {
	app := App{}

	file := flag.String("file", "data.csv", "The file to parse")
	flag.Parse()

	fs := http.FileServer(http.Dir("./static"))

	games := game.CsvParser(*file)
	result := games.Filter(hasResult)

	mux := http.NewServeMux()
	mux.HandleFunc("GET /", app.Home(result[0:10]))
	mux.Handle("GET /static/", http.StripPrefix("/static/", fs))

	println("Application launched and running on http://localhost:3000")
	http.ListenAndServe("localhost:3000", mux)

}

func hasResult(g game.Game) bool {
	return g.HasResult()
}
