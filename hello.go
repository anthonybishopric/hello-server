package main

import (
	"fmt"
	"html"
	"log"
	"net/http"
)

func SayHello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
}

func main() {
	http.HandleFunc("/", SayHello)
	s := &http.Server{
		Addr: ":43770",
	}
	log.Fatal(s.ListenAndServe())
}
