package main

import (
	"log"
	"os"
	"path/filepath"
	"strings"
)

const (
	kubeDirectory = "/.kube/"
)

func main() {
	// Final $KUBECONFIG
	var allKubeconfigPaths strings.Builder

	homeDirectoryPath := os.Getenv("HOME")
	basePath := filepath.Join(homeDirectoryPath, ".kube")

	additionalKubeconfigPath := filepath.Join(basePath, "config-files")
	defaultKubeconfigPath := filepath.Join(basePath, "config")

	data, err := os.ReadFile(defaultKubeconfigPath)
	if err != nil {
		log.Fatal(err)
	}

	// If default kubeconfig is present, add to path to $KUBECONFIG
	if len(data) != 0 {

		allKubeconfigPaths.WriteString(defaultKubeconfigPath)
		allKubeconfigPaths.WriteString(":")
	}

	// Add paths to additional kubeconfigs if any are present
	directories, err := os.ReadDir(additionalKubeconfigPath)

	totalDirectories := len(directories)
	for index, directory := range directories {

		path := filepath.Join(basePath, directory.Name())
		allKubeconfigPaths.WriteString(path)

		if index != (totalDirectories - 1) {
			allKubeconfigPaths.WriteString(":")
		}
	}

	os.Setenv("KUBECONFIG", allKubeconfigPaths.String())
	log.Printf("KUBECONFIG has been set to: %s", os.Getenv("KUBECONFIG"))
}
