package pl

import (
	"github.com/siwonpawel/go-stdnum/stdnum/validation"
	"log"
	"regexp"
	"strconv"
	"strings"
)

var nipValidRegex = regexp.MustCompile("^\\d{10}$")
var nipNumericalWeights = []int{6, 5, 7, 2, 3, 4, 5, 6, 7}

var nipResultCreator = validation.NewResultCreator(country, nipIdentifierName)

func ValidateNIP(number string) *validation.Result {
	cleanedNumber, warnings := validation.Cleanse(number, country)

	debugInfo := validation.DebugInfo{
		CleanedInput: cleanedNumber,
	}

	if !nipValidRegex.MatchString(cleanedNumber) {
		return nipResultCreator.Fail(number, warnings, validation.InvalidLength, debugInfo)
	}

	isValid, err := validateNIP(cleanedNumber)
	if err != nil {
		return nipResultCreator.Fail(number, warnings, err.Error(), debugInfo)
	}

	if isValid {
		return nipResultCreator.Ok(number, warnings, debugInfo)
	} else {
		return nipResultCreator.Fail(number, warnings, validation.InvalidNumber, debugInfo)
	}
}

func validateNIP(number string) (bool, error) {

	sum := 0
	splitNumber := strings.Split(number, "")
	for i := 0; i < 9; i++ {
		numVal, err := strconv.Atoi(splitNumber[i])
		if err != nil {
			log.Printf("Error [ %v ] converting to number: %#v", err.Error(), numVal)
		}

		sum += numVal * nipNumericalWeights[i]
	}

	if checksum, err := strconv.Atoi(splitNumber[9]); err != nil {
		return false, err
	} else {
		return sum%11 == checksum, nil
	}
}
