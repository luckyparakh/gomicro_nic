package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	// will serve to any thing path except /bye
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Print("Hello world")
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			// w.WriteHeader(http.StatusBadRequest)
			// w.Write([]byte("bad request"))
			// OR
			http.Error(w, "bad req", http.StatusBadRequest)
			return
		}
		fmt.Fprint(w, body)
	})

	http.HandleFunc("/bye", func(w http.ResponseWriter, r *http.Request) {
		log.Print("byee")
	})

	// ListenAndServe starts an HTTP server with a given address and handler. The handler is usually nil,
	// which means to use DefaultServeMux. Handle and HandleFunc add handlers to DefaultServeMux:
	http.ListenAndServe(":9090", nil)
}
