package database

import (
	"github.com/sdslabs/nymeria/internal/database"
	"github.com/sdslabs/nymeria/internal/database/schema"
)

func CreateOTPEntry(otpEntry *schema.OTP) error {

	var existingOTP schema.OTP
	tx := database.DB.First(&existingOTP, "email = ?", otpEntry.Email)
	if tx.Error == nil {
		existingOTP.Code = otpEntry.Code
		existingOTP.ExpiresAt = otpEntry.ExpiresAt
		return database.DB.Save(&existingOTP).Error
	}

	return database.DB.Create(otpEntry).Error
}

func QueryOTPEntry(email string) (schema.OTP, error) {
	var otpEntry schema.OTP

	tx := database.DB.Where("email = ?", email).First(&otpEntry)

	return otpEntry, tx.Error
}

func VerifyOTPEntry(email string) error {

	return database.DB.Model(&schema.OTP{}).Where("email = ?", email).Update("verified", true).Error
}
