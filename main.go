package main

import (
	"log"
	"net/http"
	"os"

	"github.com/bizahmad/go-Microservices/handlers"
)

func main() {
	l := log.New(os.Stdout, "product-api", log.LstdFlags)
	hh := handlers.NewHello(l)
	gh := handlers.NewGoodbye(l)

	//http.HandleFunc() were converting this func into a handler type and registering it on the default servmux
	sm := http.NewServeMux()
	sm.Handle("/", hh)
	sm.Handle("/goodbye", gh)
	http.ListenAndServe(":8080", sm)
}
