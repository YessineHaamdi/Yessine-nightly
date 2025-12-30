package tekton

import (
	"context"
	"fmt"

	tektonv1 "github.com/tektoncd/pipeline/pkg/apis/pipeline/v1"
	tektonclient "github.com/tektoncd/pipeline/pkg/client/clientset/versioned"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// TriggerPipelineRun starts an execution of the todo-pipeline
func TriggerPipelineRun(client tektonclient.Interface, namespace string) (*tektonv1.PipelineRun, error) {
	pipelineRun := &tektonv1.PipelineRun{
		ObjectMeta: metav1.ObjectMeta{
			GenerateName: "todo-run-", // K8s will append a random hash (e.g., todo-run-xyz)
			Namespace:    namespace,
		},
		Spec: tektonv1.PipelineRunSpec{
			PipelineRef: &tektonv1.PipelineRef{
				Name: "todo-pipeline",
			},
			// If you add Workspaces later for PVCs, they go here
		},
	}

	createdRun, err := client.TektonV1().PipelineRuns(namespace).Create(context.TODO(), pipelineRun, metav1.CreateOptions{})
	if err != nil {
		return nil, fmt.Errorf("failed to create pipelinerun: %v", err)
	}

	return createdRun, nil
}
