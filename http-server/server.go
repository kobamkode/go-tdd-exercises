package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

type PlayerStore interface {
	GetPlayerScore(name string) int
}

type PlayerServer struct {
	store PlayerStore
}

func (p *PlayerServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	player := strings.TrimPrefix(r.URL.Path, "/players/")
	fmt.Fprint(w, p.store.GetPlayerScore(player))
}

func GetPlayerScore(player string) string {
	if player == "Pepper" {
		return "20"
	}

	if player == "Floyd" {
		return "10"
	}
	return ""
}

func main() {
	server := &PlayerServer{}
	log.Fatal(http.ListenAndServe(":5000", server))
}
