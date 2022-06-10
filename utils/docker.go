package utils

import "os/exec"


func DockerUpdate(version string) {
	image := GetEnv("DOCKER_IMAGE") + ":" + version
	cmd := "docker-compose pull " + image
	exec.Command("sh", "-c", cmd).Run()
	stopCmd := "docker-compose stop backend" + image
	exec.Command("sh", "-c", stopCmd).Run()
	upCmd := "docker-compose up -d backend" + image
	exec.Command("sh", "-c", upCmd).Run()
}