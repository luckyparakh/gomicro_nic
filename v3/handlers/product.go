package handlers

import (
	"log"
	"net/http"
	"v3/data"
)

type Product struct {
	l *log.Logger
}

func NewProduct(l *log.Logger) *Product {
	return &Product{l}
}

// Default MUX by default calls ServeHTTP
func (p *Product) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		p.getProducts(w, r)
	}
	w.WriteHeader(http.StatusNotImplemented)
	// OR
	// d, err := json.Marshal(lp)
	// if err != nil {
	// 	http.Error(w, "Can't marshal to json", http.StatusBadRequest)
	// }
	// fmt.Fprint(w, string(d))
	// // w.Write(d)
}

func (p *Product) getProducts(w http.ResponseWriter, r *http.Request) {
	lp := data.GetProducts()
	if err := lp.ToJson(w); err != nil {
		w.Write([]byte("error while encoding"))
	}
}
