package main

import (
	"article/handlers"
	"fmt"
	"net/http"
)

const PORT = "8000"

func main() {
	http.HandleFunc("/", Gethome)
	http.HandleFunc("/users", handlers.HandlerUsers)
	//	http.HandleFunc("/users", getusers).Methods("GET")
	//	http.HandleFunc("/users", getusers).Methods("GET")

	http.ListenAndServe(":"+PORT, nil)
}
func Gethome(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Nodir Developer")
}
