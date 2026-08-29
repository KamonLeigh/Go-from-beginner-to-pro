package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

var names []Name

type Name struct {
	Name string `json:"name"`
}

type Names struct {
	Names []Name `json:"names"`
}

type Response struct {
	Ok bool `json:"ok"`
}

func getNames(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	data := Names{}

	data.Names = append(data.Names, names...)

	json, err := json.Marshal(data)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json")

	w.Write(json)

}

func addNames(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	name := Name{}

	defer r.Body.Close()

	data, err := io.ReadAll(r.Body)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = json.Unmarshal(data, &name)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	names = append(names, name)

	res := Response{
		Ok: true,
	}

	json, err := json.Marshal(res)

	if err != nil {
		http.Error(w, "Method not allowed", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(json)
}

func main() {
	http.HandleFunc("/names", getNames)
	http.HandleFunc("/", addNames)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
