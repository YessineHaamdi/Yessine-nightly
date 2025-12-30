package tekton

import (
	"fmt"
	"os"
	"path/filepath"

	tektonclient "github.com/tektoncd/pipeline/pkg/client/clientset/versioned"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

// NewClient builds a Tekton clientset using KUBECONFIG, default kubeconfig location, or in-cluster config.
func NewClient() (tektonclient.Interface, error) {
	// 1) Try KUBECONFIG env var
	if kube := os.Getenv("KUBECONFIG"); kube != "" {
		cfg, err := clientcmd.BuildConfigFromFlags("", kube)
		if err == nil {
			return tektonclient.NewForConfig(cfg)
		}
		// if error, fallthrough to other methods
	}

	// 2) Try user's default kubeconfig (e.g., $HOME/.kube/config)
	home, err := os.UserHomeDir()
	if err == nil {
		kube := filepath.Join(home, ".kube", "config")
		if _, err := os.Stat(kube); err == nil {
			cfg, err := clientcmd.BuildConfigFromFlags("", kube)
			if err == nil {
				return tektonclient.NewForConfig(cfg)
			}
		}
	}

	// 3) Try in-cluster config (useful when running inside Kubernetes)
	cfg, err := rest.InClusterConfig()
	if err == nil {
		return tektonclient.NewForConfig(cfg)
	}

	return nil, fmt.Errorf("could not create tekton client: no kubeconfig found and in-cluster config failed: %w", err)
}
