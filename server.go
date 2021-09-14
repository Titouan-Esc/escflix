package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

/*
	Ce fichier sert à la création de notre server et du router et de leur configuration
*/

// ! Correspond au différent composants pour notre projet
type server struct {
	router *mux.Router // Router
	store  Store       // Store qui permet d'accéder à notre BDD
}

// ! Func qui nous renvoie une nouvelle instance de Server
func newServer() *server {
	s := &server{
		router: mux.NewRouter(), // Instancier le router
	}

	s.routes()
	return s
}

// ! Func qui sert à délégué le HTTP au router
func (s *server) serveHTTP(rw http.ResponseWriter, r *http.Request) {
	logRequestMiddleware(s.router.ServeHTTP).ServeHTTP(rw, r) // Délègue le HTTP à notre router gorilla/mux
}

// ! Func qui sert à formatter toute les réponses en format JSON
func (s *server) respond(rw http.ResponseWriter, _ *http.Request, data interface{}, status int) {
	rw.Header().Add("Content-Type", "application/json") // Configuration du Header
	rw.WriteHeader(status)                              // Code de retour de notre réponse

	if data == nil {
		return
	}

	err := json.NewEncoder(rw).Encode(data) // Encoder en json data
	if err != nil {
		log.Printf("Cannot format json. error=%v\n", err)
	}
}

// ! Func qui permet de décoder du JSON
func (s *server) decode(rw http.ResponseWriter, r *http.Request, v interface{}) error {
	return json.NewDecoder(r.Body).Decode(v)

}
