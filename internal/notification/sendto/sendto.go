package sendto

import (
	"bytes"
	"fmt"
	"net/smtp"
	"strings"
	"text/template"

	"github.com/Npeka/go-ecommerce/global"
	"go.uber.org/zap"
)

const (
	SMTPHost     = "smtp.gmail.com"
	SMTPPort     = "587"
	SMTPEmail    = "seimeicc@gmail.com"
	SMTPPassword = "ctfp eowe ngzu vtma"
)

type EmailAddress struct {
	Address string `json:"address"`
	Name    string `json:"name"`
}

type Email struct {
	From    EmailAddress `json:"from"`
	To      []string     `json:"to"`
	Subject string       `json:"subject"`
	Body    string       `json:"body"`
}

func BuildMessage(mail Email) string {
	msg := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\r\n"
	msg += fmt.Sprintf("From: %s\r\n", mail.From.Address)
	msg += fmt.Sprintf("To: %s\r\n", strings.Join(mail.To, ";"))
	msg += fmt.Sprintf("Subject: %s\r\n", mail.Subject)
	msg += fmt.Sprintf("\r\n%s\r\n", mail.Body)
	return msg
}

func SendTextEmailOtp(to []string, from string, otp string) error {
	contentEmail := Email{
		From: EmailAddress{
			Address: from,
			Name:    "Go Ecommerce",
		},
		To:      to,
		Subject: "OTP Verification",
		Body:    fmt.Sprintf("Your OTP is %s. Please enter this OTP to verify your account.", otp),
	}

	msg := BuildMessage(contentEmail)

	auth := smtp.PlainAuth("", SMTPEmail, SMTPPassword, SMTPHost)
	err := smtp.SendMail(SMTPHost+":"+SMTPPort, auth, from, to, []byte(msg))
	if err != nil {
		global.Logger.Error("Failed to send email: ", zap.Error(err))
		return err
	}
	return nil
}

func SendTemplateEmailOtp(
	to []string, from string, nameTemplate string,
	dataTemplate map[string]interface{},
) error {
	lg := global.Logger

	htmlTemplate, err := getMailTemplate(nameTemplate, dataTemplate)
	if err != nil {
		lg.Error("Failed to get email template: ", zap.Error(err))
		return err
	}

	err = send(to, from, htmlTemplate)
	if err != nil {
		lg.Error("Failed to send email: ", zap.Error(err))
	}

	return err
}

func getMailTemplate(nameTemplate string, dataTemplate map[string]interface{}) (string, error) {
	htmlTemplate := new(bytes.Buffer)
	t := template.Must(template.New(nameTemplate).ParseFiles("templates-email/" + nameTemplate))
	err := t.Execute(htmlTemplate, dataTemplate)
	if err != nil {
		return "", err
	}
	return htmlTemplate.String(), nil
}

func send(to []string, from string, htmlTemplate string) error {
	contentEmail := Email{
		From:    EmailAddress{Address: from, Name: "Go Ecommerce"},
		To:      to,
		Subject: "OTP Verification",
		Body:    htmlTemplate,
	}

	msg := BuildMessage(contentEmail)

	auth := smtp.PlainAuth("", SMTPEmail, SMTPPassword, SMTPHost)
	err := smtp.SendMail(SMTPHost+":"+SMTPPort, auth, from, to, []byte(msg))
	if err != nil {
		global.Logger.Error("Failed to send email: ", zap.Error(err))
		return err
	}
	return nil
}
