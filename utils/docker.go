package utils

import (
	"fmt"
	"os/exec"
	"time"
)

func printLogs(cmd *exec.Cmd) {
	fmt.Println(cmd.Args)
	out, err := cmd.Output()
	if err != nil {
		fmt.Println(err)
	}
	// get the output
	fmt.Println(string(out))
	time.Sleep(time.Second * 5)
}

func DockerUpdate(version string) bool {
	printLogs(exec.Command("docker", "pull", "n4z3m/probat-fstack:latest"))
	printLogs(exec.Command("docker-compose", "stop", "backend"))
	printLogs(exec.Command("docker-compose", "rm", "-f", "backend"))
	printLogs(exec.Command("docker-compose", "up", "-d", "backend"))
	return true
}
