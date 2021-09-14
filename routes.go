package main

/*
	Ce fichier sert uniquement à faire une liste exaustif des routes
*/

// ! Créer la func pour le chemin de nos routes
func (s *server) routes() {
	s.router.HandleFunc("/", s.handleIndex()).Methods("GET")                             // Afficher dans l'index
	s.router.HandleFunc("/api/movies/{id:[0-9]+}", s.handleMovieDetail()).Methods("GET") // Affiche un seul film grâce à sont index
	s.router.HandleFunc("/api/movies/", s.handleMovieList()).Methods("GET")              // Affiche tout les films
	s.router.HandleFunc("/api/movies/", s.handleMovieCreate()).Methods("POST")           // Créer un film
}
