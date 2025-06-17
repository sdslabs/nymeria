// Copyright (c) 2025 SDSLabs
// SPDX-License-Identifier: MIT

package totp

import (
	"github.com/pkg/errors"

	"github.com/pquerna/otp"
	stdtotp "github.com/pquerna/otp/totp"
)

// Generate a new key with the TOTP package
// TODO: update issue and accountName parameters
func GenerateKey(issuer string, accountName string) (*otp.Key, error) {
	key, err := stdtotp.Generate(stdtotp.GenerateOpts{
		Issuer:      issuer,
		AccountName: accountName,
		SecretSize:  20,
		Digits:      otp.DigitsSix,
		Period:      30,
	})

	if err != nil {
		return nil, errors.WithStack(err)
	}

	return key, err
}
