package handlers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/bizahmad/go-Microservices/data"
	"github.com/gorilla/mux"
)

type LoggerProducts struct {
	l *log.Logger
}

func NewProduct(l *log.Logger) *LoggerProducts {
	return &LoggerProducts{l}
}

func (p *LoggerProducts) GetProducts(rw http.ResponseWriter, h *http.Request) {
	lp := data.GetProducts()
	err := lp.ToJSON(rw)

	if err != nil {
		http.Error(rw, "Unable to marshal json", http.StatusInternalServerError)
	}
}

func (p *LoggerProducts) AddProduct(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Hanlde POST Product")

	prod := &data.Product{}

	err := prod.FromJSON(r.Body)
	if err != nil {
		http.Error(rw, "Unable to unmarshal json", http.StatusBadRequest)
	}

	data.AddProduct(prod)
}

func (p *LoggerProducts) UpdateProducts(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	p.l.Println("Handle PUT Product")

	prod := &data.Product{}
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(rw, "Unable to convert id", http.StatusBadRequest)
	}

	err = prod.FromJSON(r.Body)
	if err != nil {
		http.Error(rw, "Unable to unmarshal json", http.StatusBadRequest)
	}

	err = data.UpdateProduct(id, prod)
	if err == data.ErrProductNotFound {
		http.Error(rw, "Product not found", http.StatusNotFound)
		return
	}

	if err != nil {
		http.Error(rw, "Product not found", http.StatusNotFound)
		return
	}

}
