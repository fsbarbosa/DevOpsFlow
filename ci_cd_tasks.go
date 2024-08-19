package main

import (
	"log"
	"net/http"
	"os"
	"os/exec"
)

func main() {
	// Setup HTTP endpoints
	setupRoutes()

	// Start the HTTP server
	log.Fatal(http.ListenAndServe(":8080", nil))
}

// setupRoutes defines the HTTP endpoints and their corresponding handler functions.
func setupRoutes() {
	http.HandleFunc("/build", handleBuild)
	http.HandleFunc("/test", handleTest)
	http.HandleFunc("/deploy", handleDeploy)
	http.HandleFunc("/rollback", handleRollback)
}

// handleBuild executes the project build process.
func handleBuild(w http.ResponseWriter, r *http.Request) {
	if runCommand("go", "build", ".") {
		respondWithMessage(w, "Build successful")
	} else {
		respondWithError(w, "Build failed")
	}
}

// handleTest executes the project test process.
func handleTest(w http.ResponseWriter, r *http.Request) {
	if runCommand("go", "test", "./...") {
		respondWithMessage(w, "Tests passed")
	} else {
		respondWithError(w, "Tests failed")
	}
}

// handleDeploy performs the project deployment process.
func handleDeploy(w http.ResponseWriter, r *http.Request) {
	os.Setenv("DEPLOYED_VERSION", "v1.0")
	respondWithMessage(w, "Deployed successfully")
}

// handleRollback performs the project rollback process.
func handleRollback(w http.ResponseWriter, r *http.Request) {
	os.Setenv("DEPLOYED_VERSION", "v0.9")
	respondWithMessage(w, "Rollback successful")
}

// runCommand executes a system command and returns a boolean status.
func runCommand(name string, args ...string) bool {
	cmd := exec.Command(name, args...)
	cmd.Env = append(os.Environ())
	if err := cmd.Run(); err != nil {
		return false
	}
	return true
}

// respondWithMessage sends a success response back to the client.
func respondWithMessage(w http.ResponseWriter, message string) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(message))
}

// respondWithError sends an error response back to the client.
func respondWithError(w http.ResponseWriter, message string) {
	http.Error(w, message, http.StatusInternalServerError)
}