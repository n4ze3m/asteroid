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
	fmt.Println(string(out))
	time.Sleep(time.Second * 2)
}

func DockerUpdate(version string) bool {
	printLogs(exec.Command("docker-compose", "pull", "backend"))
	printLogs(exec.Command("docker-compose", "stop", "backend"))
	printLogs(exec.Command("docker-compose", "start", "backend"))
	return true
}
