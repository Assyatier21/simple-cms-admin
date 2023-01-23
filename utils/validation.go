package utils

import "regexp"

func IsValidAlphabet(s string) bool {
	regex, _ := regexp.Compile(`^[a-zA-Z]*$`)
	return regex.MatchString(s)
}

func IsValidNumeric(s string) bool {
	regex, _ := regexp.Compile(`([0-9])`)
	return regex.MatchString(s)
}

func IsValidAlphaNumeric(s string) bool {
	regex, _ := regexp.Compile(`(^[a-zA-Z0-9]*$)`)
	return regex.MatchString(s)
}

func IsValidSlug(s string) bool {
	regex, _ := regexp.Compile(`^[a-z0-9-]+$`)
	return regex.MatchString(s)
}
