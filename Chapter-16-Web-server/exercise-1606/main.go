package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./index.html")
	})

	http.Handle(
		"/statics/",
		http.StripPrefix(
			"/statics/",
			http.FileServer(http.Dir("./statics")), // was ./public
		),
	)

	// http.Handle(
	//    "/statics/",
	//    http.FileServer(http.Dir(".")),  // serve from project root, path /statics/body.css → ./statics/body.css
	// )

	log.Fatal(http.ListenAndServe(":8081", nil))
}
