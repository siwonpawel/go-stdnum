package validation

import (
	"regexp"
	"strings"
)

const (
	InvalidNumber        = "Invalid."
	InvalidLength        = "Invalid input length."
	MissingCountryPrefix = "Missing country identifier at beginning."
)

var cleanseRegex = regexp.MustCompile("[^A-z0-9]|[\\\\\\[\\]`^]")

func Cleanse(number, country string) (cleanedNumber string, warnings []string) {
	cleanedNumber = strings.ToUpper(cleanseRegex.ReplaceAllString(number, ""))

	if !strings.HasPrefix(cleanedNumber, country) {
		return cleanedNumber, []string{MissingCountryPrefix}
	} else {
		return strings.Trim(cleanedNumber, country), []string{}
	}
}
