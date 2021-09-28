package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/smtp"
	"os"
	"path"
	"runtime"
	"time"

	"github.com/joho/godotenv"
	"github.com/nats-io/nats.go"
)

type User struct {
	Name    string `json:"name"`
	Email   string `json:"email"`
	Address string `json:"address"`
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	Subscribe()
}

func Subscribe() {
	// urls := nats.DefaultURL
	urls := "nats-server://nats-server:4222"
	options := []nats.Option{nats.Name("sales service")}
	options = setupConnections(options)

	nc, err := nats.Connect(urls, options...)
	if err != nil {
		log.Fatal(err)
	}

	smtpHost := "smtp.gmail.com"
	smtpPort := "587"

	from := os.Getenv("EMAIL")
	fromPassword := os.Getenv("EMAIL_PASSWORD")

	auth := smtp.PlainAuth("", from, fromPassword, smtpHost)

	subjectUsers, i := "new-user-registered", 0
	nc.Subscribe(subjectUsers, func(msg *nats.Msg) {
		log.Printf("[#%d] Received on [%s]: '%s'", i, msg.Subject, string(msg.Data))
		var userData User
		err = json.Unmarshal([]byte(msg.Data), &userData)
		if err != nil {
			log.Println(err)
		}
		log.Println(userData)

		var body bytes.Buffer
		mimeHeaders := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
		body.Write([]byte(fmt.Sprintf("Welcome to Warung Serba Ada \n%s\n\n", mimeHeaders)))
		filepath := path.Join("template.txt")
		t, err := template.ParseFiles(filepath)
		if err != nil {
			log.Println(err)
		}

		t.Execute(&body, struct {
			Name    string
			Email   string
			Address string
		}{
			Name:    userData.Name,
			Email:   userData.Email,
			Address: userData.Address,
		})

		err = smtp.SendMail(smtpHost+":"+smtpPort, auth, from, []string{userData.Email}, body.Bytes())
		if err != nil {
			fmt.Println(err)
			return
		}
		log.Println("Email Sent!")

	})
	nc.Flush()

	if err := nc.LastError(); err != nil {
		log.Fatal(err)
	}

	log.Printf("Listening on [%s]", subjectUsers)
	log.SetFlags(log.LstdFlags)

	runtime.Goexit()
}

func setupConnections(options []nats.Option) []nats.Option {
	totalWait := 10 * time.Minute
	reconnectDelay := time.Second

	options = append(options, nats.ReconnectWait(reconnectDelay))
	options = append(options, nats.MaxReconnects(int(totalWait/reconnectDelay)))
	options = append(options, nats.DisconnectErrHandler(func(nc *nats.Conn, err error) {
		log.Printf("Disconnected due to:%s, will attempt reconnects for %.0fm", err, totalWait.Minutes())
	}))
	options = append(options, nats.ReconnectHandler(func(nc *nats.Conn) {
		log.Printf("Reconnected [%s]", nc.ConnectedUrl())
	}))
	options = append(options, nats.ClosedHandler(func(nc *nats.Conn) {
		log.Fatalf("Exiting: %v", nc.LastError())
	}))
	return options
}
