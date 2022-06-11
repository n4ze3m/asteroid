package main

import (
	"fmt"
	"log"
	"time"

	"github.com/joho/godotenv"
	"github.com/n4ze3m/asteroid/utils"
)

func init() {

	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	for {
		astro, found := utils.Select()
		if found {

			fmt.Println("Backuping... current database")
			// utils.Backup()
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
