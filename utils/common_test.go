package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsValidAlphabet(t *testing.T) {
	result := IsValidAlphabet("alphabet")
	assert.True(t, result)

	result = IsValidAlphabet("123")
	assert.False(t, result)
}

func TestIsValidNumeric(t *testing.T) {
	result := IsValidNumeric("123")
	assert.True(t, result)

	result = IsValidNumeric("invalid format string")
	assert.False(t, result)
}

func TestIsValidAlphaNumeric(t *testing.T) {
	result := IsValidAlphaNumeric("Alpha123")
	assert.True(t, result)

	result = IsValidAlphaNumeric("!@#")
	assert.False(t, result)
}

func TestIsValidAlphaNumericHyphen(t *testing.T) {
	result := IsValidAlphaNumericHyphen("valid-number-2-with-hyphen")
	assert.True(t, result)

	result = IsValidAlphaNumericHyphen("")
	assert.False(t, result)
}

func TestFormattedTime(t *testing.T) {
	result := FormattedTime("2022-12-20T12:34:56Z")
	assert.Equal(t, "2022-12-20 12:34:56", result)

	result = FormattedTime("invalid time string")
	assert.Equal(t, "", result)
}
