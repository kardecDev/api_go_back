package validation

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/kardecDev/api_go_back/internal/handler/httperr"
)

func ValidateHttpData(d interface{}) *httperr.RestErr {
	val := validator.New(validator.WithRequiredStructEnabled())

	//extract json tag name
	val.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
		if name == "-" {
			return ""
		}
		return name
	})

	if err := val.Struct(d); err != nil {
		var errorsCauses []httperr.Fields
		for _, e := range err.(validator.ValidationErrors) {
			cause := httperr.Fields{}
			fieldName := e.Field()
			switch e.Tag() {
			case "required":
				cause.Message = fmt.Sprintf("%s is required", fieldName)
				cause.Field = fieldName
				cause.Value = e.Field()
			case "uuid4":
				cause.Message = fmt.Sprintf("%s is not a valid uuid", fieldName)
				cause.Field = fieldName
				cause.Value = e.Field()
			case "boolean":
				cause.Message = fmt.Sprintf("%s is not a valid boolean", fieldName)
				cause.Field = fieldName
				cause.Value = e.Field()
			case "min":
				cause.Message = fmt.Sprintf("%s must be grather than %s", fieldName, e.Param())
				cause.Field = fieldName
				cause.Value = e.Field()
			case "max":
				cause.Message = fmt.Sprintf("%s must be less than %s", fieldName, e.Param())
				cause.Field = fieldName
				cause.Value = e.Field()
			case "email":
				cause.Message = fmt.Sprintf("%s is not a valid email", fieldName)
				cause.Field = fieldName
				cause.Value = e.Field()
			case "containsany":
				cause.Message = fmt.Sprintf("%s must contain at least one of the following characters: !@#$%%*", fieldName)
				cause.Field = fieldName
				cause.Value = e.Value()
			default:
				cause.Message = "invalid field"
				cause.Field = fieldName
				cause.Value = e.Field()

			}
			errorsCauses = append(errorsCauses, cause)
		}
		return httperr.NewBadRequestValidationError("some fields are invalid", errorsCauses)
	}
	return nil
}
