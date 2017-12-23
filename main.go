package main

import (
	"log"
	"net/http"
	"os"

	"bitbucket.org/bitcreator/kubernetes-micro/handlers"
	"bitbucket.org/bitcreator/kubernetes-micro/version"
	"os/signal"
	"syscall"
	"context"
)

func main()  {
	log.Printf("Starting the service. \nVersion: %s - %s (%s)", version.Release, version.BuildTime, version.Commit)

	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("Port is not set.")
	}

	router := handlers.Router(version.BuildTime, version.Commit, version.Release)

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	srv := &http.Server{
		Addr: ":" + port,
		Handler: router,
	}

	go func() {
		log.Fatal(http.ListenAndServe(":" + port, router))
	}()

	log.Print("The service is ready to listen and serve.")


	killSignal := <-interrupt
	switch killSignal {
	case os.Interrupt:
		log.Print("Got SIGINT signal.")
	case syscall.SIGTERM:
		log.Print("Got SIGTERM signal.")
		
	}

	log.Print("Service is shutting down.")
	srv.Shutdown(context.Background())
	log.Print("Service is shut down.")
}
