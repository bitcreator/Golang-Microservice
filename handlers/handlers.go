package handlers

import (
	"github.com/gorilla/mux"
	"sync/atomic"
	"log"
	"time"
)

func Router(buildTime, commit, release string) *mux.Router{
	isReady := &atomic.Value{}
	isReady.Store(false)

	go func() {
		log.Printf("Ready probe is negative.")

		time.Sleep(1 * time.Second)
		isReady.Store(true)

		log.Printf("Ready probe is positive.")
	}()

	r := mux.NewRouter()
	r.HandleFunc("/version", version(buildTime, commit, release)).Methods("GET")
	r.HandleFunc("/healthz", healthz)
	r.HandleFunc("/readyz", readyz(isReady))

	return r
}