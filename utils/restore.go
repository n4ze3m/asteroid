package utils

import (
	"fmt"
	"os/exec"
)

func Restore() {
	
	FILENAME := "restore.sql"
	cmd := fmt.Sprintf("mysql -h %s -u %s -p%s %s < %s", HOST, USERNAME, PASSWORD, DATABASE, FILENAME)
	err := exec.Command("sh", "-c", cmd).Run()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Restore complete! locally")
}
