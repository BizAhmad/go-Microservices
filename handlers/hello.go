package handlers

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

//Hello ...
type Hello struct {
	l *log.Logger
}

//NewHello ...
func NewHello(l *log.Logger) *Hello {
	return &Hello{l}
}

func (h *Hello) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	//responsewriter is an interface used by the http response
	//request is the actual request with the method

	h.l.Println("Hello World")
	//set the body of the request into d and dont worry about any errors
	d, err := ioutil.ReadAll(r.Body)
	if err != nil {
		// rw.WriteHeader(http.StatusBadRequest)
		// rw.Write([]byte("Oops"))
		// return
		// instead you can do this:
		http.Error(rw, "Oops", http.StatusBadRequest)
		return
	}
	fmt.Fprintf(rw, "Data:  %s, \n", d)
}
