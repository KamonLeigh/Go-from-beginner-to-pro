package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type messageData struct {
	Message string `json:"message"`
}

func getDataAndReturnReaponse() messageData {
	r, err := http.Get("http://localhost:8081")

	if err != nil {
		log.Fatal(err)
	}

	defer r.Body.Close()

	data, err := io.ReadAll(r.Body)

	if err != nil {
		log.Fatal(err)
	}

	message := messageData{}

	err = json.Unmarshal(data, &message)

	if err != nil {
		log.Fatal(err)
	}

	return message

}

func main() {
	data := getDataAndReturnReaponse()
	fmt.Println(data)
}
