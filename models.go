package main

import "fmt"

/*
	Création de tout les models de notre BDD
*/

// ! Struct d'un movie
type Movie struct {
	ID          int64  `db:"id"`
	Title       string `db:"title"`
	ReleaseDate string `db:"release_date"`
	Duration    int    `db:"duration"`
	TrailerURL  string `db:"trailer_url"`
}

func (m Movie) String() string {
	return fmt.Sprintf("id=%v title=%v release date=%v duration=%v trailer url=%v\n", m.ID, m.Title, m.ReleaseDate, m.Duration, m.TrailerURL)
}
