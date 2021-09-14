package main

import (
	"log"
	"net/http"
)

/*
	Ce fichier sert à stocker les middlewares
*/

// ! Func qui correspond aux logs des requêtes
func logRequestMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		log.Printf("[%v] %v", r.Method, r.RequestURI)
		next.ServeHTTP(rw, r)
	}
}
