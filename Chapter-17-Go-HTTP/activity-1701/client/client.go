package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type ResponseData struct {
	Electric int
	Boogaloo int
}

func getDataAndReturnResponse() ResponseData {
	r, err := http.Get("http://localhost:8081")

	if err != nil {
		log.Fatal(err)
	}

	defer r.Body.Close()

	data, err := io.ReadAll(r.Body)

	if err != nil {
		log.Fatal(err)
	}

	response := ResponseData{}

	err = json.Unmarshal(data, &response)

	return response
}

func main() {

	data := getDataAndReturnResponse()

	fmt.Printf("Electric count: %d\n", data.Electric)
	fmt.Printf("Boogaloo count: %d", data.Boogaloo)
}
