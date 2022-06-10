package utils

import (
	"fmt"
	"os/exec"
)

func printLogs(cmd *exec.Cmd) {
	out, err := cmd.Output()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("response", string(out))
}

func DockerUpdate(version string) bool {
	containerName := "backend"
	// cmd := "docker-compose pull " + containerName
	printLogs(exec.Command("docker-compose", "-v"))
	printLogs(exec.Command("~/.docker/cli-plugins/docker-compose", "-v"))

	printLogs(exec.Command("docker-compose", "pull", containerName))

	stopCmd := "docker-compose stop backend" + containerName
	printLogs(exec.Command("sh", "-c", stopCmd))
	upCmd := "docker-compose up -d backend" + containerName
	printLogs(exec.Command("sh", "-c", upCmd))

	return true
}
