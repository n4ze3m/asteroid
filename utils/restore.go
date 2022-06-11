package utils

import (
	"fmt"
	"os"
	"os/exec"
)

func updateFile(sql string) {
	// create file not exist
	FILENAME := "restore.sql"
	file, err := os.Create(FILENAME)
	if err != nil {
		fmt.Println(err)
	}
	// write to file
	_, err = file.WriteString(sql)
	if err != nil {
		fmt.Println(err)
	}
	// close file
	err = file.Close()
	if err != nil {
		fmt.Println(err)
	}
}

func Restore(sql string) {
	HOST := GetEnv("HOST")
	USERNAME := GetEnv("USERNAME")
	PASSWORD := GetEnv("PASSWORD")
	DATABASE := GetEnv("DATABASE")
	FILENAME := "restore.sql"
	updateFile(sql)
	cmd := fmt.Sprintf("mysql -h %s -P 3156 --protocol=tcp -u %s -p%s %s < %s", HOST, USERNAME, PASSWORD, DATABASE, FILENAME)
	err := exec.Command("sh", "-c", cmd).Run()
	if err != nil {
		fmt.Println(err)
	}
	// delete file
	err = os.Remove(FILENAME)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Restore complete! locally")
}
