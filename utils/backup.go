package utils

import (
	"bytes"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"os/exec"
	"time"
)

func Backup() {
	HOST := GetEnv("HOST")
	USERNAME := GetEnv("USERNAME")
	PASSWORD := GetEnv("PASSWORD")
	DATABASE := GetEnv("DATABASE")
	FILENAME := fmt.Sprintf("%s-%s.sql", DATABASE, time.Now().Format("2006-01-02-15-04-05"))
	cmd := fmt.Sprintf("mysqldump -h %s -P 3156 --no-tablespaces --protocol=tcp -u %s -p%s %s > %s", HOST, USERNAME, PASSWORD, DATABASE, FILENAME)
	err := exec.Command("sh", "-c", cmd).Run()

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Backup complete! locally")
	// upload to server
	url := "https://cdn.appdocker.xyz/api/upload/backup"
	// multipart form
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	// file
	part, err := writer.CreateFormFile("file", FILENAME)
	if err != nil {
		fmt.Println(err)
	}
	// read file
	file, err := os.Open(FILENAME)
	if err != nil {
		fmt.Println(err)
	}
	io.Copy(part, file)
	// close writer
	err = writer.Close()
	if err != nil {
		fmt.Println(err)
	}
	// request
	request, err := http.NewRequest("POST", url, body)
	if err != nil {
		fmt.Println(err)
	}
	// set header
	request.Header.Set("Content-Type", writer.FormDataContentType())
	// send request
	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		fmt.Println(err)
	}
	// close response
	// pirnt the response
	defer response.Body.Close()
	fmt.Println("Backup complete! on server")
}
