package pl

import (
	"errors"
	"github.com/siwonpawel/go-stdnum/stdnum/validation"
	"regexp"
	"strconv"
	"strings"
)

var peselValidRegex = regexp.MustCompile("^\\d{2}((0[1-9])|(1[1-2]))((0[1-9])|([1-2][0-9])|3[0-1])\\d{5}$")
var peselNumericalWeights = []int{1, 3, 7, 9, 1, 3, 7, 9, 1, 3}

var peselResultCreator = validation.NewResultCreator(country, peselIdentifierName)

func ValidatePESEL(number string) *validation.Result {
	cleanedNumber, warnings := validation.Cleanse(number, country)

	debugInfo := validation.DebugInfo{
		CleanedInput: cleanedNumber,
	}

	if !peselValidRegex.MatchString(cleanedNumber) {
		return peselResultCreator.Fail(number, warnings, validation.InvalidLength, debugInfo)
	}

	isValid, err := validatePESEL(cleanedNumber)
	if err != nil {
		return peselResultCreator.Fail(number, warnings, err.Error(), debugInfo)
	}

	if isValid {
		return peselResultCreator.Ok(number, warnings, debugInfo)
	} else {
		return peselResultCreator.Fail(number, warnings, validation.InvalidNumber, debugInfo)
	}
}

func validatePESEL(cleanedNumber string) (bool, error) {
	if len(cleanedNumber) != 11 {
		return false, errors.New(validation.InvalidLength)
	}

	sum := 0
	split := strings.Split(cleanedNumber, "")
	for i := 0; i < 10; i++ {
		atoi, err := strconv.Atoi(split[i])
		if err != nil {
			return false, err
		}

		sum += atoi * peselNumericalWeights[i]
	}

	checksum, _ := strconv.Atoi(split[10])
	if checksum != 0 {
		checksum = 10 - checksum
	}

	return (sum)%10 == checksum, nil
}
