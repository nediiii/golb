package services

import (
	"context"
	"log"
	"time"

	"github.com/mailgun/mailgun-go/v4"
	"gopkg.in/go-playground/validator.v9"
)

// Your available domain names can be found here:
// (https://app.mailgun.com/app/domains)
var yourDomain string = "mail.macrohard.cloud" // e.g. mg.yourcompany.com

const sysDomain = "blog.macrohard.cloud"

// You can find the Private API Key in your Account Menu, under "Settings":
// (https://app.mailgun.com/app/account/security)
var privateAPIKey string = "25f028417d88c8a2ab0299b28ed55f49-9dda225e-bb22f5c0"

// send email
func send(recipient, sender, subject, body string) {
	// Create an instance of the Mailgun Client
	mg := mailgun.NewMailgun(yourDomain, privateAPIKey)

	// TODO: delete it
	log.Println("sender", sender)
	log.Println("subject", subject)
	log.Println("body", body)
	log.Println("recipient", recipient)

	// The message object allows you to add attachments and Bcc recipients
	message := mg.NewMessage(sender, subject, body, recipient)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	// Send the message with a 10 second timeout
	resp, id, err := mg.Send(ctx, message)

	if err != nil {
		log.Fatal(err)
	}
	log.Printf("ID: %s Resp: %s\n", id, resp)
}

// Reply reply comment email notification
func Reply(sender, recipient, postSlug string) {
	// validate email format
	v := validator.New()
	errs := v.Var(recipient, "required,email")
	if errs != nil {
		log.Println(errs) // output: Key: "" Error:Field validation for "" failed on the "email" tag
		return
	}
	subject := "Comment Notice"
	body := sender + " reply your comment. plz visit https://" + sysDomain + "/post/" + postSlug + "/ to check it."
	send(recipient, "notice@"+sysDomain, subject, body)
}
