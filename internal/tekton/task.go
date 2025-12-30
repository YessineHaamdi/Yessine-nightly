package tekton

import (
	"context"
	"fmt"

	tektonv1 "github.com/tektoncd/pipeline/pkg/apis/pipeline/v1"
	tektonclient "github.com/tektoncd/pipeline/pkg/client/clientset/versioned"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// CreateTodoTask creates a Task that simulates building the application
func CreateTodoTask(client tektonclient.Interface, namespace string) (*tektonv1.Task, error) {
	taskName := "build-todo-app"

	// Define the Task structure
	task := &tektonv1.Task{
		ObjectMeta: metav1.ObjectMeta{
			Name:      taskName,
			Namespace: namespace,
		},
		Spec: tektonv1.TaskSpec{
			Description: "A task to build the todo application",
			Steps: []tektonv1.Step{
				{
					Name:    "compile",
					Image:   "golang:1.21-alpine",
					Command: []string{"/bin/sh"},
					Args:    []string{"-c", "echo 'Compiling Todo App...'; sleep 2; echo 'Build Success!'"},
				},
				{
					Name:    "test",
					Image:   "golang:1.21-alpine",
					Command: []string{"echo"},
					Args:    []string{"Running Unit Tests... Passed."},
				},
			},
		},
	}

	// Create the Task in the cluster
	// Check if it exists first to avoid conflict, or use Update logic (omitted for brevity)
	createdTask, err := client.TektonV1().Tasks(namespace).Create(context.TODO(), task, metav1.CreateOptions{})
	if err != nil {
		return nil, fmt.Errorf("failed to create task: %v", err)
	}

	return createdTask, nil
}