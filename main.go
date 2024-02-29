package main

import (
	"flag"

	"github.com/passageidentity/passage-go"
	"github.com/xlund/badminton-tracker/game"
)

type App struct {
	psg *passage.App
}

func main() {
	// psgApiKey := os.Getenv("PASSAGE_API_KEY")
	// if psgApiKey == "" {
	// 	fmt.Println("PASSAGE_API_KEY environment variable not set")
	// 	return
	// }
	// psg, err := passage.New("bYQcMjnqyyPFrud9PPbobIV7", &passage.Config{
	// 	APIKey: psgApiKey,
	// })
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// app := &App{
	// 	psg,
	// }
	// mux := http.NewServeMux()
	// mux.HandleFunc("GET /", app.sessionMiddleware(app.dashboardHandler()))
	// mux.HandleFunc("GET /sign-out", app.sessionMiddleware(app.signOutHandler()))
	// mux.HandleFunc("GET /login", app.loginHandler())
	// mux.HandleFunc("GET /favicon.ico", app.favicon())
	// port := os.Getenv("PORT")
	// if port == "" {
	// 	port = "3000"
	// }
	// fmt.Printf("Application launched and running on http://localhost:%s\n", port)
	// http.ListenAndServe("localhost:"+port, mux)
	file := flag.String("file", "data.csv", "The file to parse")
	flag.Parse()
	games := game.CsvParser(*file)
	println(games[12].ResultString())
}
