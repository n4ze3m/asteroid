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

			fmt.Println("Backuping... current database")
			utils.Backup()
			time.Sleep(time.Second * 5)
			fmt.Println("Updating...")
			ok := utils.DockerUpdate(astro.Version)
			if ok {
				fmt.Println("Updated to", astro.Version)
				utils.Update(astro.Id)
			} else {
				fmt.Println("Failed to update to", astro.Version)
			}
		}
		time.Sleep(2 * time.Minute)
	}
}
