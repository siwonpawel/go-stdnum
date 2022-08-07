package pl

import (
	"github.com/siwonpawel/go-stdnum/stdnum/validation"
)

func init() {
	validation.RegisterValidator(ValidateNIP, country, nipIdentifierName)
	validation.RegisterValidator(ValidatePESEL, country, peselIdentifierName)
	validation.RegisterValidator(ValidateREGON, country, regonIdentifierName)
}
