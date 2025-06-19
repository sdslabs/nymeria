package smtp

import (
	"bytes"
	"crypto/tls"
	"fmt"
	"html/template"
	"math/rand"
	"net/smtp"
	"os"
	"path/filepath"
	"time"

	"github.com/sdslabs/nymeria/internal/config"
	"github.com/sdslabs/nymeria/internal/logger"
)

func generateOTP() string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return fmt.Sprintf("%06d", r.Intn(1000000)) // 6-digit OTP
}

// sendEmail sends an OTP email using an SMTP client with TLS. Falls back to plain text if template is missing.
func sendEmail(email, otp string) error {
	from := config.AppConfig.MailFrom
	password := config.AppConfig.MailPassword
	smtpHost := config.AppConfig.SMTPHost
	smtpPort := config.AppConfig.SMTPPort

	// Email subject
	subject := "Your OTP Code"

	// Path to email template
	emailTemplatePath := filepath.Join("internal", "templates", "email_template.html")

	// Check if template file exists
	var body bytes.Buffer
	_, err := os.Stat(emailTemplatePath)
	if err == nil {
		// Template exists, parse and execute
		tmpl, err := template.ParseFiles(emailTemplatePath)
		if err != nil {
			logger.Err(err).Msg("Failed to read email template")
			return err
		}

		// Create structured data matching the template expectations
		emailData := struct {
			Subject          string
			AppName          string
			Title            string
			Message          string
			VerificationCode string
			ExpiryTime       int
		}{
			Subject:          subject,
			AppName:          "Nymeria", // You can make this configurable
			Title:            "Email Verification",
			Message:          "Please use the verification code below to complete your authentication.",
			VerificationCode: otp,
			ExpiryTime:       5, // 5 minutes expiry
		}

		if err := tmpl.Execute(&body, emailData); err != nil {
			logger.Err(err).Msg("Failed to execute email template")
			return err
		}
	} else {
		// Template does not exist, send plain text email
		logger.Warn().Msg("Template not found, sending plain text email")
		body.WriteString(fmt.Sprintf("Hello,\n\nYour OTP is: %s\nThis OTP will expire in 5 minutes.\n\nRegards,\nTeam", otp))
	}

	// Create email headers
	message := fmt.Sprintf("From: %s\r\n", from) +
		fmt.Sprintf("To: %s\r\n", email) +
		fmt.Sprintf("Subject: %s\r\n", subject) +
		"MIME-Version: 1.0\r\n"

	// Set Content-Type based on template availability
	bodyString := body.String()
	if len(bodyString) > 0 && bodyString[0] == '<' {
		message += "Content-Type: text/html; charset=\"utf-8\"\r\n\r\n"
	} else {
		message += "Content-Type: text/plain; charset=\"utf-8\"\r\n\r\n"
	}

	message += bodyString

	// Setup TLS connection
	tlsConfig := &tls.Config{
		InsecureSkipVerify: true, // Set true only if SMTP server uses self-signed certs
		ServerName:         smtpHost,
	}

	// Connect to SMTP server
	conn, err := tls.Dial("tcp", smtpHost+":"+smtpPort, tlsConfig)
	if err != nil {
		logger.Info().Str("smtpHost", smtpHost).Str("smtpPort", smtpPort).Msg("Failed to connect to SMTP server")
		logger.Err(err).Msg("Failed to connect to SMTP server")
		return err
	}

	client, err := smtp.NewClient(conn, smtpHost)
	if err != nil {
		logger.Err(err).Msg("Failed to create SMTP client")
		return err
	}
	defer client.Close()

	// Authenticate
	auth := smtp.PlainAuth("", from, password, smtpHost)
	if err := client.Auth(auth); err != nil {
		logger.Err(err).Msg("SMTP authentication failed")
		return err
	}

	// Set sender and recipient
	if err := client.Mail(from); err != nil {
		logger.Err(err).Msg("Failed to set sender")
		return err
	}

	if err := client.Rcpt(email); err != nil {
		logger.Err(err).Msg("Failed to set recipient")
		return err
	}

	// Write email data
	w, err := client.Data()
	if err != nil {
		logger.Err(err).Msg("Failed to get SMTP data writer")
		return err
	}

	_, err = w.Write([]byte(message))
	if err != nil {
		logger.Err(err).Msg("Failed to write email content")
		return err
	}

	err = w.Close()
	if err != nil {
		logger.Err(err).Msg("Failed to close SMTP writer")
		return err
	}

	// Quit SMTP session
	if err := client.Quit(); err != nil {
		logger.Err(err).Msg("Failed to close SMTP connection")
		return err
	}

	logger.Info().Str("email", email).Msg("OTP email sent successfully")
	return nil
}
