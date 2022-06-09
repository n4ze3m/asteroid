package utils

import (
	"fmt"
	"os/exec"
	"time"
)

func Backup() {

	FILENAME := fmt.Sprintf("%s-%s.sql", DATABASE, time.Now().Format("2006-01-02-15-04-05"))
	cmd := fmt.Sprintf("mysqldump -h %s -u %s -p%s %s > %s", HOST, USERNAME, PASSWORD, DATABASE, FILENAME)
	err := exec.Command("sh", "-c", cmd).Run()

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Backup complete! locally")
}
