package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
)

type Name struct {
	Name string `json:"name"`
}

type Names struct {
	Names []Name `json:"names"`
}

type Response struct {
	Ok bool `json:"ok"`
}

var url = "http://localhost:8080"

func addNameAndParseResponse(nameToAdd string) error {
	name := Name{
		Name: nameToAdd,
	}

	jsonBytes, err := json.Marshal(name)

	if err != nil {
		return err
	}

	r, err := http.Post(fmt.Sprintf("%s/", url), "application/json", bytes.NewBuffer(jsonBytes))

	if err != nil {
		return err
	}

	defer r.Body.Close()
	data, err := io.ReadAll(r.Body)

	if err != nil {
		return err
	}

	res := Response{}

	err = json.Unmarshal(data, &res)

	if err != nil {
		return err
	}

	if !res.Ok {
		return errors.New("User POST failed")
	}

	return nil
}

func getDataAndParseResponse() []Name {

	r, err := http.Get(fmt.Sprintf("%s/names", url))

	if err != nil {
		log.Fatal(err)
	}

	defer r.Body.Close()

	data, err := io.ReadAll(r.Body)

	if err != nil {
		log.Fatal(err)
	}

	names := Names{}

	err = json.Unmarshal(data, &names)

	if err != nil {
		log.Fatal(err)
	}

	return names.Names
}

func main() {
	err := addNameAndParseResponse("electronic")

	if err != nil {
		log.Fatal(err)
	}

	err = addNameAndParseResponse("device")

	if err != nil {
		log.Fatal(err)
	}

	data := getDataAndParseResponse()

	for _, name := range data {
		log.Println(name.Name)
	}

}
