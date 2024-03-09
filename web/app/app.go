package app

import (
	"database/sql"
	"net/http"
	"os"

	_ "github.com/mattn/go-sqlite3"

	"github.com/xlund/badminton-tracker/internal/handlers"
	"github.com/xlund/badminton-tracker/internal/repos"
)

type App struct {
	PlayerRepo *repos.PlayerRepo
}

const dbPath = "./badminton.db"

func LoadApp() {

	os.Remove(dbPath)
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		panic(err)
	}
	playerRepo := repos.NewPlayerRepo(db)
	playerRepo.Migrate()
	port, exists := os.LookupEnv("SERVER_PORT")
	if !exists {
		port = "3000"
	}
	router := http.NewServeMux()
	handlers.LoadHandlers(router)
	println("Application launched and running on http://localhost:" + port)
	http.ListenAndServe("localhost:"+port, router)
}
