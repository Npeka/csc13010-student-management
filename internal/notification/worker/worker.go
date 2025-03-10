package worker

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/smtp"
	"strings"
	"text/template"

	"github.com/csc13010-student-management/internal/events"
	"github.com/csc13010-student-management/internal/notification"
	kafkas "github.com/csc13010-student-management/pkg/kafka"
	"github.com/csc13010-student-management/pkg/logger"
	"github.com/opentracing/opentracing-go"
	"github.com/pkg/errors"
	"github.com/segmentio/kafka-go"
)

type notificationWorker struct {
	nr notification.INotificationRepository
	lg *logger.LoggerZap
}

func NewNotificationWorker(
	nr notification.INotificationRepository,
	lg *logger.LoggerZap,
) notification.INotificationWorker {
	return &notificationWorker{
		nr: nr,
		lg: lg,
	}
}

func (nw *notificationWorker) Start(kurl string) {
	krs := []kafkas.KafkaReader{
		{
			Topic:   string(events.NotiCreate),
			GroupID: fmt.Sprintf("%v.g", events.NotiCreate),
			Handler: nw.HandleNotiCreateEvent,
			MaxIns:  1,
		},
	}
	kafkas.StartKafkaConsumers(kurl, krs)
}

type NotiTemplate struct {
	Subject  string
	Template string
}

var notiTemplate = map[events.NotiType]NotiTemplate{
	events.NotiUserResetPasswordOTP: {
		Subject:  "OTP Reset Password",
		Template: "user_reset_password_otp.html",
	},
	events.NotiStudentStatusChanged: {
		Subject:  "Your status has been changed",
		Template: "student_status_changed.html",
	},
}

func (nw *notificationWorker) HandleNotiCreateEvent(ctx context.Context, msg kafka.Message) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "notiWorker.HandleCreateNotiEvent")
	defer span.Finish()

	var event events.NotificationEvent
	if err := json.Unmarshal(msg.Value, &event); err != nil {
		logger.LoggerFuncError(nw.lg, errors.Wrap(err, "notiWorker.HandleCreateNotiEvent.json.Unmarshal"))
		return err
	}

	if event.Email != "" {
		err := SendTemplateEmail(
			[]string{event.Email},
			"seimeicc@gmail.com",
			notiTemplate[event.Type].Subject,
			notiTemplate[event.Type].Template,
			event.Data,
		)
		if err != nil {
			return err
		}
	}

	return nil
}

const (
	SMTPHost     = "smtp.gmail.com"
	SMTPPort     = "587"
	SMTPEmail    = "seimeicc@gmail.com"
	SMTPPassword = "ctfp eowe ngzu vtma"
)

type EmailConfig struct {
	SMTPHost     string
	SMTPPort     string
	SMTPEmail    string
	SMTPPassword string
}

var emailConfig = EmailConfig{
	SMTPHost:     "smtp.gmail.com",
	SMTPPort:     "587",
	SMTPEmail:    "seimeicc@gmail.com",
	SMTPPassword: "ctfp eowe ngzu vtma",
}

type Email struct {
	From    string
	To      []string
	Subject string
	Body    string
}

func BuildMessage(mail Email) string {
	msg := "MIME-Version: 1.0\r\nContent-Type: text/html; charset=\"UTF-8\"\r\n"
	msg += fmt.Sprintf("From: %s\r\n", mail.From)
	msg += fmt.Sprintf("To: %s\r\n", strings.Join(mail.To, ";"))
	msg += fmt.Sprintf("Subject: %s\r\n", mail.Subject)
	msg += "\r\n" + mail.Body
	return msg
}

func SendEmailBase(to []string, from string, subject string, body string) error {
	email := Email{
		From:    from,
		To:      to,
		Subject: subject,
		Body:    body,
	}

	msg := BuildMessage(email)
	auth := smtp.PlainAuth("", emailConfig.SMTPEmail, emailConfig.SMTPPassword, emailConfig.SMTPHost)

	err := smtp.SendMail(emailConfig.SMTPHost+":"+emailConfig.SMTPPort, auth, from, to, []byte(msg))
	if err != nil {
		log.Printf("Lỗi gửi email: %v", err)
		return err
	}

	log.Println("Email gửi thành công!")
	return nil
}

// SendTemplateEmail: Gửi email sử dụng template
func SendTemplateEmail(to []string, from string, subject string, templateName string, data map[string]interface{}) error {
	htmlBody, err := GetMailTemplate(templateName, data)
	if err != nil {
		return err
	}

	return SendEmailBase(to, from, subject, htmlBody)
}

// GetMailTemplate: Load và parse template HTML
func GetMailTemplate(templateName string, data map[string]interface{}) (string, error) {
	var buf bytes.Buffer
	t, err := template.ParseFiles("internal/notification/templates/" + templateName)
	if err != nil {
		log.Printf("Lỗi tải template: %v", err)
		return "", err
	}

	err = t.Execute(&buf, data)
	if err != nil {
		log.Printf("Lỗi render template: %v", err)
		return "", err
	}

	return buf.String(), nil
}
