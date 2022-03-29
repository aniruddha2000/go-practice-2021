package main

import (
	"crypto/tls"
	"fmt"
	"log"

	"net/smtp"

	"github.com/aniruddha2000/goSendEmail/config"
	"gopkg.in/gomail.v2"
)

type EmailData struct {
	ToEmail   []string
	EmailBody []byte
}

type EmailDataGoMail struct {
	ToEmail   string
	EmailBody string
}

func main() {
	config, err := config.LoadConfig("./config")
	if err != nil {
		log.Fatalf("Error loading config: %v", err)
	}
	log.Println(config)

	sendMailSMTP(config)
	sendMailGoMail(config)
}

func sendMailSMTP(config *config.EmailSMTP) {

	toEmail := []EmailData{
		{
			[]string{"riki.bask@gmail.com"},
			[]byte("This is a test messesge from Go SMTP"),
		},
	}

	auth := smtp.PlainAuth("", config.Email, config.Password, config.EmailHost)

	err := smtp.SendMail(config.EmailHost+":"+config.EmailPortString, auth, config.Email,
		toEmail[0].ToEmail, toEmail[0].EmailBody)
	if err != nil {
		log.Println(config.EmailHost, config.EmailPort)
		log.Fatalf("Couldn't send email : %v", err)
	}
	fmt.Println("Email Sent successfully!")
}

func sendMailGoMail(config *config.EmailSMTP) {
	toEmailGomail := []EmailDataGoMail{
		{
			"riki.bask@gmail.com",
			"This is a test messesge from Go Mail package",
		},
	}

	mail := gomail.NewMessage()
	mail.SetHeaders(map[string][]string{
		"From":    {mail.FormatAddress(config.Email, "ricky")},
		"To":      {toEmailGomail[0].ToEmail},
		"Subject": {"This is mail from Go Mail package"},
	})
	mail.SetBody("text/plain", toEmailGomail[0].EmailBody)

	dial := gomail.NewDialer(config.EmailHost, config.EmailPort, config.Email, config.Password)
	dial.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	if err := dial.DialAndSend(mail); err != nil {
		log.Fatalf("Couldn't send email : %v", err)
	}
	fmt.Println("Email Sent successfully!")
}
