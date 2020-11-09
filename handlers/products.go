package handlers

import (
	"fmt"
	"log"
	"net/http"
	"nick/data"
	"regexp"
	"strconv"
)

type Products struct {
	l *log.Logger
}

func NewProducts(l *log.Logger) *Products {
	return &Products{l}
}
func (p *Products) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		if r.URL.Path == "/" {
			p.getProducts(w, r)
			return
		} else {
			id, err := getIdFromURL(r.URL.Path)
			if err != nil {
				p.l.Println(err)
				http.Error(w, "Invalid URL", http.StatusBadRequest)
				return
			}
			p.getProduct(id, w, r)
			return
		}

	}
	if r.Method == http.MethodPost {
		p.addProduct(w, r)
		return
	}
	if r.Method == http.MethodPut {
		// expect the id in URI
		// p.l.Println("PUT", r.URL.Path)
		id, err := getIdFromURL(r.URL.Path)
		if err != nil {
			p.l.Println(err)
			http.Error(w, "Invalid URL", http.StatusBadRequest)
			return
		}
		// p.l.Println("got id", id)
		p.updateProduct(id, w, r)
		return
	}
	if r.Method == http.MethodPatch {
		id, err := getIdFromURL(r.URL.Path)
		if err != nil {
			p.l.Println(err)
			http.Error(w, "Invalid URL", http.StatusBadRequest)
			return
		}
		p.modifyProduct(id, w, r)
		return
	}
	if r.Method == http.MethodDelete {
		id, err := getIdFromURL(r.URL.Path)
		if err != nil {
			p.l.Println(err)
			http.Error(w, "Invalid URL", http.StatusBadRequest)
			return
		}
		p.deleteProduct(id, w, r)
		return
	}
	// catch all
	w.WriteHeader(http.StatusMethodNotAllowed)
}
func getIdFromURL(path string) (int, error) {
	// expect the id in URI
	// p.l.Println("PUT", r.URL.Path)
	reg := regexp.MustCompile(`/([0-9]+)`)
	g := reg.FindAllStringSubmatch(path, -1)
	// p.l.Println("PUT", g)
	// g = [[/32 32]]
	if len(g) != 1 {

		return -1, fmt.Errorf("Invalid URL more than one id")
		// will be raised if /89/90
	}
	if len(g[0]) != 2 {

		return -1, fmt.Errorf("Invalid URL more than one capture group")
	}
	idString := g[0][1]
	// p.l.Println("PUT", idString)
	id, err := strconv.Atoi(idString)
	if err != nil {

		return -1, fmt.Errorf("Invalid URL can't convert id to int type")
	}
	return id, nil
}
func (p *Products) getProducts(w http.ResponseWriter, r *http.Request) {
	lp := data.GetProducts()
	err := lp.ToJSON(w)
	if err != nil {
		http.Error(w, "Unable to marshal json", http.StatusInternalServerError)
	}
}
func (p *Products) getProduct(id int, w http.ResponseWriter, r *http.Request) {
	prod, err := data.GetProduct(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	_ = prod.ToJSON(w)

}
func (p *Products) addProduct(w http.ResponseWriter, r *http.Request) {
	prod := &data.Product{}
	err := prod.FromJSON(r.Body)
	if err != nil {
		http.Error(w, "Unable to unmarshal json", http.StatusBadRequest)
		return
	}
	// p.l.Printf("Prod: %#v", prod)
	savedProd := data.AddProduct(prod)
	_ = savedProd.ToJSON(w)
}
func (p *Products) updateProduct(id int, w http.ResponseWriter, r *http.Request) {
	prod := &data.Product{}
	err := prod.FromJSON(r.Body)
	if err != nil {
		http.Error(w, "Unable to unmarshal json", http.StatusBadRequest)
		return
	}
	updatedProd, err := data.UpdateProduct(id, prod)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	_ = updatedProd.ToJSON(w)
}
func (p *Products) modifyProduct(id int, w http.ResponseWriter, r *http.Request) {
	prod := &data.Product{}
	err := prod.FromJSON(r.Body)
	if err != nil {
		http.Error(w, "Unable to unmarshal json", http.StatusBadRequest)
		return
	}
	modifiedProd, err := data.ModifyProduct(id, prod)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	_ = modifiedProd.ToJSON(w)
}
func (p *Products) deleteProduct(id int, w http.ResponseWriter, r *http.Request) {
	deletedProduct, err := data.DeleteProduct(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	_ = deletedProduct.ToJSON(w)
}
