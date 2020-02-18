package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/bizahmad/go-Microservices/data"
)

type Products struct {
	l *log.Logger
}

func NewProduct(l *log.Logger) *Products {
	return &Products{l}
}

func (p *Products) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	lp := data.GetProducts()
	d, err := json.Marshal(lp)

	if err != nil {
		http.Error(rw, "Unable to marshal jason", http.StatusInternalServerError)
	}

	rw.Write(d)
}
