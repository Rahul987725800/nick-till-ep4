package data

import (
	"encoding/json"
	"fmt"
	"io"
	"time"
)

type Product struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	SKU         string  `json:"sku"`
	CreatedOn   string  `json:"-"`
	UpdatedOn   string  `json:"-"`
	DeletedOn   string  `json:"-"`
}

func (p *Product) FromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(p)
}

type Products []*Product

func (p *Products) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(p)
}
func GetProducts() Products {
	return productList
}
func AddProduct(p *Product) *Product {
	p.ID = getNextID()
	productList = append(productList, p)
	return p
}
func UpdateProduct(id int, p *Product) (*Product, error) {
	index, err := findIndexOfProductById(id)
	if err != nil {
		return nil, err
	}
	p.ID = id
	productList[index] = p
	return p, nil
}
func ModifyProduct(id int, p *Product) (*Product, error) {
	index, err := findIndexOfProductById(id)
	if err != nil {
		return nil, err
	}
	existingProd := productList[index]
	if p.Name != "" {
		existingProd.Name = p.Name
	}
	if p.Description != "" {
		existingProd.Description = p.Description
	}
	if p.Price != 0 {
		existingProd.Price = p.Price
	}
	if p.SKU != "" {
		existingProd.SKU = p.SKU
	}
	return existingProd, nil
}
func DeleteProduct(id int) (*Product, error) {
	index, err := findIndexOfProductById(id)
	if err != nil {
		return nil, err
	}
	removed := productList[index]
	for i := index; i < len(productList)-1; i++ {
		productList[i] = productList[i+1]
	}
	productList = productList[:len(productList)-1]
	return removed, nil
}
func GetProduct(id int) (*Product, error) {
	index, err := findIndexOfProductById(id)
	if err != nil {
		return nil, err
	}
	prod := productList[index]
	return prod, nil
}
func findIndexOfProductById(id int) (int, error) {
	index := -1
	for i, v := range productList {
		if v.ID == id {
			index = i
		}
	}
	if index == -1 {
		return index, fmt.Errorf("Product with id %d doesn't exist", id)
	}
	return index, nil
}
func (p *Product) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(p)
}
func getNextID() int {
	lp := productList[len(productList)-1]
	return lp.ID + 1
}

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
