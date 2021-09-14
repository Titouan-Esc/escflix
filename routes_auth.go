package main

import (
	"fmt"
	"net/http"
)

/*
	Ce fichier sert à créer les fonctions qui vont être appelé dans nos chemin de route
*/

// ! Func handleIndex
func (s *server) handleIndex() http.HandlerFunc { // handleIndex() nous renvoie une fonction
	return func(rw http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(rw, "Welcome to Goflix")
	}
}
