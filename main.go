package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func loadEnvVars() {
	os.Setenv("PORT", "8080")
}

func main() {
	loadEnvVars()

	router := mux.NewRouter()

	router.HandleFunc("/deploy", DeployHandler).Methods("POST")
	router.HandleFunc("/test", TestHandler).Methods("POST")

	port := os.Getenv("PORT")
	fmt.Printf("Server started on port %s\n", port)

	log.Fatal(http.ListenAndServe(":"+port, router))
}

func DeployHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Deploying application...")
}

func TestHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Running tests...")
}