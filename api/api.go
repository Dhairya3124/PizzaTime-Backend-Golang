package api

import (
	"context"
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/Dhairya3124/PizzaTime-Backend-Golang/internal/database"
	"github.com/Dhairya3124/PizzaTime-Backend-Golang/state"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type PizzaServer struct {
	http.Handler
	state.State
}
type CreatePlayerParams struct {
	Name sql.NullString 
Age  sql.NullInt32
DateCreated sql.NullTime
Gender      sql.NullString
TotalPizza  sql.NullInt32
LoggedPizza sql.NullInt32
Coins sql.NullInt32


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
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		
		var params CreatePlayerParams
		if err := json.NewDecoder(r.Body).Decode(&params); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		params.DateCreated = sql.NullTime{Time: time.Now(), Valid: true}
		
		playerParams := database.CreatePlayerParams{
			Name:        params.Name,
			Age:         params.Age,
			DateCreated: params.DateCreated,
			Gender:      params.Gender,
			TotalPizza:  params.TotalPizza,
			LoggedPizza: params.LoggedPizza,
			Coins:       params.Coins,
		}
		player, err := p.State.DB.CreatePlayer(context.Background(), playerParams)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		json.NewEncoder(w).Encode(player)
		
	}

}
