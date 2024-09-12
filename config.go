package main

import (
	"fmt"
	"os"
)

type Environment struct {
	Name           string
	DeployPath     string
	BuildCommand   string
	TestCommand    string
	ReleaseCommand string
}

type PipelineConfig struct {
	Environments []Environment
}

func main() {
	developmentDeployPath := os.Getenv("DEV_DEPLOY_PATH")
	stagingDeployPath := os.Getenv("STAGING_DEPLOY_PATH")
	productionDeployPath := os.Getenv("PRODUCTION_DEPLOY_PATH")

	pipelineConfig := PipelineConfig{
		Environments: []Environment{
			{
				Name:           "development",
				DeployPath:     developmentDeployPath,
				BuildCommand:   "go build -o devApp",
				TestCommand:    "go test ./...",
				ReleaseCommand: "go build -o devAppRelease",
			},
			{
				Name:           "staging",
				DeployPath:     stagingDeployPath,
				BuildCommand:   "go build -o stagingApp",
				TestCommand:    "go test --race ./...",
				ReleaseCommand: "go build -o stagingAppRelease",
			},
			{
				Name:           "production",
				DeployPath:     productionDeployPath,
				BuildCommand:   "go build -ldflags=\"-s -w\" -o prodApp",
				TestCommand:    "go test -v ./...",
				ReleaseCommand: "go build -ldflags=\"-s -w\" -o prodAppRelease",
			},
		},
	}

	for _, env := range pipelineConfig.Environments {
		executeEnvironmentCommands(env)
	}
}

func executeEnvironmentCommands(env Environment) {
	fmt.Printf("Executing commands for environment: %s\n", env.Name)
	fmt.Printf("Build command: %s\n", env.BuildCommand)
	fmt.Printf("Test command: %s\n", env.TestCommand)
	fmt.Printf("Release command: %s\n", env.ReleaseCommand)
	fmt.Println("Commands simulated. Replace with actual execution as needed.")
}