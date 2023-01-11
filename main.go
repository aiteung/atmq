package main

import (
	"encoding/base64"
	"fmt"
	"log"
	"os"

	"github.com/aiteung/atmail"
	"google.golang.org/api/gmail/v1"
)

func filereadmime(fname string) (fileData string) {
	fileBytes, err := os.ReadFile(fname)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
	fileData = base64.StdEncoding.EncodeToString(fileBytes)
	return fileData
}

func main() {
	filename := "README.md"
	to := "awangga@gmail.com"
	from := "Penerbit Buku Pedia<penerbit@bukupedia.co.id>"
	body := "ini bodi isinya ya <b>luar biasa</b>"
	subject := "subjek email"
	fileMIMEType := "text/plain; charset=utf-8"
	fileData := filereadmime(filename)
	fmt.Println(fileMIMEType)
	srv := atmail.GetGmailService("client_secret.json", "token.json", gmail.GmailSendScope)
	message := atmail.GenerateGmailMessageWithAttachment(from, to, subject, body, fileMIMEType, filename, fileData)
	// Send the message
	resp, err := srv.Users.Messages.Send("me", &message).Do()
	if err != nil {
		log.Printf("Error: %v", err)
	} else {
		fmt.Println(resp.LabelIds)
	}
}
