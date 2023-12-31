package main

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Book struct {
	Id    int64  `json:"id"`
	Name  string `json:"name"`
	Genre string `json:"genre"`
	Year  int64  `json:"year"`
}

func stringToInt64(s string) (int64, error) {
	numero, err := strconv.ParseInt(s, 0, 64)
	if err != nil {
		return 0, err
	}
	return numero, err
}

func setupRoutesForBooks(router *mux.Router) {
	enableCORS(router)
	router.HandleFunc("/Books", func(w http.ResponseWriter, r *http.Request) {
		Books, err := getBooks()
		if err == nil {
			respondWithSuccess(Books, w)
		} else {
			respondWithError(err, w)
		}
	}).Methods(http.MethodGet)
	router.HandleFunc("/Book/{id}", func(w http.ResponseWriter, r *http.Request) {
		idAsString := mux.Vars(r)["id"]
		id, err := stringToInt64(idAsString)
		if err != nil {
			respondWithError(err, w)
			return
		}
		Book, err := getBookById(id)
		if err != nil {
			respondWithError(err, w)
		} else {
			respondWithSuccess(Book, w)
		}
	}).Methods(http.MethodGet)

	router.HandleFunc("/Book", func(w http.ResponseWriter, r *http.Request) {
		var Book Book
		err := json.NewDecoder(r.Body).Decode(&Book)
		if err != nil {
			respondWithError(err, w)
		} else {
			err := createBook(Book)
			if err != nil {
				respondWithError(err, w)
			} else {
				respondWithSuccess(true, w)
			}
		}
	}).Methods(http.MethodPost)

	router.HandleFunc("/Book", func(w http.ResponseWriter, r *http.Request) {
		var Book Book
		err := json.NewDecoder(r.Body).Decode(&Book)
		if err != nil {
			respondWithError(err, w)
		} else {
			err := updateBook(Book)
			if err != nil {
				respondWithError(err, w)
			} else {
				respondWithSuccess(true, w)
			}
		}
	}).Methods(http.MethodPut)
	router.HandleFunc("/Book/{id}", func(w http.ResponseWriter, r *http.Request) {
		idAsString := mux.Vars(r)["id"]
		id, err := stringToInt64(idAsString)
		if err != nil {
			respondWithError(err, w)
			return
		}
		err = deleteBook(id)
		if err != nil {
			respondWithError(err, w)
		} else {
			respondWithSuccess(true, w)
		}
	}).Methods(http.MethodDelete)
}

func enableCORS(router *mux.Router) {
	router.PathPrefix("/").HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", AllowedCORSDomain)
	}).Methods(http.MethodOptions)
	router.Use(middlewareCors)
}
func middlewareCors(next http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, req *http.Request) {
			// Just put some headers to allow CORS...
			w.Header().Set("Access-Control-Allow-Origin", AllowedCORSDomain)
			w.Header().Set("Access-Control-Allow-Credentials", "true")
			w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
			w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
			// and call next handler!
			next.ServeHTTP(w, req)
		})
}

// Helper functions for respond with 200 or 500 code
func respondWithError(err error, w http.ResponseWriter) {
	w.WriteHeader(http.StatusInternalServerError)
	json.NewEncoder(w).Encode(err.Error())
}

func respondWithSuccess(data interface{}, w http.ResponseWriter) {

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(data)
}
