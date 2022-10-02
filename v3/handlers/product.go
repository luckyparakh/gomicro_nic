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

func (p *Product) UpdateProduct(w http.ResponseWriter, r *http.Request) {
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
func (p *Product) AddProducts(w http.ResponseWriter, r *http.Request) {
	prod := &data.Product{}

	if err := prod.FromJson(r.Body); err != nil {
		http.Error(w, "bad request from add Prod", http.StatusBadRequest)
	}
	data.AddProd(prod)
}
func (p *Product) GetProducts(w http.ResponseWriter, r *http.Request) {
	lp := data.GetProducts()
	if err := lp.ToJson(w); err != nil {
		w.Write([]byte("error while encoding"))
	}
}
