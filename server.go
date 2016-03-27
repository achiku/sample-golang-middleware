package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
)

// AppHandler app handler
type AppHandler func(http.ResponseWriter, *http.Request) (int, error)

func (h AppHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	status, err := h(w, r)
	if status == http.StatusInternalServerError {
		w.Header().Set("Content-type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		res := map[string]string{"message": err.Error()}
		json.NewEncoder(w).Encode(res)
	}
}

func hello(w http.ResponseWriter, r *http.Request) (int, error) {
	fmt.Fprintf(w, "hello, world!")
	return http.StatusOK, nil
}

func errorhello(w http.ResponseWriter, r *http.Request) (int, error) {
	return http.StatusInternalServerError, errors.New("error")
}

func main() {
	singleHostHello1 := NewSingleHost(AppHandler(hello), "somehost1")
	singleHostHello2 := SingleHost2(AppHandler(hello), "somehost2")
	appendHello := AppendMiddleware(AppHandler(hello), " append this!")
	mux := http.NewServeMux()
	mux.Handle("/hello", AppHandler(hello))
	mux.Handle("/error/hello", AppHandler(errorhello))
	mux.Handle("/singlehost1/hello", singleHostHello1)
	mux.Handle("/singlehost2/hello", singleHostHello2)
	mux.Handle("/append/hello", appendHello)
	if err := http.ListenAndServe(":8011", mux); err != nil {
		log.Fatal(err)
	}
}
