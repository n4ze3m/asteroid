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
			ok := utils.DockerUpdate(astro.Version)
			if ok {
				fmt.Println("Updated to", astro.Version)
			} else {
				fmt.Println("Failed to update to", astro.Version)
			}
		}
		time.Sleep(2 * time.Minute)
	}
}
