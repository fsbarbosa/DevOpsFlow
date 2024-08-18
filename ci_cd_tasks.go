package main

import (
	"log"
	"net/http"
	"os"
	"os/exec"
)

func main() {
	http.HandleFunc("/build", build)
	http.HandleFunc("/test", test)
	http.HandleFunc("/deploy", deploy)
	http.HandleFunc("/rollback", rollback)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func build(w http.ResponseWriter, r *http.Request) {
	cmd := exec.Command("go", "build", ".")
	cmd.Env = append(os.Environ())
	if err := cmd.Run(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write([]byte("Build successful"))
}

func test(w http.ResponseWriter, r *http.Request) {
	cmd := exec.Command("go", "test", "./...")
	cmd.Env = append(os.Environ())
	if err := cmd.Run(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write([]byte("Tests passed"))
}

func deploy(w http.ResponseWriter, r *http.Request) {
	os.Setenv("DEPLOYED_VERSION", "v1.0")
	w.Write([]byte("Deployed successfully"))
}

func rollback(w http.ResponseWriter, r *http.Request) {
	os.Setenv("DEPLOYED_VERSION", "v0.9")
	w.Write([]byte("Rollback successful"))
}