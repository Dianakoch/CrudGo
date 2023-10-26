package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func main() {
	bd, err := getDB()
	if err != nil {
		log.Printf("Error with database" + err.Error())
		return
	} else {
		err = bd.Ping()
		if err != nil {
			log.Printf("Error making connection to DB. Please check credentials. The error is: " + err.Error())
			return
		}
	}

	router := mux.NewRouter()
	setupRoutesForBooks(router)

	port := ":8000"

	server := &http.Server{
		Handler:      router,
		Addr:         port,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Printf("Servidor iniciado en %s", port)
	log.Fatal(server.ListenAndServe())
}
