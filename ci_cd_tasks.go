package main

import (
	"log"
	"net/http"
	"os"
	"os/exec"
	"io"
	"bufio"
)

func main() {
	setupRoutes()

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalf("Failed to start HTTP server: %v", err)
	}
}

func setupRoutes() {
	http.HandleFunc("/build", handleBuild)
	http.HandleFunc("/test", handleTest)
	http.HandleFunc("/deploy", handleDeploy)
	http.HandleFunc("/rollback", handleRollback)
	http.HandleFunc("/stream-logs", handleStreamLogs) // New route for log streaming
}

func handleBuild(w http.ResponseWriter, r *http.Request) {
	if runCommand(w, "go", "build", ".") {
		respondWithMessage(w, "Build successful")
	} else {
		respondWithError(w, "Build failed")
	}
}

func handleTest(w http.ResponseWriter, r *http.Request) {
	if runCommand(w, "go", "test", "./...") {
		respondWithMessage(w, "Tests passed")
	} else {
		respondWithError(w, "Tests failed")
	}
}

func handleDeploy(w http.ResponseWriter, r *http.Request) {
	if err := os.Setenv("DEPLOYED_VERSION", "v1.0"); err != nil {
		log.Printf("Could not set DEPLOYED_VERSION: %v", err)
		respondWithError(w, "Deployment failed due to environment setup issue")
		return
	}
	respondWithMessage(w, "Deployed successfully")
}

func handleRollback(w http.ResponseWriter, r *http.Request) {
	if err := os.Setenv("DEPLOYED_VERSION", "v0.9"); err != nil {
		log.Printf("Could not set DEPLOYED_VERSION: %v", err)
		respondWithError(w, "Rollback failed due to environment setup issue")
		return
	}
	respondWithMessage(w, "Rollback successful")
}

func runCommand(w http.ResponseWriter, name string, args ...string) bool {
	cmd := exec.Command(name, args...)
	cmd.Env = append(os.Environ())
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		log.Printf("Error obtaining stdout: %v", err)
		return false
	}
	if err := cmd.Start(); err != nil {
		log.Printf("Failed to start command: %v", err)
		return false
	}
	// Stream the output
	scanner := bufio.NewScanner(stdout)
	for scanner.Scan() {
		io.WriteString(w, scanner.Text()+"\n")
	}
	if err := cmd.Wait(); err != nil {
		log.Printf("Command finished with error: %v", err)
		return false
	}
	return true
}

func handleStreamLogs(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "text/plain")
	io.WriteString(w, "Log streaming setup endpoint\n")
	// Any streaming logic or SSE setup can be added here
}

func respondWithMessage(w http.ResponseWriter, message string) {
	w.WriteHeader(http.StatusOK)
	if _, err := w.Write([]byte(message)); err != nil {
		log.Printf("Error writing success response: %v", err)
	}
}

func respondWithError(w http.ResponseWriter, message string) {
	http.Error(w, message, http.StatusInternalServerError)
}