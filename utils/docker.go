package utils

import (
	"fmt"
	"os/exec"
)

func prinnError(err error) {
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Success")
	}
}


func DockerUpdate(version string) bool {
	containerName := "backend"
	prinnError(exec.Command("sh", "-c", "docker -v").Run())
	cmd := "docker-compose pull " + containerName
	prinnError(exec.Command("sh", "-c", cmd).Run())
	stopCmd := "docker-compose stop backend" + containerName
	prinnError(exec.Command("sh", "-c", stopCmd).Run())
	upCmd := "docker-compose up -d backend" + containerName
	prinnError(exec.Command("sh", "-c", upCmd).Run())
	return true
}