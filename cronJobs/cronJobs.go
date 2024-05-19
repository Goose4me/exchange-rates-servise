package cronjobs

import (
	"log"

	"github.com/robfig/cron/v3"
	"gses2.app/api"
	"gses2.app/mail"
	"gses2.app/rate"
)

type CronWrapper struct {
	c *cron.Cron
}

var CronJobsService CronWrapper

func cronSendMails() {
	log.Printf("Start sending mails cron job")

	var emails []string

	api.GetSubscribersEmails(&emails)

	currencyRate, err := rate.GetCurrencyRateFor("usd", "uah")

	if err != nil {
		log.Print("Error geting currency rate from usd to uah: " + err.Error())

		return
	}

	subject := "Exchange Rates"

	mail.SendMails(&emails, &subject, &currencyRate)

	log.Printf("Send emails to: %v", emails)
}

func SetupCronJobs() {
	if CronJobsService.c != nil {
		return
	}

	CronJobsService = CronWrapper{
		c: cron.New(),
	}

	CronJobsService.c.AddFunc("@midnight", cronSendMails)

	CronJobsService.c.Start()
}
