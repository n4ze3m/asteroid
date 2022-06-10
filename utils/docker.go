package utils

import (
	"fmt"
	"os/exec"
	"time"
)

func printLogs(cmd *exec.Cmd) {
	out, err := cmd.Output()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("response", string(out))
}

func DockerUpdate(version string) bool {
	containerName := "probat_backend"
	// cmd := "docker-compose pull " + containerName
	printLogs(exec.Command("docker", "pull", "n4z3m/probat-fstack:latest"))
	time.Sleep(5 * time.Second)
	printLogs(exec.Command("docker", "stop", containerName))
	time.Sleep(5 * time.Second)
	printLogs(exec.Command("docker", "start", containerName))
	return true
}
