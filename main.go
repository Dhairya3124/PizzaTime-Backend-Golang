package main

import (
	"log"
	"net/http"

	"github.com/Dhairya3124/PizzaTime-Backend-Golang/api"
)
func main(){
	p:=api.NewPizzaServer()
	if err := http.ListenAndServe(":5000", p); err != nil {
		log.Fatalf("could not listen on port 5000 %v", err)
	}
}