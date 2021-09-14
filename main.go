package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	fmt.Println("Goflix")

	if err := run(); err != nil {
		fmt.Sprintf(os.Stderr.Name(), "%s\n", err)
		os.Exit(1)
	}
}

/*
	Déplacer le code d'exécution dans une nouvelle fonction qui renvoie juste une erreur, et délégai le code d'erreur à la fonction main() ce qui rend le code un peu plus flexible
*/

func run() error {
	srv := newServer()      // Appel la fonction qui nous initialiser notre server
	srv.store = &dbStore{}  // Associer notre store BDD à notre store server
	err := srv.store.Open() // Ouvrir notre BDD
	if err != nil {
		return err
	}
	defer srv.store.Close()

	http.HandleFunc("/", srv.serveHTTP) // Tout les fonctions HTTP va être géré par serveHTTP qui lui même est géré par le router donc gorilla/mux
	log.Printf("Serving HTTP on, port 9000")

	err = http.ListenAndServe(":9000", nil)
	if err != nil {
		return err
	}
	return nil
}
