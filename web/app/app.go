package app

import (
	"net/http"
	"os"

	"github.com/xlund/badminton-tracker/internal/handlers"
)

func LoadApp() {
	port, exists := os.LookupEnv("SERVER_PORT")
	if !exists {
		port = "3000"
	}
	router := http.NewServeMux()
	handlers.LoadHandlers(router)
	println("Application launched and running on http://localhost:" + port)
	http.ListenAndServe("localhost:"+port, router)
}
