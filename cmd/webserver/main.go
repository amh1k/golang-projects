package main

import (
	"log"
	"net/http"

	poker "github.com/amh1k/golang-projects"
)
const dbFileName = "game.db.json"
// func (i *InMemoryPlayerStore)GetPlayerScore(name string)int{
// 	return 123
// }
// func (i *InMemoryPlayerStore) RecordWin(name string) {}
func main() {
	store, close, err := poker.FileSystemPlayerStoreFromFile(dbFileName)
	if err != nil {
		log.Fatal(err)
	}
	defer close()
	server := poker.NewPlayerServer(store)
	if err := http.ListenAndServe(":5000", server); err != nil {
		log.Fatalf("could not listen on port 5000 %v", err)
	}
	//http.HandlerFunc is basically a type which is converting PlayerServer to a function recognized by http.ListenAndServer
	// We are basically type casting
	
}