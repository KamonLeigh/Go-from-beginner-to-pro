package main

import (
	"log"
	"net/http"
	"strings"
	"text/template"
)

type Visitor struct {
	Name string
}

type Hello struct {
	tpl *template.Template
}

func (h Hello) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	vl := r.URL.Query()

	vist := Visitor{}

	name, ok := vl["name"]

	if ok {
		vist.Name = strings.Join(name, ",")
	}

	h.tpl.Execute(w, vist)
}

func newHello(tplpath string) (*Hello, error) {
	temp, error := template.ParseFiles(tplpath)

	if error != nil {
		return nil, error
	}

	return &Hello{temp}, nil
}

func main() {
	hello, err := newHello("./index.html")

	if err != nil {
		log.Fatal(err)
	}

	http.Handle("/", hello)

	log.Fatal(http.ListenAndServe(":8081", nil))
}
