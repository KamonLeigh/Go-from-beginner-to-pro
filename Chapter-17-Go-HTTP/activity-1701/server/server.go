package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type ResultData struct {
	Electric int
	Boogaloo int
}

type Db struct {
	Names []string `json:"names"`
}

type server struct{}

func (srv server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	db := Db{Names: []string{"Electric", "Electric", "Electric", "Boogaloo", "Boogaloo", "Boogaloo", "Boogaloo"}}

	result := ResultData{}

	for _, value := range db.Names {
		if value == "Electric" {
			result.Electric++
		}

		if value == "Boogaloo" {
			result.Boogaloo++
		}
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)

}

func main() {
	log.Fatal(http.ListenAndServe(":8081", server{}))
}
