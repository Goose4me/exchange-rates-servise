package mail

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/mailersend/mailersend-go"
)

type MailersendWrapper struct {
	ms *mailersend.Mailersend
}

var MailService MailersendWrapper

func sendMail(receipient mailersend.Recipient, subject *string, text *string) {
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	from := mailersend.From{
		Name:  "Example Company",
		Email: os.Getenv("MAILERSEND_INFO_MAIL"),
	}

	recipients := []mailersend.Recipient{
		receipient,
	}

	message := MailService.ms.Email.NewMessage()

	message.SetFrom(from)
	message.SetRecipients(recipients)
	message.SetSubject(*subject)
	message.SetText(*text)

	_, err := MailService.ms.Email.Send(ctx, message)

	if err != nil {
		log.Print("Sending for " + receipient.Email + " failed. Error: " + err.Error())
	}
}

func SendMails(recipients *[]string, subject *string, text *string) {
	log.Print(*recipients)
	for _, receipient := range *recipients {
		sendMail(mailersend.Recipient{
			Name:  receipient,
			Email: receipient,
		}, subject, text)
	}
}

func SetupMailService() {
	if MailService.ms != nil {
		return
	}

	// Create an instance of the mailersend client
	MailService = MailersendWrapper{
		ms: mailersend.NewMailersend(os.Getenv("MAILERSEND_API_KEY")),
	}
}
