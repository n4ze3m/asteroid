package utils

import (
	"fmt"
	"os"
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
	printLogs(exec.Command("docker", "rm -f", containerName))
	time.Sleep(5 * time.Second)
	// recreate container with --name <containerName>, env current enviroment variables
	cmd := "docker run -d --name " + containerName + " -e " + os.Getenv("ENV_VARIABLES") + " n4z3m/probat-fstack:lastest" 
	printLogs(exec.Command("sh", "-c", cmd))
	return true
}
