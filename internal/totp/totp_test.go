// Copyright (c) 2025 SDSLabs
// SPDX-License-Identifier: MIT

package totp_test

import (
	"testing"

	"github.com/sdslabs/nymeria/internal/totp"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGenerator(t *testing.T) {
	key, err := totp.GenerateKey("sdslabs.co", "segfault")
	require.NoError(t, err)
	assert.Equal(t, "sdslabs.co", key.Issuer())
}
