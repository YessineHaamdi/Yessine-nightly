package tekton

import (
	"context"
	"fmt"

	tektonv1 "github.com/tektoncd/pipeline/pkg/apis/pipeline/v1"
	tektonclient "github.com/tektoncd/pipeline/pkg/client/clientset/versioned"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// CreateTodoPipeline creates a Pipeline that orchestrates the build task
func CreateTodoPipeline(client tektonclient.Interface, namespace string) (*tektonv1.Pipeline, error) {
	pipelineName := "todo-pipeline"

	pipeline := &tektonv1.Pipeline{
		ObjectMeta: metav1.ObjectMeta{
			Name:      pipelineName,
			Namespace: namespace,
		},
		Spec: tektonv1.PipelineSpec{
			Tasks: []tektonv1.PipelineTask{
				{
					Name: "build-step",
					TaskRef: &tektonv1.TaskRef{
						Name: "build-todo-app", // Matches the Task name created earlier
						Kind: "Task",
					},
				},
			},
		},
	}

	createdPipeline, err := client.TektonV1().Pipelines(namespace).Create(context.TODO(), pipeline, metav1.CreateOptions{})
	if err != nil {
		return nil, fmt.Errorf("failed to create pipeline: %v", err)
	}

	return createdPipeline, nil
}

// DeleteTodoPipeline deletes the example pipeline created by this service.
func DeleteTodoPipeline(client tektonclient.Interface, namespace string) error {
	pipelineName := "todo-pipeline"
	// Use Background or Foreground deletion options as appropriate; here we just delete.
	if err := client.TektonV1().Pipelines(namespace).Delete(context.TODO(), pipelineName, metav1.DeleteOptions{}); err != nil {
		return fmt.Errorf("failed to delete pipeline %s: %v", pipelineName, err)
	}
	return nil
}