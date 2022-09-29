package handlers

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type hello struct {
	l *log.Logger
}

func NewHello(l *log.Logger) *hello {
	return &hello{
		l: l,
	}
}

// hello became part of Handler Interface
func (h *hello) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.l.Print("logging")
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		// w.WriteHeader(http.StatusBadRequest)
		// w.Write([]byte("bad request"))
		// OR
		http.Error(w, "bad req", http.StatusBadRequest)
		return
	}
	fmt.Fprintln(w, string(body))
}
