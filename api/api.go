package api

import (
	"context"
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/Dhairya3124/PizzaTime-Backend-Golang/internal/database"
	"github.com/Dhairya3124/PizzaTime-Backend-Golang/state"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type PizzaServer struct {
	http.Handler
	state.State
}

func NewPizzaServer() *PizzaServer {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	p := new(PizzaServer)
	router := http.NewServeMux()
	router.Handle("/api/v1/player", http.HandlerFunc(p.playersHandler))
	p.Handler = router
	p.State = state.New()

	dbURL := os.Getenv("DB_URI")
	if dbURL == "" {
		log.Fatal("DATABASE_URL environment variable is required but not set")
	}
	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatal("Failed to connect to the database:", err)
	}
	defer db.Close()
	databaseQueries := database.New(db)
	p.State.DB = databaseQueries

	return p
}
func (p *PizzaServer) playersHandler(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case http.MethodGet:
		playerData, err := p.State.DB.GetPlayers(context.Background())

		if err != nil {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(err)
		} else {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(playerData)

		}

	case http.MethodPost:
		w.WriteHeader(http.StatusOK)
	}

}
