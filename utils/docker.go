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
	time.Sleep(1 * time.Second)
	printLogs(exec.Command("docker", "stop", containerName))
	time.Sleep(1 * time.Second)
	// remove container
	printLogs(exec.Command("docker", "rm", containerName))
	time.Sleep(1 * time.Second)
	// get enviroment variables\
	envValues := os.Environ()
	var env string = ""
	for _, value := range envValues {
		env += "-e " + value + " "
	}
	// create container depand on database
	fmt.Println("docker", "run", "-d", "--name", containerName, "--network", "probat", env, "n4z3m/probat-fstack:latest")
	printLogs(exec.Command("docker", "run", "-d", "--name", containerName, "--network", "probat", env, "n4z3m/probat-fstack:latest"))
	return true
}
