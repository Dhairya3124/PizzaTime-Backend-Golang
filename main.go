package main

import (
	"log"
	"net/http"
)
func main(){
	if err := http.ListenAndServe(":5000", nil); err != nil {
		log.Fatalf("could not listen on port 5000 %v", err)
	}
}