package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/mux"

	"github.com/bizahmad/go-Microservices/handlers"
)

func main() {
	l := log.New(os.Stdout, "product-api", log.LstdFlags)
	//create the handlers
	ph := handlers.NewProduct(l)

	//router has concept of subrooters
	sm := mux.NewRouter()

	getRouter := sm.Methods("GET").Subrouter()
	getRouter.HandleFunc("/", ph.GetProducts)

	putRooter := sm.Methods(http.MethodPut).Subrouter()
	putRooter.HandleFunc("/{id:[0-9]+}", ph.UpdateProducts)

	postRooter := sm.Methods(http.MethodPost).Subrouter()
	postRooter.HandleFunc("/", ph.AddProduct)
	//sm.Handle("/products", ph)

	//creating our own server config
	//More control over the server's behavior is available by creating a custom Server:
	s := &http.Server{
		Addr:         ":3030",
		Handler:      sm,
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}

	//start the server
	go func() {
		err := s.ListenAndServe()
		if err != nil {
			l.Printf("Error starting %s\n", err)
			os.Exit(1)
		}
	}()

	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, os.Interrupt) //whenever a signal is received by the system itll send the message on that channel
	signal.Notify(sigChan, os.Kill)      //trap a forceful kill to the program

	// Block until a signal is received.
	sig := <-sigChan
	l.Println("Received terminate, gracefgul shutdown", sig)

	// if the handlers are working after 30 seconds then shut them down forcefully
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	s.Shutdown(ctx)
}
