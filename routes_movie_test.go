package main

import (
	"bytes"
	"encoding/json"
	"log"
	"testing"
)

// ! Simuler un store de film
type testStore struct {
	movieId int64
	movies  []*Movie
}

// ! Créer les func qui sont dans l'interface store

func (t testStore) Open() error { // func Open()
	return nil
}

func (t testStore) Close() error { // func Close()
	return nil
}

func (t testStore) GetMovies() ([]*Movie, error) { // func GetMovies()
	return t.movies, nil
}

func (t testStore) GetMovieById(id int64) (*Movie, error) { // func GetMovieById()
	for _, m := range t.movies {
		if m.ID == id {
			return m, nil
		}
	}
	return nil, nil
}

func (t testStore) CreateMovie(m *Movie) error { // func CreateMovie()
	t.movieId++                    // Incrémenter l'ID de movie
	m.ID = t.movieId               // Affecter l'ID de movie à l'ID incrémenté
	t.movies = append(t.movies, m) // Ajouter le film à notre liste de film
	return nil
}

func TestMovieCreateUnit(t *testing.T) {
	//  Création d'un nouveau server
	srv := newServer()
	srv.store = &testStore{}

	p := struct { // Utilisation d'un struct annonime
		Title       string `json:"title"`
		ReleaseDate string `json:"release_date"`
		Duration    int    `json:"duration"`
		TrailerURL  string `json:"trailer_url"`
	}{
		Title:       "Inception",
		ReleaseDate: "2010-07-18",
		Duration:    148,
		TrailerURL:  "http://url",
	}

	//  Création d'un buffer de bytes
	var buf bytes.Buffer

	// Mettre le contenue encoder en JSON dans le buffer
	err := json.NewEncoder(&buf).Encode(p)
	if err != nil {
		log.Printf("Cannot Encode JSON. error=%v\n", err)
		return
	}
}
