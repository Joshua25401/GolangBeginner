package validators

import "github.com/go-playground/validator/v10"

type ValidateError struct {
	FailedField string
	Tag         string
	Value       string
}

func ValidateStruct(obj interface{}) []*ValidateError {
	var errors []*ValidateError

	validate := validator.New()
	err := validate.Struct(obj)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			errors = append(errors, &ValidateError{
				FailedField: err.StructNamespace(),
				Tag:         err.Tag(),
				Value:       err.Param(),
			})
		}
	}
	return errors
}
