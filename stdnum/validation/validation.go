package validation

var validators = []validator{}

func RegisterValidator(validationFunction func(string) *Result, country, identifier string) {
	v := validator{
		Country:            country,
		Identifier:         identifier,
		validationFunction: validationFunction,
	}
	validators = append(validators, v)
}

type validator struct {
	Country            string
	Identifier         string
	validationFunction func(string) *Result
}

type Validator interface {
	Validate(string) *Result
}

func (v *validator) Validate(input string) *Result {
	return v.validationFunction(input)
}

func ValidateAll(input string) Results {
	results := Results{}
	for _, validator := range validators {
		results = append(results, validator.Validate(input))
	}

	return results
}
