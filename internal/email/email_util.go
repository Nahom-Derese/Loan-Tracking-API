package emailutil

import (
	"fmt"
	"log"
	"net/smtp"

	"github.com/Nahom-Derese/Loan-Tracking-API/bootstrap"
)

func SendVerificationEmail(recipientEmail string, VerificationToken string, env *bootstrap.Env) error {
	// Email configuration
	from := env.SenderEmail
	password := env.SenderPassword
	smtpHost := env.SmtpHost
	smtpPort := env.SmtpPort

	baseUrl := env.ServerAddress

	subject := "Subject: Account Verification\n"
	mime := "MIME-Version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	url := fmt.Sprintf("%v/users/verify-email/%v", baseUrl, VerificationToken)
	// print the url
	fmt.Println(url)
	body := Emailtemplate(url)
	message := []byte(subject + mime + "\n" + body)
	auth := smtp.PlainAuth("", from, password, smtpHost)

	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, []string{recipientEmail}, message)
	if err != nil {
		log.Fatalf("Failed to send email: %v", err)
	}

	return nil

}
