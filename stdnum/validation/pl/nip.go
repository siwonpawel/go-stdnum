package pl

import (
	"github.com/siwonpawel/go-stdnum/stdnum/validation"
	"log"
	"regexp"
	"strconv"
	"strings"
)

var nipValidRegex = regexp.MustCompile("\\d{10}")
var nipCleanseRegex = regexp.MustCompile("[^A-z0-9]|[\\\\\\[\\]`^]")
var nipNumericalWeights = []int{6, 5, 7, 2, 3, 4, 5, 6, 7}

var nipResultCreator = validation.NewResultCreator(country, nipIdentifierName)

var compact = func(src string) string {
	return nipCleanseRegex.ReplaceAllString(src, "")
}

func cleanse(number string) (cleanedNumber string, warnings []string) {
	cleanedNumber = strings.ToUpper(compact(number))

	if !strings.HasPrefix(cleanedNumber, country) {
		return cleanedNumber, []string{missingCountryPrefix}
	} else {
		return strings.Trim(cleanedNumber, country), []string{}
	}
}

func Validate(number string) *validation.Result {
	cleanedNumber, warnings := cleanse(number)

	debugInfo := validation.DebugInfo{
		CleanedInput: cleanedNumber,
	}

	if !nipValidRegex.MatchString(cleanedNumber) {
		return nipResultCreator.Fail(number, warnings, invalidLength, debugInfo)
	}

	isValid, err := validate(cleanedNumber)
	if err != nil {
		return nipResultCreator.Fail(number, warnings, err.Error(), debugInfo)
	}

	if isValid {
		return nipResultCreator.Ok(number, warnings, debugInfo)
	} else {
		return nipResultCreator.Fail(number, warnings, invalidNumber, debugInfo)
	}
}

func validate(number string) (bool, error) {

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
