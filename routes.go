package main

/*
	Ce fichier sert uniquement à faire une liste exaustif des routes
*/

// ! Créer la func pour le chemin de nos routes
func (s *server) routes() {
	s.router.HandleFunc("/", s.handleIndex()).Methods("GET")
	s.router.HandleFunc("/api/movies/{id:[0-9]+}", s.handleMovieDetail()).Methods("GET")
	s.router.HandleFunc("/api/movies/", s.handleMovieList()).Methods("GET")
	s.router.HandleFunc("/api/movies/", s.handleMovieCreate()).Methods("POST")
}
