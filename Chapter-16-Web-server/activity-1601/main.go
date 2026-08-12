package main

import (
	"fmt"
	"log"
	"net/http"
)

type pageWithCounter struct {
	Counter int
	Content string
	Heading string
}

func (p *pageWithCounter) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	p.Counter++
	msg := fmt.Sprintf("<h1>%s</h1>\n<p>%s</p>\n<p>Views: %d", p.Heading, p.Content, p.Counter)
	w.Write([]byte(msg))
}

func main() {

	hom := pageWithCounter{
		0,
		"This is my main page",
		"Hello World",
	}

	chap1 := pageWithCounter{
		0,
		"This is the first chapter",
		"Chapter 1",
	}

	chap2 := pageWithCounter{
		0,
		"This is the second chapter",
		"Chapter 2",
	}

	http.Handle("/", &hom)
	http.Handle("/chapter1", &chap1)
	http.Handle("/chapter2", &chap2)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
