package main

import (
	"log"
	"net/http"
)

// func (i *InMemoryPlayerStore)GetPlayerScore(name string)int{
// 	return 123
// }
// func (i *InMemoryPlayerStore) RecordWin(name string) {}
func main() {
	server := &PlayerServer{NewInMemoryPlayerStore()}
	log.Fatal(http.ListenAndServe(":5000", server))
	//http.HandlerFunc is basically a type which is converting PlayerServer to a function recognized by http.ListenAndServer
	// We are basically type casting
	
}