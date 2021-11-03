package mail

import (
	"log"
	"os"
	"strconv"

	"git.xenonstack.com/util/test-portal/config"
	"go.uber.org/zap"
	gomail "gopkg.in/gomail.v2"
)

// function for sending mail
func SendMailV2(to string, sub string, template string) {
	if config.DisableMail == "true" {
		log.Println("mail is disabled")
		return
	}

	// creating new message with default settings
	m := gomail.NewMessage()

	// setting mail headers from, to and subject
	m.SetHeader("From", os.Getenv("HIRING_PORTAL_MAIL_FROM"))
	m.SetHeader("To", to)
	m.SetHeader("Subject", sub)

	//path is from where main.go is running
	// embedding static images
	m.Embed(config.MailPath + "/mail/images/xenonstack.png")
	m.Embed(config.MailPath + "/mail/images/testcompleted.png")

	// set body of mail
	m.SetBody("text/html", template)

	// port of smtp mail
	port, _ := strconv.Atoi(os.Getenv("HIRING_PORTAL_MAIL_SMTP_PORT"))

	//use port 465 for TLS, other than 465 it will send without TLS.
	// connect to smtp server using mail admin username and password
	d := gomail.NewPlainDialer(os.Getenv("HIRING_PORTAL_MAIL_SMTP_HOST"), port, os.Getenv("HIRING_PORTAL_MAIL_USERID"), os.Getenv("HIRING_PORTAL_MAIL_PASS"))

	if port == 465 {
		d.SSL = true
	}

	// send above mail message
	err := d.DialAndSend(m)
	zap.S().Info("error in sending mail......", err)

}
