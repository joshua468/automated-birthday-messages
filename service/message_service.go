package services

import (
	"fmt"
	"net/smtp"

	"github.com/joshua468/automated-birthday-messenger/db"
	"github.com/joshua468/automated-birthday-messenger/models"
)

func SendBirthdayMessage(employee models.Employee) error {
	cfg := db.InitDB()
	to := []string{employee.Email}
	subject := "Happy Birthday!"
	body := fmt.Sprintf("Happy Birthday, %s! Wishing you a wonderful day!", employee.Name)

	msg := []byte("From: " + cfg.EmailUser + "\n" +
		"To: " + employee.Email + "\n" +
		"Subject: " + subject + "\n\n" +
		body)

	err := smtp.SendMail(fmt.Sprintf("%s:%s", cfg.SMTP_SERVER, cfg.SMTP_PORT),
		smtp.PlainAuth("", cfg.EmailUser, cfg.EmailPass, cfg.SMTP_SERVER), cfg.EmailUser, to, msg)
	if err != nil {
		return err
	}
	return nil
}
