package main

import (
	"log"
	"net/http"
)
type InMemoryPlayerStore struct{}

func (i *InMemoryPlayerStore)GetPlayerScore(name string)int{
	return 123
}
func main() {
	server := &PlayerServer{&InMemoryPlayerStore{}}
	//http.HandlerFunc is basically a type which is converting PlayerServer to a function recognized by http.ListenAndServer
	// We are basically type casting
	log.Fatal(http.ListenAndServe(":5000", server))
}