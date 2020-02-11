package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	//handle function on this path
	//takes the fuction and creates an http handler from it and adds it to the default servemux
	//then when i call ListenAndServe
	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		//responsewriter is an interface used by the http response
		//request is the actual request with the method
		log.Println("Hello World")
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

	})
	//this will execute when we curl 127.0.0.1:8080/goodbye
	http.HandleFunc("/goodbye", func(http.ResponseWriter, *http.Request) {
		log.Println("peace out world")

	})
	//constructs an http server and registers a default handler to it.
	// if i put in null then it will use servemux https://golang.org/src/net/http/server.go?s=61509:61556#L2378
	http.ListenAndServe(":8080", nil) //(port to lisen on, handler)
}
