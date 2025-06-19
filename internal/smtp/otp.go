package smtp

import (
	"errors"
	"regexp"
	"time"

	"gorm.io/gorm"

	"github.com/sdslabs/nymeria/internal/config"
	database "github.com/sdslabs/nymeria/internal/database/otp"
	"github.com/sdslabs/nymeria/internal/database/schema"
	"github.com/sdslabs/nymeria/internal/logger"
)

func SendOTPHandler(email string, isIITRCheck bool) (otp string, err error) {

	smtpHost := config.AppConfig.SMTPHost
	smtpPort := config.AppConfig.SMTPPort

	if smtpHost == "" || smtpPort == "" {
		logger.Warn().Msg("SMTP not configured")
		return "", errors.New("SMTP not configured")
	}

	re := regexp.MustCompile(`^.*@.*iitr\.ac\.in$`)
	isIITR := re.MatchString(email)

	if isIITRCheck && !isIITR {
		return "", errors.New("email should be of IITR domain")
	}

	otp = generateOTP()
	expiry := time.Now().Add(time.Duration(config.AppConfig.OTPExpiry) * time.Minute) // OTP expires in 5 minutes

	otpEntry, err := database.QueryOTPEntry(email)

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			otpEntry = schema.OTP{
				Email:     email,
				Code:      otp,
				ExpiresAt: expiry,
				Verified:  false,
			}
		} else {
			logger.Err(err).Msg("Failed to query OTP")
			return "", errors.New("failed to send OTP")
		}
	}

	if otpEntry.Verified {
		return "", errors.New("email already verified")
	}

	otpEntry.Code = otp
	otpEntry.ExpiresAt = expiry

	err = database.CreateOTPEntry(&otpEntry)
	if err != nil {
		logger.Err(err).Msg("Failed to store OTP")
		return "", errors.New("failed to store OTP")
	}

	// Send OTP to email
	err = sendEmail(email, otp)
	if err != nil {
		logger.Err(err).Msg("Failed to send OTP")
		return "", errors.New("failed to send OTP")
	}

	return otp, nil
}

func VerifyOTPHandler(email string, otp string) error {

	smtpHost := config.AppConfig.SMTPHost
	smtpPort := config.AppConfig.SMTPPort

	if smtpHost == "" || smtpPort == "" {
		logger.Warn().Msg("SMTP not configured")
		return errors.New("SMTP not configured")
	}

	otpEntry, err := database.QueryOTPEntry(email)

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("OTP not found")
		}
		logger.Err(err).Msg("Failed to query OTP")
		return errors.New("failed to verify OTP")
	}

	if otpEntry.Verified {
		return errors.New("email already verified")
	}

	if otpEntry.Code != otp {
		return errors.New("wrong OTP")
	}

	if time.Now().After(otpEntry.ExpiresAt) {
		return errors.New("OTP expired")
	}

	err = database.VerifyOTPEntry(email)

	if err != nil {
		logger.Err(err).Msg("Failed to verify OTP")
		return errors.New("failed to verify OTP")
	}

	return nil
}
