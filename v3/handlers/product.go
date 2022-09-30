package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"regexp"
	"strconv"
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
	if r.Method == http.MethodPost {
		p.addProducts(w, r)
	}
	if r.Method == http.MethodPut {
		p.updateProduct(w, r)
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
func (p *Product) updateProduct(w http.ResponseWriter, r *http.Request) {
	reg := regexp.MustCompile(`/([0-9]+)`)
	g := reg.FindAllStringSubmatch(r.URL.Path, -1)
	if len(g) != 1 {
		http.Error(w, "Invalid URL", http.StatusBadRequest)
		return
	}
	if len(g[0]) != 2 {
		http.Error(w, "Invalid URL", http.StatusBadRequest)
		return
	}
	idString := g[0][1]
	id, err := strconv.Atoi(idString)
	if err != nil {
		http.Error(w, "Cannot ID into number", http.StatusBadRequest)
		return
	}
	prod := &data.Product{}
	if err := prod.FromJson(r.Body); err != nil {
		http.Error(w, "bad request from add Prod", http.StatusBadRequest)
		return
	}
	err = prod.UpdateProd(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	e := json.NewEncoder(w)
	e.Encode(prod)
	w.WriteHeader(http.StatusOK)
}
func (p *Product) addProducts(w http.ResponseWriter, r *http.Request) {
	prod := &data.Product{}

	if err := prod.FromJson(r.Body); err != nil {
		http.Error(w, "bad request from add Prod", http.StatusBadRequest)
	}
	data.AddProd(prod)
}
func (p *Product) getProducts(w http.ResponseWriter, r *http.Request) {
	lp := data.GetProducts()
	if err := lp.ToJson(w); err != nil {
		w.Write([]byte("error while encoding"))
	}
}
