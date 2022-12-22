package utils

import (
	"errors"
	"fmt"
	"log"
	"regexp"
	"time"
)

var (
	ErrNotFound = errors.New("data not found")
	TimeNow     = fmt.Sprintf("%d-%d-%d %d:%d:%d", time.Now().Year(), time.Now().Month(), time.Now().Day(), time.Now().Hour(), time.Now().Minute(), time.Now().Second())
)

func IsValidAlphabet(s string) bool {
	regex, _ := regexp.Compile(`^[a-zA-Z ]*$`)
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

func IsValidAlphaNumericHyphen(s string) bool {
	regex, _ := regexp.Compile(`[a-zA-Z0-9-]+`)
	return regex.MatchString(s)
}

func FormattedTime(ts string) string {
	t, err := time.Parse(time.RFC3339, ts)
	if err != nil {
		log.Println(err)
		return ""
	}

	formattedTime := t.Format("2006-01-02 15:04:05")
	return formattedTime
}
