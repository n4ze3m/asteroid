package main

import (
	"fmt"
	"time"

	"github.com/n4ze3m/asteroid/utils"
)

func main() {
	for {
		astro, found := utils.Select()
		if found {
			utils.DockerUpdate(astro.Version)
			fmt.Println("Updated to", astro.Version)
		}
		time.Sleep(2 * time.Minute)
	}
}
