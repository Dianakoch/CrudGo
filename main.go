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
		log.Printf("Error con la bd" + err.Error())
		return
	} else {
		err = bd.Ping()
		if err != nil {
			log.Printf("Error conectando con la base de datos, verifica credenciales, error: " + err.Error())
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
