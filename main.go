package main

import (
	"fmt"
	"log"
	"net/http"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "route: %q", r.URL.Path)
}

func main() {
	http.HandleFunc("/", IndexHandler)
	// to add : color functionality
	log.Fatal(http.ListenAndServe(":8080", nil))
}

//modification distante pour exo 4 partie 2 oulala je suis un employé qui peut pas bosser sur un autre fichier et qui complique la vie de tout le monde
