package main

import (
	"fmt"
	"log"

	// Import your internal packages (match module name in go.mod)
	"tekton-backend/internal/tekton"
)

func main() {
	// 1. Initialize Client
	client, err := tekton.NewClient()
	if err != nil {
		log.Fatalf("failed to create tekton client: %v", err)
	}

	namespace := "default" // or a specific namespace

	// 2. Create the Task Definition
	fmt.Println("Creating Task...")
	_, err = tekton.CreateTodoTask(client, namespace)
	if err != nil {
		// Note: In real app, handle "already exists" errors gracefully
		log.Printf("Warning creating task: %v", err)
	}

	// 3. Create the Pipeline Definition
	fmt.Println("Creating Pipeline...")
	_, err = tekton.CreateTodoPipeline(client, namespace)
	if err != nil {
		log.Printf("Warning creating pipeline: %v", err)
	}

	// 4. Trigger the PipelineRun (kept for backwards compatibility)
	fmt.Println("Triggering PipelineRun...")
	pr, err := tekton.TriggerPipelineRun(client, namespace)
	if err != nil {
		log.Printf("Error triggering run: %v", err)
	} else {
		fmt.Printf("Successfully started PipelineRun: %s\n", pr.Name)
		fmt.Println("Check status with: tkn pipelinerun logs -f " + pr.Name)
	}

	// Start HTTP server with handlers to create/delete pipeline
	fmt.Println("Starting HTTP server on :8080")
	http.HandleFunc("/pipeline/create", api.CreatePipelineHandler)
	http.HandleFunc("/pipeline/delete", api.DeletePipelineHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
