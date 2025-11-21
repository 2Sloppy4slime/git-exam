package main

import (
	"fmt"
	"log"
	"net/http"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "route: %q", r.URL.Path)
}
func ColorHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Colors</h1><style>*{background-color: #006400;}</style>")
}
func main() {
	http.HandleFunc("/", IndexHandler)
	http.HandleFunc("/color", ColorHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

//modification distante pour exo 4 partie 2 oulala je suis un employé qui peut pas bosser sur un autre fichier et qui complique la vie de tout le monde
