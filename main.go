package main

import (
    "fmt"
    "log"
    "net/http"
    "os"
    "time"

    "github.com/gorilla/mux"
)

func loadEnvVars() {
    os.Setenv("PORT", "8080")
}

func main() {
    loadEnvVars()

    router := mux.NewRouter()

    loggedRouter := logRequest(router)

    router.HandleFunc("/deploy", DeployHandler).Methods("POST")
    router.HandleFunc("/test", TestHandler).Methods("POST")

    port := os.Getenv("PORT")
    fmt.Printf("Server started on port %s\n", port)

    log.Fatal(http.ListenAndServe(":"+port, loggedRouter))
}

func DeployHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "Deploying application...")
}

func TestHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "Running tests...")
}

func logRequest(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        startTime := time.Now()
        log.Printf("Started %s %s", r.Method, r.URL.Path)

        next.ServeHTTP(w, r) 

        log.Printf("Completed %s in %v", r.URL.Path, time.Since(startTime))
    })
}