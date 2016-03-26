package main

import (
	"fmt"
	"log"
	"net/http"
)

func hello(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "hello, world!")
}

func main() {
	singleHostHello1 := NewSingleHost(http.HandlerFunc(hello), "somehost1")
	singleHostHello2 := SingleHost2(http.HandlerFunc(hello), "somehost2")
	mux := http.NewServeMux()
	mux.Handle("/hello", http.HandlerFunc(hello))
	mux.Handle("/singlehost1/hello", singleHostHello1)
	mux.Handle("/singlehost2/hello", singleHostHello2)
	if err := http.ListenAndServe(":8011", mux); err != nil {
		log.Fatal(err)
	}
}
