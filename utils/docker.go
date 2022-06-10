package utils

import (
	"bytes"
	"fmt"
	"os/exec"
)

func printLogs(cmd *exec.Cmd) {
	var out bytes.Buffer
	cmd.Stdout = &out

	err := cmd.Run()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(out.String())
}

func DockerUpdate(version string) bool {
	containerName := "backend"
	cmd := "docker-compose pull " + containerName
	printLogs(exec.Command("sh", "-c", cmd))

	stopCmd := "docker-compose stop backend" + containerName
	printLogs(exec.Command("sh", "-c", stopCmd))
	upCmd := "docker-compose up -d backend" + containerName
	printLogs(exec.Command("sh", "-c", upCmd))

	return true
}
