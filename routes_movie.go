package main

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

/*
	Ce fichier contient toute les requêtes spécifique aux films
*/

// ! Struct dédié au formatage JSON
type jsonMovie struct {
	ID          int64  `json:"id"`
	Title       string `json:"title"`
	ReleaseDate string `json:"release_date"`
	Duration    int    `json:"duration"`
	TrailerURL  string `json:"trailer_url"`
}

// ! Func pour retourner tout les films
func (s *server) handleMovieList() http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		movies, err := s.store.GetMovies() // Récupérer la liste des films
		if err != nil {
			log.Printf("Cannot load movies. err=%v", err)
			s.respond(rw, r, nil, http.StatusInternalServerError)
			return
		}

		var resp = make([]jsonMovie, len(movies))

		// ? Convertir les éléments de movies en jsonMovies
		for i, m := range movies {
			resp[i] = mapMovieToJson(m)
		}

		// ? Formater en JSON
		s.respond(rw, r, resp, http.StatusOK)
	}
}

// ! Func qui va retourner un seul film
func (s *server) handleMovieDetail() http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id, err := strconv.ParseInt(vars["id"], 10, 64)
		if err != nil {
			log.Printf("Cannot parse id to int. err=%v\n", err)
			s.respond(rw, r, nil, http.StatusBadRequest)
			return
		}

		m, err := s.store.GetMovieById(id)
		if err != nil {
			log.Printf("Cannot load movie. err=%v\n", err)
			s.respond(rw, r, nil, http.StatusInternalServerError)
			return
		}

		var resp = mapMovieToJson(m)
		s.respond(rw, r, resp, http.StatusOK)
	}
}

// ! Func pour créer un film
func (s *server) handleMovieCreate() http.HandlerFunc {
	// ? Le struct permet de dire quel champ sont attendu pour la création du film
	type request struct {
		Title       string `json:"title"`
		ReleaseDate string `json:"release_date"`
		Duration    int    `json:"duration"`
		TrailerURL  string `json:"trailer_url"`
	}

	return func(rw http.ResponseWriter, r *http.Request) {
		req := request{}
		err := s.decode(rw, r, &req)
		if err != nil {
			log.Printf("Cannot parse movie body. err=%v\n", err)
			s.respond(rw, r, nil, http.StatusBadRequest)
			return
		}

		// ? Créer un film
		m := &Movie{
			ID:          0,
			Title:       req.Title,
			ReleaseDate: req.ReleaseDate,
			Duration:    req.Duration,
			TrailerURL:  req.TrailerURL,
		}

		// ? Stocker le film dans la BDD
		err = s.store.CreateMovie(m)
		if err != nil {
			log.Printf("Cannot create movie in DB. err=%v\n", err)
			s.respond(rw, r, nil, http.StatusInternalServerError)
			return
		}

		var resp = mapMovieToJson(m)
		s.respond(rw, r, resp, http.StatusOK)
	}
}

// ! Va Maper le struct Movie en jsonMovie
func mapMovieToJson(m *Movie) jsonMovie {
	return jsonMovie{
		ID:          m.ID,
		Title:       m.Title,
		ReleaseDate: m.ReleaseDate,
		Duration:    m.Duration,
		TrailerURL:  m.TrailerURL,
	}
}
