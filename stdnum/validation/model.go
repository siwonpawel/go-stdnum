package validation

type Result struct {
	Input          string
	Country        string
	IdentifierName string
	IsValid        bool
	Warnings       []string
	Error          validationError
	DebugInfo      DebugInfo
}

type validationError struct {
	msg string
}

func (e *validationError) Error() string {
	return e.msg
}

type DebugInfo struct {
	CleanedInput string
}

type ResultCreator struct {
	country        string
	identifierName string
}

func NewResultCreator(country, identifierName string) ResultCreator {
	return ResultCreator{
		country:        country,
		identifierName: identifierName,
	}
}

func (rc *ResultCreator) Ok(validationInput string, warnings []string, debugInfo DebugInfo) *Result {
	return &Result{
		Input:          validationInput,
		Country:        rc.country,
		IdentifierName: rc.identifierName,
		IsValid:        true,
		Warnings:       warnings,
		DebugInfo:      debugInfo,
	}
}

func (rc *ResultCreator) Fail(validationInput string, warnings []string, errorMessage string, debugInfo DebugInfo) *Result {
	return &Result{
		Input:          validationInput,
		Country:        rc.country,
		IdentifierName: rc.identifierName,
		IsValid:        false,
		Warnings:       warnings,
		Error:          validationError{msg: errorMessage},
		DebugInfo:      debugInfo,
	}
}

type Results []*Result

func (r Results) GetValid() Results {
	filtered := make(Results, 0)

	for _, v := range r {
		if v.IsValid {
			filtered = append(filtered, v)
		}
	}

	return filtered
}
