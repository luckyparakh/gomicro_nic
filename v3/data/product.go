package data

import (
	"encoding/json"
	"errors"
	"io"
	"time"
)

// Product defines the structure for an API product
type Product struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float32 `json:"price"`
	SKU         string  `json:"sku"`
	CreatedOn   string  `json:"-"`
	UpdatedOn   string  `json:"-"`
	DeletedOn   string  `json:"-"`
}

func (p *Product) FromJson(r io.Reader) error {
	d := json.NewDecoder(r)
	return d.Decode(p)
}

type Products []*Product

func GetProducts() Products {
	return productList
}

func (p *Products) ToJson(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(p)
}
func AddProd(p *Product) {
	p.ID = getId()
	productList = append(productList, p)
}
func getId() int {
	return len(productList) + 1
}

var errProdNotFound = errors.New("Product not found")

func (p *Product) UpdateProd(id int) error {
	foundId := findProd(id)
	if foundId == -1 {
		return errProdNotFound
	}
	p.ID = id
	productList[foundId] = p
	return nil
}

func findProd(id int) int {
	for i, prod := range productList {
		if prod.ID == id {
			return i
		}
	}
	return -1
}

// productList is a hard coded list of products for this
// example data source
var productList = []*Product{
	&Product{
		ID:          1,
		Name:        "Latte",
		Description: "Frothy milky coffee",
		Price:       2.45,
		SKU:         "abc323",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
	&Product{
		ID:          2,
		Name:        "Espresso",
		Description: "Short and strong coffee without milk",
		Price:       1.99,
		SKU:         "fjd34",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
}
