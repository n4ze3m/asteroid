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
	printLogs(exec.Command("docker-compose", "pull", "bcakend"))
	printLogs(exec.Command("docker-compose", "stop", "bcakend"))
	printLogs(exec.Command("docker-compose", "start", "bcakend"))
	return true
}
