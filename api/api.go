package api

import (
	"context"
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strconv"
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

func NewPizzaServer() *PizzaServer {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	p := new(PizzaServer)
	router := http.NewServeMux()
	router.Handle("/api/v1/player", http.HandlerFunc(p.playersHandler))
	router.Handle("/api/v1/player/{id}", http.HandlerFunc(p.playersHandlerById))
	router.Handle("/api/v1/pizza/{id}",http.HandlerFunc(p.pizzaHandler))

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

	databaseQueries := database.New(db)
	p.State.DB = databaseQueries
	// Todo: Close the database after the request
	// defer db.Close()

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

		var params database.CreatePlayerParams
		if err := json.NewDecoder(r.Body).Decode(&params); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		params.DateCreated = time.Now()

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
func (p *PizzaServer) playersHandlerById(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		idStr := r.PathValue("id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		player, err := p.State.DB.GetPlayer(context.Background(), int32(id))
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		} else {
			json.NewEncoder(w).Encode(player)

		}
	case http.MethodPost:
		idStr := r.PathValue("id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		var params database.UpdatePlayerParams

		if err := json.NewDecoder(r.Body).Decode(&params); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		params.ID = int32(id)

		player, err := p.State.DB.UpdatePlayer(context.Background(), params)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		} else {
			json.NewEncoder(w).Encode(player)
		}

	case http.MethodDelete:
		idStr := r.PathValue("id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		err = p.State.DB.DeletePlayer(context.Background(), int32(id))
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
	}

}
func (p *PizzaServer) pizzaHandler(w http.ResponseWriter, r *http.Request) {
}