package api

import (
	"fmt"
	"net/http"
)
type PizzaServer struct{
	http.Handler
}
func NewPizzaServer()*PizzaServer{
	p:=new(PizzaServer)
	router:=http.NewServeMux()
	router.Handle("/api/v1/player",http.HandlerFunc(p.playersHandler))
	p.Handler = router
	return p
}
func (p *PizzaServer) playersHandler(w http.ResponseWriter,r *http.Request){
	switch r.Method{
	case http.MethodGet:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, `{"message": "Hello, World!"}`)
	case http.MethodPost:
		w.WriteHeader(http.StatusOK)
	}

}