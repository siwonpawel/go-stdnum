package pl

import (
	"errors"
	"github.com/siwonpawel/go-stdnum/stdnum/validation"
	"regexp"
	"strconv"
	"strings"
)

var regonValidRegex = regexp.MustCompile("^\\d{2]\\d{7}(\\d{5})?")

var nineDigitRegonWeights = []int{8, 9, 2, 3, 4, 5, 6, 7}
var fourteenDigitRegonWeights = []int{2, 4, 8, 5, 0, 9, 7, 3, 6, 1, 2, 4, 8}

var regonResultCreator = validation.NewResultCreator(country, regonIdentifierName)

func ValidateREGON(number string) *validation.Result {
	cleanedNumber, warnings := validation.Cleanse(number, country)

	debugInfo := validation.DebugInfo{
		CleanedInput: cleanedNumber,
	}

	if !regonValidRegex.MatchString(cleanedNumber) {
		return regonResultCreator.Fail(number, warnings, validation.InvalidLength, debugInfo)
	}

	isValid, err := validateRegon(cleanedNumber)
	if err != nil {
		return regonResultCreator.Fail(number, warnings, err.Error(), debugInfo)
	}

	if isValid {
		return regonResultCreator.Ok(number, warnings, debugInfo)
	} else {
		return regonResultCreator.Fail(number, warnings, validation.InvalidNumber, debugInfo)
	}
}

func validateRegon(number string) (bool, error) {
	if len(number) == 14 {
		return validateRegonLong(number)
	} else if len(number) == 9 {
		return validateRegonShort(number)
	} else if len(number) == 7 {
		return validateRegonShort("00" + number)
	} else {
		return false, errors.New(validation.InvalidLength)
	}
}

func validateRegonShort(number string) (bool, error) {
	if len(number) != 9 {
		return false, errors.New(validation.InvalidLength)
	}

	splitted := strings.Split(number, "")
	sum := 0
	for i := 0; i < 8; i++ {
		digit, err := strconv.Atoi(splitted[i])
		if err != nil {
			return false, nil
		}

		sum += digit * nineDigitRegonWeights[i]
	}

	lastDigit, err := strconv.Atoi(splitted[8])
	if err != nil {
		return false, nil
	}

	checksum := sum % 11
	if checksum == 10 {
		checksum = 0
	}

	return checksum == lastDigit, nil
}

func validateRegonLong(number string) (bool, error) {
	if len(number) != 14 {
		return false, errors.New(validation.InvalidLength)
	}

	splitted := strings.Split(number, "")
	sum := 0
	for i := 0; i < 13; i++ {
		digit, err := strconv.Atoi(splitted[i])
		if err != nil {
			return false, nil
		}

		sum += digit * fourteenDigitRegonWeights[i]
	}

	lastDigit, err := strconv.Atoi(splitted[13])
	if err != nil {
		return false, nil
	}

	checksum := sum % 11
	if checksum == 10 {
		checksum = 0
	}

	return checksum == lastDigit, nil
}
