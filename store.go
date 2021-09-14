package main

import (
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

/*
	Ce fichier contient toute les fonctions qui vont intéragir avec la BDD
*/

// ! Interface pour accéder à notre BDD
type Store interface {
	Open() error  // Ouvrir la connexion vers notre BDD SQL
	Close() error // Fermer la connexion à notre BDD SQL

	GetMovies() ([]*Movie, error)          // Afficher la liste des films
	GetMovieById(id int64) (*Movie, error) // Récupérer un film avec l'id

	CreateMovie(m *Movie) error // Créer un film
}

// ! Implémenter les fonctions de Store
type dbStore struct {
	db *sqlx.DB
}

// ! Var pour le schema de table
var schema = `
CREATE TABLE IF NOT EXISTS movie
(
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	title TEXT,
	release_date TEXT,
	duration INTEGER,
	trailer_url TEXT
)
`

// ! Func Open()
func (store *dbStore) Open() error {
	db, err := sqlx.Connect("sqlite3", "goflix.db") // Connexion à la BDD
	if err != nil {
		return err
	}

	log.Println("Connected to DB")

	db.MustExec(schema)
	store.db = db // Affilier store à db
	return nil
}

// ! Func Close()
func (store *dbStore) Close() error {
	return store.db.Close()
}

// ! Func GetMovies()
func (store *dbStore) GetMovies() ([]*Movie, error) {
	var movies []*Movie // Variable qui contiendra ce qu'il à était récupérer de la base

	err := store.db.Select(&movies, "SELECT * FROM movie")
	if err != nil {
		return movies, err
	}
	return movies, nil
}

// ! Func pour récupérer un seul film
func (store *dbStore) GetMovieById(id int64) (*Movie, error) {
	var movie = &Movie{}
	err := store.db.Get(movie, "SELECT * FROM movie WHERE id=$1", id)
	if err != nil {
		return movie, nil
	}

	return movie, nil
}

// ! Func de la création d'un film
func (store *dbStore) CreateMovie(m *Movie) error {
	res, err := store.db.Exec("INSERT INTO movie (title, release_date, duration, trailer_url) VALUES (?, ?, ?, ?)",
		m.Title, m.ReleaseDate, m.Duration, m.TrailerURL)
	if err != nil {
		return err
	}

	m.ID, err = res.LastInsertId()
	return err
}
